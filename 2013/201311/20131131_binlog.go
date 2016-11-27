/*
http://segmentfault.com/a/1190000003072963
一个事件头有 19 字节，依次排列为四字节的时间戳、一字节的当前事件类型、四字节的服务端 ID、四字节的当前事件长度描述、四字节的下个事件位置（方便跳转）以及两字节的标识。

　　用 ASCII Diagram 表示如下：

+---------+---------+---------+------------+-------------+-------+
|timestamp|type code|server_id|event_length|next_position|flags  |
|4 bytes  |1 byte   |4 bytes  |4 bytes     |4 bytes      |2 bytes|
+---------+---------+---------+------------+-------------+-------+

struct BinlogEventHeader
{
    int   timestamp;
    char  type_code;
    int   server_id;
    int   event_length;
    int   next_position;
    char  flags[2];
};
如果你要直接用这个结构体来读取数据的话，需要加点手脚。

因为默认情况下 GCC 或者 G++ 编译器会对结构体进行字节对齐，这样读进来的数据就不对了，因为 Binlog 并不是对齐的。为了统一我们需要取消这个结构体的字节对齐，一个方法是使用 #pragma pack(n)，一个方法是使用 __attribute__((__packed__))，还有一种情况是在编译器编译的时候强制把所有的结构体对其取消，即在编译的时候使用 fpack-struct 参数，如：

```sh
$ g++ temp.cpp -o a -fpack-struct=1

而具体有哪些事件则在 https://github.com/mysql/mysql-server/blob/5.7/libbinlogevents/include/binlog_event.h#L245
里面被定义。如有个 FORMAT_DESCRIPTION_EVENT 事件的 type_code 是 15、UPDATE_ROWS_EVENT 的 type_code 是 31。
还有那个 next_position，在 v4 版本中代表从 Binlog 一开始到下一个事件开始的偏移量，比如到第一个事件的 next_position 就是 4，因为文件头有一个字节的长度。然后接下去对于事件 n 和事件 n + 1 来说，他们有这样的关系：

next_position(n + 1) = next_position(n) + event_length(n)

事件体

　　事实上在 Binlog 事件中应该是有三个部分组成，header、post-header 和 payload，不过通常情况下我们把 post-header 和 payload 都归结为事件体，实际上这个 post-header 里面放的是一些定长的数据，只不过有时候我们不需要特别地关心。想要深入了解可以去查看 MySQL 的官方文档。

　　所以实际上一个真正的事件体由两部分组成，用 ASCII Diagram 表示就像这样：

+=====================================+
| event  | fixed part (post-header)   |
| data   +----------------------------+
|        | variable part (payload)    |
+=====================================+
　　而这个 post-header 对于不同类型的事件来说长度是不一样的，同种类型来说是一样的，而这个长度的预先规定将会在一个“格式描述事件”中定好。

格式描述事件

　　在上文我们有提到过，在 Magic Number 之后跟着的是一个格式描述事件（Format Description Event），其实这只是在 v4 版本中的称呼，在以前的版本里面叫起始事件（Start Event）。

　　在 v4 版本中这个事件的结构如下面的 ASCII Diagram 所示。

+=====================================+
| event  | timestamp         0 : 4    |
| header +----------------------------+
|        | type_code         4 : 1    | = FORMAT_DESCRIPTION_EVENT = 15
|        +----------------------------+
|        | server_id         5 : 4    |
|        +----------------------------+
|        | event_length      9 : 4    | >= 91
|        +----------------------------+
|        | next_position    13 : 4    |
|        +----------------------------+
|        | flags            17 : 2    |
+=====================================+
| event  | binlog_version   19 : 2    | = 4
| data   +----------------------------+
|        | server_version   21 : 50   |
|        +----------------------------+
|        | create_timestamp 71 : 4    |
|        +----------------------------+
|        | header_length    75 : 1    |
|        +----------------------------+
|        | post-header      76 : n    | = array of n bytes, one byte per event
|        | lengths for all            |   type that the server knows about
|        | event types                |
+=====================================+
　　这个事件的 type_code 是 15，然后 event_length 是大于等于 91 的值的，这个主要取决于所有事件类型数。

　　因为从第 76 字节开始后面的二进制就代表一个字节类型的数组了，一个字节代表一个事件类型的 post-header 长度，即每个事件类型固定数据的长度。

　　那么按照上述的一些线索来看，我们能非常快地写出一个简单的解读 Binlog 格式描述事件的代码。

*/
// go-mysqlbinlog: a simple binlog tool to sync remote MySQL binlog.
// go-mysqlbinlog supports semi-sync mode like facebook mysqlbinlog.
// see http://yoshinorimatsunobu.blogspot.com/2014/04/semi-synchronous-replication-at-facebook.html
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
)

var host = flag.String("host", "127.0.0.1", "MySQL host")
var port = flag.Int("port", 3306, "MySQL port")
var user = flag.String("user", "root", "MySQL user, must have replication privilege")
var password = flag.String("password", "", "MySQL password")

var flavor = flag.String("flavor", "mysql", "Flavor: mysql or mariadb")

var file = flag.String("file", "", "Binlog filename")
var pos = flag.Int("pos", 4, "Binlog position")

var semiSync = flag.Bool("semisync", false, "Support semi sync")
var backupPath = flag.String("backup_path", "", "backup path to store binlog files")

var rawMode = flag.Bool("raw", false, "Use raw mode")

func main() {
	flag.Parse()

	b := replication.NewBinlogSyncer(101, *flavor)

	if err := b.RegisterSlave(*host, uint16(*port), *user, *password); err != nil {
		fmt.Printf("Register slave error: %v \n", err)
		return
	}

	b.SetRawMode(*rawMode)

	if *semiSync {
		if err := b.EnableSemiSync(); err != nil {
			fmt.Printf("Enable semi sync replication mode err: %v\n", err)
			return
		}
	}

	pos := mysql.Position{*file, uint32(*pos)}
	if len(*backupPath) > 0 {
		// must raw mode
		b.SetRawMode(true)

		err := b.StartBackup(*backupPath, pos, 0)
		if err != nil {
			fmt.Printf("Start backup error: %v\n", err)
			return
		}
	} else {
		s, err := b.StartSync(pos)
		if err != nil {
			fmt.Printf("Start sync error: %v\n", err)
			return
		}

		for {
			e, err := s.GetEvent()
			if err != nil {
				fmt.Printf("Get event error: %v\n", err)
				return
			}

			e.Dump(os.Stdout)
		}
	}

}