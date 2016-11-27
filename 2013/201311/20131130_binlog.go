package main

import (
    "fmt"
    "github.com/2tvenom/myreplication"
    "github.com/golang/glog"
    "flag"
)

var (
    host     = "localhost"
    port     = 3306
    username = "bonly"
    password = ""
)

func main() {
    flag.Parse();
    flag.Set("alsologtostderr", "true");
    newConnection := myreplication.NewConnection()
    serverId := uint32(2)
    err := newConnection.ConnectAndAuth(host, port, username, password)

    if err != nil {
        panic("Client not connected and not autentificate to master server with error:" + err.Error())
    }
    //Get position and file name
    pos, filename, err := newConnection.GetMasterStatus()

    if err != nil {
        panic("Master status fail: " + err.Error())
    }

    el, err := newConnection.StartBinlogDump(pos, filename, serverId)

    if err != nil {
        panic("Cant start bin log: " + err.Error())
    }
    events := el.GetEventChan()
    go func() {
        for {
            event := <-events

            switch e := event.(type) {
            case *myreplication.QueryEvent:
                //Output query event
                glog.Info(e.GetQuery())
            case *myreplication.IntVarEvent:
                //Output last insert_id  if statement based replication
                glog.Info(e.GetValue())
            case *myreplication.WriteEvent:
                //Output Write (insert) event
                glog.Info("Write", e.GetTable())
                //Rows loop
                for i, row := range e.GetRows() {
                    //Columns loop
                    for j, col := range row {
                        //Output row number, column number, column type and column value
                        glog.Info(fmt.Sprintf("%d %d %d %v", i, j, col.GetType(), col.GetValue()))
                    }
                }
            case *myreplication.DeleteEvent:
                //Output delete event
                glog.Info("Delete", e.GetTable())
                for i, row := range e.GetRows() {
                    for j, col := range row {
                        glog.Info(fmt.Sprintf("%d %d %d %v", i, j, col.GetType(), col.GetValue()))
                    }
                }
            case *myreplication.UpdateEvent:
                //Output update event
                glog.Info("Update", e.GetTable())
                //Output old data before update
                for i, row := range e.GetRows() {
                    for j, col := range row {
                        glog.Info(fmt.Sprintf("%d %d %d %v", i, j, col.GetType(), col.GetValue()))
                    }
                }
                //Output new
                for i, row := range e.GetNewRows() {
                    for j, col := range row {
                        glog.Info(fmt.Sprintf("%d %d %d %v", i, j, col.GetType(), col.GetValue()))
                    }
                }
            default:
            }
        }
    }()
    err = el.Start()
    glog.Info(err.Error())
}