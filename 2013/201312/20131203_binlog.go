package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/garyburd/redigo/redis"
	// "iconv"
)

var host = flag.String("host", "127.0.0.1", "MySQL host")
// var host = flag.String("host", "120.25.106.243", "MySQL host")
var port = flag.Int("port", 3306, "MySQL port")
var user = flag.String("user", "root", "MySQL user, must have replication privilege")
// var user = flag.String("user", "db_admin", "MySQL user, must have replication privilege")
var password = flag.String("password", "", "MySQL password")
// var password = flag.String("password", "db_admin2015", "MySQL password")

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

		cli, err := redis.Dial("tcp", "192.168.1.13:6379");
		if  err != nil{
			fmt.Println("connect fail: ", err);
			return;
		}
		defer cli.Close();	

		for {
			e, err := s.GetEvent()
			if err != nil {
				fmt.Printf("Get event error: %v\n", err)
				return
			}

			

			if e.Header.EventType == replication.INTVAR_EVENT{
				fmt.Println("==============INTVAR_EVENT==================");
				// //写入
				// if _, err = cli.Do("SET", os.Args[1], os.Args[2]); err != nil{
				// 	fmt.Println("set fail: ", err);
				// }	
				// e.Dump(os.Stdout)
			}
			e.Dump(os.Stdout)
		}
	}

}