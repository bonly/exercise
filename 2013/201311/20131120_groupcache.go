package main
/*
把Key+时间作为键值，
时间就是每次更新缓存的计算标准
*/
import (
    "flag"
    "fmt"
    "github.com/golang/groupcache"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strconv"
    "strings"
)

var (
    // peers_addrs = []string{"127.0.0.1:8001", "127.0.0.1:8002", "127.0.0.1:8003"}
    //rpc_addrs = []string{"127.0.0.1:9001", "127.0.0.1:9002", "127.0.0.1:9003"}
    index = flag.Int("index", 0, "peer index")
)

func main() {
    flag.Parse()
    peers_addrs := make([]string, 3)
    rpc_addrs := make([]string, 3)
    if len(os.Args) > 0 {
        for i := 1; i < 4; i++ {
            peers_addrs[i-1] = os.Args[i]
            rpcaddr := strings.Split(os.Args[i], ":")[1]
            port, _ := strconv.Atoi(rpcaddr)
            rpc_addrs[i-1] = ":" + strconv.Itoa(port+1000)
        }
    }
    if *index < 0 || *index >= len(peers_addrs) {
        fmt.Printf("peer_index %d not invalid\n", *index)
        os.Exit(1)
    }
    peers := groupcache.NewHTTPPool(addrToURL(peers_addrs[*index]))
    var stringcache = groupcache.NewGroup("SlowDBCache", 64<<20, groupcache.GetterFunc(
        func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
            result, err := ioutil.ReadFile(key)
            if err != nil {
                log.Fatal(err)
                return err
            }
            fmt.Printf("asking for %s from dbserver\n", key)
            dest.SetBytes([]byte(result))
            return nil
        }))

    peers.Set(addrsToURLs(peers_addrs)...)

    http.HandleFunc("/zk", func(rw http.ResponseWriter, r *http.Request) {
        log.Println(r.URL.Query().Get("key"))
        var data []byte
        k := r.URL.Query().Get("key")
        fmt.Printf("cli asked for %s from groupcache\n", k)
        stringcache.Get(nil, k, groupcache.AllocatingByteSliceSink(&data))
        rw.Write([]byte(data))
    })
    go http.ListenAndServe(rpc_addrs[*index], nil)
    rpcaddr := strings.Split(os.Args[1], ":")[1]
    log.Fatal(http.ListenAndServe(":"+rpcaddr, peers))
}

func addrToURL(addr string) string {
    return "http://" + addr
}

func addrsToURLs(addrs []string) []string {
    result := make([]string, 0)
    for _, addr := range addrs {
        result = append(result, addrToURL(addr))
    }
    return result
}

/*
1、groupcache介绍： http://www.csdn.net/article/2013-07-30/2816399-groupcache-readme-go

2、groupcache常用框架：

    



    一般常用以上的框架去使用groupcache，此框架以及框架示例代码可通过https://github.com/capotej/groupcache-db-experiment 下载， 框架示例核心源码解读：

func main() {

    var port = flag.String("port", "8001", "groupcache port")
    flag.Parse()

    peers := groupcache.NewHTTPPool("http://localhost:" + *port)

    client := new(client.Client)

    var stringcache = groupcache.NewGroup("SlowDBCache", 64<<20, groupcache.GetterFunc(
        func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
            result := client.Get(key)
            fmt.Printf("asking for %s from dbserver\n", key)
            dest.SetBytes([]byte(result))
            return nil
        }))

    peers.Set("http://localhost:8001", "http://localhost:8002", "http://localhost:8003")

    frontendServer := NewServer(stringcache)

    i, err := strconv.Atoi(*port)
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
    var frontEndport = ":" + strconv.Itoa(i+1000)
    go frontendServer.Start(frontEndport)

    fmt.Println(stringcache)
    fmt.Println("cachegroup slave starting on " + *port)
    fmt.Println("frontend starting on " + frontEndport)
    http.ListenAndServe("127.0.0.1:"+*port, http.HandlerFunc(peers.ServeHTTP))
}
   理解以上这段代码需要首先理解groupcache中的peer如何与HttpPool产生关联，关键代码段：

func NewHTTPPoolOpts(self string, o *HTTPPoolOptions) *HTTPPool {
    if httpPoolMade {
        panic("groupcache: NewHTTPPool must be called only once")
    }
    httpPoolMade = true

    opts := HTTPPoolOptions{}
    if o != nil {
        opts = *o
    }
    if opts.BasePath == "" {
        opts.BasePath = defaultBasePath
    }
    if opts.Replicas == 0 {
        opts.Replicas = defaultReplicas
    }

    p := &HTTPPool{
        basePath:    opts.BasePath,
        self:        self,
        peers:       consistenthash.New(opts.Replicas, opts.HashFn),
        httpGetters: make(map[string]*httpGetter),
    }
    RegisterPeerPicker(func() PeerPicker { return p })
    return p
}
   通过RegisterPeerPicker将获取httppool的对象返回函数注册到全局的portPicker，这样在调用Group的Get接口时，通过调用initPeers接口返回HTTPPool对象，HTTPPool与groupcache的关联就是通过portPicker函数变量；

3、groupcache源码解读

    A、在使用以上框架的时候或许你会困惑，通过key分片，然后去远端获取数据时，在远端仍然是通过调用HTTPPool的ServeHttp来进行处理，我们先来看下该接口代码实现：

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Parse request.
    if !strings.HasPrefix(r.URL.Path, p.basePath) {
        panic("HTTPPool serving unexpected path: " + r.URL.Path)
    }
    parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
    if len(parts) != 2 {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    groupName := parts[0]
    key := parts[1]

    // Fetch the value for this group/key.
    group := GetGroup(groupName)
    if group == nil {
        http.Error(w, "no such group: "+groupName, http.StatusNotFound)
        return
    }
    var ctx Context
    if p.Context != nil {
        ctx = p.Context(r)
    }

    group.Stats.ServerRequests.Add(1)
    var value []byte
    err := group.Get(ctx, key, AllocatingByteSliceSink(&value))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Write the value to the response body as a proto message.
    body, err := proto.Marshal(&pb.GetResponse{Value: value})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/x-protobuf")
    w.Write(body)
}
    首先通过GetGroup获取本地的group对象指针，然后group.Get(ctx, key, AllocatingByteSliceSink(&value))获取数据，而在Frontend中也是通过调用Get接口获取数据，这样会不会形成死循环 ？ 为解答这一问题，首先我们来看下groupcache中的Get接口：

func (g *Group) Get(ctx Context, key string, dest Sink) error {
    g.peersOnce.Do(g.initPeers)
    g.Stats.Gets.Add(1)
    if dest == nil {
        return errors.New("groupcache: nil dest Sink")
    }
    value, cacheHit := g.lookupCache(key)

    if cacheHit {
        fmt.Printf("key %s cache hit!\n", key)
        g.Stats.CacheHits.Add(1)
        return setSinkView(dest, value)
    }

    // Optimization to avoid double unmarshalling or copying: keep
    // track of whether the dest was already populated. One caller
    // (if local) will set this; the losers will not. The common
    // case will likely be one caller.
    destPopulated := false
    value, destPopulated, err := g.load(ctx, key, dest)
    if err != nil {
        return err
    }
    if destPopulated {
        return nil
    }
    return setSinkView(dest, value)
}
    大概流程：先执行initPeers获取远端peer，查本地缓存是否有数据，如果命中，返回数据，否则load数据，看load实现：

func (g *Group) load(ctx Context, key string, dest Sink) (value ByteView, destPopulated bool, err error) {
    g.Stats.Loads.Add(1)
    viewi, err := g.loadGroup.Do(key, func() (interface{}, error) {
        g.Stats.LoadsDeduped.Add(1)
        var value ByteView
        var err error
        if peer, ok := g.peers.PickPeer(key); ok {
            value, err = g.getFromPeer(ctx, peer, key)
            if err == nil {
                g.Stats.PeerLoads.Add(1)
                return value, nil
            }
            g.Stats.PeerErrors.Add(1)
            // TODO(bradfitz): log the peer's error? keep
            // log of the past few for /groupcachez?  It's
            // probably boring (normal task movement), so not
            // worth logging I imagine.
        }
        value, err = g.getLocally(ctx, key, dest)
        if err != nil {
            g.Stats.LocalLoadErrs.Add(1)
            return nil, err
        }
        g.Stats.LocalLoads.Add(1)
        destPopulated = true // only one caller of load gets this return value
        g.populateCache(key, value, &g.mainCache)
        return value, nil
    })
    if err == nil {
        value = viewi.(ByteView)
    }
    return
}
    大概流程：根据key选择一个固定的远端peer，如果获取成功，那么从远端获取数据，否则getLocally直接从后端（数据库或者其他数据服务）获取数据；读到这里仍然无法解答这一疑惑，继续看PickPeer接口：

func (p *HTTPPool) PickPeer(key string) (ProtoGetter, bool) {
    p.mu.Lock()
    defer p.mu.Unlock()
    if p.peers.IsEmpty() {
        return nil, false
    }
    if peer := p.peers.Get(key); peer != p.self {
        return p.httpGetters[peer], true
    }
    return nil, false
}
    根据Key获取一个Peer，如果获取的peer是自己，那么认为失败，所以会选择从后端数据服务获取获取，并缓存在本地的maincache中，读到这里，困惑就消除了，一个key会选取固定的peer，所以如果已经定位到某个peer获取数据，peer再次调用Get接口时，如果lookupCache失败，那么就会调用getLocally尝试从数据服务获取，而不会循环从另外的peer去获取，形成死循环效应；

    B、LRU缓存算法为常用算法，一般采用list和hash数据结构结合实现，在此不再讲解；

    C、singleflight.go是为了保证当没有命中本地缓存是，同一个key在同一时刻只有一个在去remote peer或后端数据服务获取数据；

总结：

    groupcache精小而又强大，直接集成在自己的服务内部，推及使用！
*/