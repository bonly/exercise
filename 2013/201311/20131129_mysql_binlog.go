package main
import (
    "github.com/siddontang/go-mysql/replication"
    "github.com/siddontang/go-mysql/mysql"
    "os"
    "fmt"
)

func main(){
    binlogFile := "/home/opt/maria/mydb/mysql-bin";
    binlogPos := uint32(313);
    // Create a binlog syncer with a unique server id, the server id must be different from other MySQL's. 
    // flavor is mysql or mariadb
    syncer := replication.NewBinlogSyncer(100, "mysql")

    // Register slave, the MySQL master is at 127.0.0.1:3306, with user root and an empty password
    syncer.RegisterSlave("127.0.0.1", 3306, "bonly", "")

    // Start sync with sepcified binlog file and position
    streamer, _ := syncer.StartSync(mysql.Position{binlogFile, binlogPos})

    // or you can start a gtid replication like
    // streamer, _ := syncer.StartSyncGTID(gtidSet)
    // the mysql GTID set likes this "de278ad0-2106-11e4-9f8e-6edd0ca20947:1-2"
    // the mariadb GTID set likes this "0-1-100"

    for {
        ev, err := streamer.GetEvent()
        if err != nil{
            fmt.Println(err);
            return;
        }
        // Dump event
        ev.Dump(os.Stdout)
    }
}
/*
//or use timeout with GetEventTimeout, and you should deal with the timeout exception 
import (
    "time"
) 
for {
    // timeout value won't be set too large, otherwise it may waste lots of memory
    ev, _ := streamer.GetEventTimeout(time.Second * 1)
    // Dump event
    ev.Dump(os.Stdout)
}
*/