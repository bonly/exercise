package main

import (
    "encoding/json"
    "fmt"
    "github.com/codegangsta/martini"
    "github.com/golang/groupcache"
    "github.com/hashicorp/memberlist"
    "log"
    "net/http"
    "os"
    "strings"
    "time"
)

const GroupcachePort = 8000

type eventDelegate struct {
    peers []string
    pool  *groupcache.HTTPPool
}

func (e *eventDelegate) NotifyJoin(node *memberlist.Node) {
    uri := e.groupcacheURI(node.Addr.String())
    e.removePeer(uri)
    e.peers = append(e.peers, uri)
    if e.pool != nil {
        e.pool.Set(e.peers...)
    }
    log.Print("Add peer: " + uri)
    log.Printf("Current peers: %v", e.peers)
}

func (e *eventDelegate) NotifyLeave(node *memberlist.Node) {
    uri := e.groupcacheURI(node.Addr.String())
    e.removePeer(uri)
    e.pool.Set(e.peers...)
    log.Print("Remove peer: " + uri)
    log.Printf("Current peers: %v", e.peers)
}

func (e *eventDelegate) NotifyUpdate(node *memberlist.Node) {
    log.Print("Update the node: %+v\n", node)
}

func (e *eventDelegate) groupcacheURI(addr string) string {
    return fmt.Sprintf("http://%s:%d", addr, GroupcachePort)
}

func (e *eventDelegate) removePeer(uri string) {
    for i := 0; i < len(e.peers); i++ {
        if e.peers[i] == uri {
            e.peers = append(e.peers[:i], e.peers[i+1:]...)
            i--
        }
    }
}

func initGroupCache() {
    eventHandler := &eventDelegate{}
    conf := memberlist.DefaultLANConfig()
    conf.Events = eventHandler
    if addr := os.Getenv("GROUPCACHE_ADDR"); addr != "" {
        conf.AdvertiseAddr = addr
    }

    list, err := memberlist.Create(conf)
    if err != nil {
        panic("Failed to created memberlist: " + err.Error())
    }

    self := list.Members()[0]
    addr := fmt.Sprintf("%s:%d", self.Addr, GroupcachePort)
    eventHandler.pool = groupcache.NewHTTPPool("http://" + addr)
    go http.ListenAndServe(addr, eventHandler.pool)

    if nodes := os.Getenv("JOIN_TO"); nodes != "" {
        if _, err := list.Join(strings.Split(nodes, ",")); err != nil {
            panic("Failed to join cluster: " + err.Error())
        }
    }
}

func main() {
    initGroupCache()
    heavy := groupcache.NewGroup("heavy", 64<<20, groupcache.GetterFunc(heavyTask))

    m := martini.Classic()
    m.Get("/_stats", func() []byte {
        v, err := json.Marshal(&heavy.Stats)
        if err != nil {
            panic(err)
        }
        return v
    })
    m.Get("/:key", func(params martini.Params) string {
        var result string
        if err := heavy.Get(nil, params["key"], groupcache.StringSink(&result)); err != nil {
            panic(err)
        }
        return result
    })
    m.Run()
}

func heavyTask(ctx groupcache.Context, key string, dst groupcache.Sink) error {
    time.Sleep(400 * time.Millisecond)
    dst.SetString("Value of " + key)
    return nil
}