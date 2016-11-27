package main

import (
"fmt"
"github.com/golang/groupcache"
)

func main(){
	peers := groupcache.NewHTTPPool("http://127.0.0.1");

	peers.Set("http://127.0.0.1", "http://120.25.106.243");

	var gp *groupcache.Group;
	{
		gp = groupcache.NewGroup("tes", 64<<20, 
			groupcache.GetterFunc(func(ctx groupcache.Context, key string, dest groupcache.Sink)error{
				fn := key;
				fmt.Println("want ", fn);
				dest.SetBytes([]byte("ok"));
				return nil;
			}));
		fmt.Println("stat: ", gp.Name());
	}
	{
		gp = groupcache.NewGroup("abc", 64<<20, 
			groupcache.GetterFunc(func(ctx groupcache.Context, key string, dest groupcache.Sink)error{
				fn := key;
				fmt.Println("want ", fn);
				dest.SetBytes([]byte("ok"));
				return nil;
			}));
		fmt.Println("stat: ", gp.Name());
	}	

	fmt.Println("stat: ", gp.Name());

	fmt.Println("last: ", groupcache.GetGroup("tes").Name());
}

/*
http://studygolang.com/articles/1345
GroupCache源码分析

概述

GitHub: https://github.com/golang/groupcache.git
memcached作者Brad Fitzpatrick用Go语言实现
分布式缓存库
数据无版本概念, 也就是一个key对应的value是不变的，并没有update
节点之间可互访，自动复制热点数据
实现
LRU

type Key interface{}

type entry struct {
    key    Key
    value interface{}  
}       

type Cache Struct {
    MaxEntries int
    OnEvicted func(key Key, value interface{})
    ll    *list.List
    cache map[interface{}]*list.Element
}
不懂Go语言或者对Go语言不熟悉的同学，可以和我一起来顺带学习Go语言 首先是Go语言将变量的类型放置于变量名之后上面的key变量，
它的类型是Key，然后需要说明的是interface{}，在Go中它可以指向任意对象，也就任意对象都是先这个接口，你可以把它看成Java/C#中的
Object, C/C++中的void*, 然后我们来看Cache，其一个字段MaxEntires表示Cache中最大可以有的KV数，OnEvicted是一个回调函数，
在触发淘汰时调用，ll是一个链表，cache是一个map，key为interface{}， value为list.Element其指向的类型是entry类型，通过数据
结构其实我们已经能够猜出LRU的实现了，没有错实现就是最基本的，Get一次放到list头，每次Add的时候判断是否满，满则淘汰掉list尾的数据

// 创建一个Cache
func New(maxEntries int) *Cache 
// 向Cache中插入一个KV
func (c *Cache) Add(key Key, value interface{})
// 从Cache中获取一个key对应的value
func (c *Cache) Get(key Key) (value interface{}, ok bool)
// 从Cache中删除一个key
func (c *Cache) Remove(key Key)
// 从Cache中删除最久未被访问的数据
func (c *Cache) RemoveOldest()
// 获取当前Cache中的元素个数
func (c *Cache) Len()
这里我们看到了Go中如何实现一个类的成员方法，在func之后加类，这种写法和Go语言很意思的一个东西有关系，在Go中接口的实现并不是和Java中那样
子，而是只要某个类只要实现了某个接口的所有方法，即可认为该类实现了该接口，类似的比如说在java中有2个接口名字不同，即使方法相同也是不一样
的，而Go里面认为这是一样的。另外Go中开头字母大写的变量名，类名，方法名表示对外可知，小写开头的表示对外不暴露。另外类实这种代码
ele.Value.(*entry).value，其中(*entry)表示将Value转成*entry类型访问

Singleflight

这个包主要实现了一个可合并操作的接口，代码也没几行

// 回调函数接口
type call struct {
    // 可以将其认为这是类似java的CountdownLatch的东西
    wg  sync.WaitGroup
    // 回调函数
    val interface{}
    // error
    err error
}

//注意这个Group是singleflight下的
type Group struct {
    // 保护m的锁
    mu sync.Mutex       // protects m
    // key/call 映射表
    m  map[string]*call // lazily initialized
}
另外有一点要说明的是Go使用开头字母大小写判断是否外部可见，大写外部可见，小写外部不可见，比如上面的Group在外部实用singleflght包
是可以访问到的，而call不能，真不理解为啥不用个export啥的关键字，唉

// 这个就是这个包的主要接口，用于向其他节点发送查询请求时，合并相同key的请求，减少热点可能带来的麻烦
// 比如说我请求key="123"的数据，在没有返回的时候又有很多相同key的请求，而此时后面的没有必要发，只要
// 等待第一次返回的结果即可
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
    // 首先获取group锁
    g.mu.Lock()
    // 映射表不存在就创建1个
    if g.m == nil {
        g.m = make(map[string]*call)
    }
    // 判断要查询某个key的请求是否已经在处理了
    if c, ok := g.m[key]; ok {
        // 已经在处理了 就释放group锁 并wait在wg上，我们看到wg加1是在某个时间段内第一次请求时加的
        // 并且在完成fn操作返回后会wg done，那时wait等待就会返回，直接返回第一次请求的结果
        g.mu.Unlock()
        c.wg.Wait()
        return c.val, c.err
    }
    // 创建回调，wg加1，把回调存到m中表示已经有在请求了，释放锁
    c := new(call)
    c.wg.Add(1)
    g.m[key] = c
    g.mu.Unlock()

    // 执行fn，释放wg
    c.val, c.err = fn()
    c.wg.Done()

    // 加锁将请求从m中删除，表示请求已经做好了
    g.mu.Lock()
    delete(g.m, key)
    g.mu.Unlock()

    return c.val, c.err
}
ByteView

一个不可变byte视图，其实就包装了一下byte数组和string，一个为null，就用另外一个

type ByteView struct {
    // If b is non-nil, b is used, else s is used.
    b []byte
    s string
}
提供的接口也很简单：

// 返回长度
func (v ByteView) Len() int
// 按[]byte返回一个拷贝
func (v ByteView) ByteSlice() []byte
// 按string返回一个拷贝
func (v ByteView) String() string {
// 返回第i个byte
func (v ByteView) At(i int) byte
// 返回ByteView的某个片断，不拷贝
func (v ByteView) Slice(from, to int) ByteView
// 返回ByteView的从某个位置开始的片断，不拷贝
func (v ByteView) SliceFrom(from int) ByteView
// 将ByteView按[]byte拷贝出来
func (v ByteView) Copy(dest []byte) int
// 判断2个ByteView是否相等
func (v ByteView) Equal(b2 ByteView) bool
// 判断ByteView是否和string相等
func (v ByteView) EqualString(s string) bool
// 判断ByteView是否和[]byte相等
func (v ByteView) EqualBytes(b2 []byte) bool
// 对ByteView创建一个io.ReadSeeker
func (v ByteView) Reader() io.ReadSeeker
// 读取从off开始的后面的数据，其实下面调用的SliceFrom，这是封装成了io.Reader的一个ReadAt方法的形式
func (v ByteView) ReadAt(p []byte, off int64) (n int, err error) {
Sink

不太好表达这个是较个啥，举个例子，比如说我调用一个方法返回需要是一个[]byte，而方法的实现我并不知道，可以它内部产生的是一个string，
方法也不知道调用的它的实需要是啥，于是我们就可以在调用方这边使用这边使用继承了Sink的类，而方法内部只要调用Sink的方法进行传值

type Sink interface {
    // SetString sets the value to s.
    SetString(s string) error

    // SetBytes sets the value to the contents of v.
    // The caller retains ownership of v.
    SetBytes(v []byte) error

    // SetProto sets the value to the encoded version of m.
    // The caller retains ownership of m.
    SetProto(m proto.Message) error

    // view returns a frozen view of the bytes for caching.
    view() (ByteView, error)
}

// 使用ByteView设置Sink，间接调用SetString，SetBytes等
func setSinkView(s Sink, v ByteView) error
这里要说明一点，Go语言比较有意思的是，类是否实现了某个接口是外挂式的，也就是不要去再类定义的时候写类似implenents之类的东西，只要类方法
中包括了所有接口的方法，就说明这个类是实现了某个接口

这里实现Sink的有5个类：

stringSink

返回值是string的Sink

type stringSink struct {
    sp *string
    v  ByteView
}

// 创建一个stringSink， 传入的是一个string的指针    
func StringSink(sp *string) Sink {
    return &stringSink{sp: sp}
}
byteViewSink

返回值是byteView的Sink

type byteViewSink struct {
    dst *ByteView
}

// 创建一个byteViewSink，传入的是一个ByteView指针
func ByteViewSink(dst *ByteView) Sink {
    if dst == nil {
        panic("nil dst")
    }
    return &byteViewSink{dst: dst}
}
protoSink

返回值是一个protobuffer message

type protoSink struct {
    dst proto.Message // authorative value
    typ string

    v ByteView // encoded
}

// 创建一个protoSink，传入的是一个proto.Messages    
func ProtoSink(m proto.Message) Sink {
    return &protoSink{
        dst: m,
    }
}
allocBytesSink

返回值是[]byte的Sink，每次Set都会重新分配内存

type allocBytesSink struct {
    dst *[]byte
    v   ByteView
}

// 创建一个allocBytesSink，传入的是一个数组分片指针   
func AllocatingByteSliceSink(dst *[]byte) Sink {
    return &allocBytesSink{dst: dst}
}
truncBytesSink

返回值是一个定长的[]byte的Sink，超过长度的会被截断，服用传入的空间

type truncBytesSink struct {
    dst *[]byte
    v   ByteView
}

// 创建一个truncBytesSink，传入的实一个数组分片指针    
func TruncatingByteSliceSink(dst *[]byte) Sink {
    return &truncBytesSink{dst: dst}
}
另外这些Sink类，都是包外不可见的，但是创建函数和Sink接口可见的，这样子在对于使用上来说，只有接口操作不会有另外的东西，对外简单清晰

HTTPPool

一个HTTPPool提供了groupcache节点之间的方式以及节点的选择的实现，但是其对于groupcache.Group是透明的，Group使用的HTTPPool
实现的1个接口PeerPicker以及httpGetter实现的接口ProtoGetter

// groupcache提供了一个节点互相访问访问的类
type httpGetter struct {
    transport func(Context) http.RoundTripper
    baseURL   string
}

// 协议为GET http://example.com/groupname/key
// response见groupcache.proto，含有2个可选项分别为[]byte和double
// 实现默认使用go自带的net/http包直接发送请求
func (h *httpGetter) Get(context Context, in *pb.GetRequest, out *pb.GetResponse) error
HttpPool实现比较简单都是直接使用Go内部自带的一些包

// 创建一个HttpPool, 只能被使用一次，主要是注册PeerPicker，以及初始化http服务
func NewHTTPPool(self string) *HTTPPool

// 设置groupcache集群的节点列表
func (p *HTTPPool) Set(peers ...string)

// 提供按key选取节点，按key作hash，但是这段代码在OS为32bit是存在bug，如果算出来的hashcode正好是-1 * 2^31时
// 会导致out of range，为啥会有这个bug看看代码你就会发现了，作者忘了-1 * 2^31 <= int32 <= 1 * 2^31 -1        
func (p *HTTPPool) PickPeer(key string) (ProtoGetter, bool)

// http服务处理函数，主要是按http://example.com/groupname/key解析请求，调用group.Get，按协议返回请求
func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request)
Cache

Cache的实现很简单基本可以认为就是直接在LRU上面包了一层，加上了统计信息，锁，以及大小限制

type cache struct {
    mu         sync.RWMutex
    nbytes     int64 // of all keys and values
    lru        *lru.Cache
    nhit, nget int64
    nevict     int64 // number of evictions
}

// 返回统计信息，加读锁
func (c *cache) stats() CacheStats
// 加入一个kv，如果cache的大小超过nbytes了，就淘汰
func (c *cache) add(key string, value ByteView)
// 根据key返回value
func (c *cache) get(key string) (value ByteView, ok bool)
// 删除最久没被访问的数据    
func (c *cache) removeOldest()
// 获取当前cache的大小
func (c *cache) bytes() int64
// 获取当前cache中kv的个数，内部会加读锁    
func (c *cache) items() int64
// 获取当前cache中kv的个数，函数名中的Locked的意思是调用方已经加锁，比如上面的stats方法    
func (c *cache) itemsLocked() int64
Group

该类是GroupCache的核心类，所有的其他包都是为该类服务的

type Group struct {
    // group名，可以理解为namespace的名字，group其实就是一个namespace
    name       string
    // 由调用方传入的回调，用于groupcache访问后端数据
    // type Getter interface {
    //     Get(ctx Context, key string, dest Sink) error
    // }
    getter     Getter
    // 通过该对象可以达到只调用一次
    peersOnce  sync.Once
    // 用于访问groupcache中其他节点的接口，比如上面的HTTPPool实现了该接口
    peers      PeerPicker
    // cache的总大小
    cacheBytes int64

    // 从本节点指向向后端获取到的数据存在在cache中
    mainCache cache

    // 从groupcache中其他节点上获取到的数据存在在该cache中，因为是用其他节点获取到的，也就是说这个数据存在多份，也就是所谓的hot
    hotCache cache
    // 用于向groupcache中其他节点访问时合并请求
    loadGroup singleflight.Group

    // 统计信息
    Stats Stats
}
由上面的结构体我们可以看出来，groupcache支持namespace概念，不同的namespace有自己的配额以及cache，不同group之间cache是独立的
也就是不能存在某个group的行为影响到另外一个namespace的情况 groupcache的每个节点的cache分为2层，由本节点直接访问后端的
存在maincache，其他存在在hotcache

// 根据name从全局变量groups（一个string/group的map）查找对于group
func GetGroup(name string) *Group
// 创建一个新的group，如果已经存在该name的group将会抛出一个异常，go里面叫panic
func NewGroup(name string, cacheBytes int64, getter Getter) *Group
// 注册一个创建新group的钩子，比如将group打印出来等，全局只能注册一次，多次注册会触发panic
func RegisterNewGroupHook(fn func(*Group))
// 注册一个服务，该服务在NewGroup开始时被调用，并且只被调用一次
func RegisterServerStart(fn func())
下面重点介绍Get方法

func (g *Group) Get(ctx Context, key string, dest Sink) error
整个groupcache核心方法就这么一个，我们来看一下该方法是怎么运作的

func (g *Group) Get(ctx Context, key string, dest Sink) error {
    // 初始化peers，全局初始化一次
    g.peersOnce.Do(g.initPeers)
    // 统计信息gets增1
    g.Stats.Gets.Add(1)
    if dest == nil {
        return errors.New("groupcache: nil dest Sink")
    }
    // 查找本地是否存在该key，先查maincache，再查hotcache
    value, cacheHit := g.lookupCache(key)

    // 如果查到九hit增1，并返回
    if cacheHit {
        g.Stats.CacheHits.Add(1)
        return setSinkView(dest, value)
    }

    // 加载该key到本地cache中，其中如果是本地直接请求后端得到的数据，并且是同一个时间段里第一个，
    // 就不需要重新setSinkView了，在load中已经设置过了，destPopulated这个参数以来底的实现
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

func (g *Group) load(ctx Context, key string, dest Sink) (value ByteView, destPopulated bool, err error) {
    // 统计loads增1
    g.Stats.Loads.Add(1)
    // 使用singleflight来达到合并请求
    viewi, err := g.loadGroup.Do(key, func() (interface{}, error) {
        // 统计信息真正发送请求的次数增1
        g.Stats.LoadsDeduped.Add(1)
        var value ByteView
        var err error
        // 选取向哪个节点发送请求，比如HTTPPool中的PickPeer实现
        if peer, ok := g.peers.PickPeer(key); ok {
            // 从groupcache中其他节点获取数据，并将数据存入hotcache
            value, err = g.getFromPeer(ctx, peer, key)
            if err == nil {
                g.Stats.PeerLoads.Add(1)
                return value, nil
            }
            g.Stats.PeerErrors.Add(1)
        }
        // 如果选取的节点就是本节点或从其他节点获取失败，则由本节点去获取数据，也就是调用getter接口的Get方法
        // 另外我们看到这里dest已经被赋值了，所以有destPopulated来表示已经赋值过不需要再赋值
        value, err = g.getLocally(ctx, key, dest)
        if err != nil {
            g.Stats.LocalLoadErrs.Add(1)
            return nil, err
        }
        g.Stats.LocalLoads.Add(1)
        destPopulated = true
        // 将获取的数据放入maincache
        g.populateCache(key, value, &g.mainCache)
        return value, nil
    })

    // 如果成功则返回
    if err == nil {
        value = viewi.(ByteView)
    }
    return
}    
这里需要说明的是A向groupcache中其他节点B发送请求，此时B是调用Get方法，然后如果本地不存在则也会走load，但是不同的是
PickPeer会发现是本身节点（HTTPPOOL的实现），然后就会走getLocally，会将数据在B的maincache中填充一份，也就是说如果
是向groupcache中其他节点发请求的，会一下子在groupcache集群内存2分数据，一份在B的maincache里，一份在A的hotcache中，
这也就达到了自动复制，越是热点的数据越是在集群内份数多，也就达到了解决热点数据的问题

总结
按Groupcache项目页面上的宣称的替代Memcached的大多数应用场景表示不大现实，GroupCache在我看来基本上只能适用于静态数据，比如下载服务中的文件缓存等
*/