package main

import (
	kcp "github.com/xtaci/kcp-go"
	"time"
	"golang.org/x/crypto/pbkdf2"
	"crypto/sha1"
)

var key = []byte("testkey");

func main(){
	pass := pbkdf2.Key(key, []byte("mysalt"), 4096, 32, sha1.New);
	block, _ := kcp.NewSalsa20BlockCrypt(pass);
	

	lst, err := kcp.ListenWithOptions("127.0.0.1:9999", block, 10, 3);
	// lst, err := kcp.Listen("127.0.0.1:9999");
	if err != nil{
		panic(err);
	}

	// go func(){
		// klst := lst;//.(*(net.Listener));
		lst.SetReadBuffer(4 * 1024 * 1024);
		lst.SetWriteBuffer(4 * 1024 * 1024);
		lst.SetDSCP(46);
		for {
			sess, err := lst.Accept();
			if err != nil{
				return;
			}
			sess.(*kcp.UDPSession).SetReadBuffer(4 * 1024 * 1024);
			sess.(*kcp.UDPSession).SetWriteBuffer(4 * 1024 * 1024);
			go handleClient(sess.(*kcp.UDPSession));
		}
	// }();
}

func handleClient(conn *kcp.UDPSession){
	conn.SetStreamMode(true)
	conn.SetWindowSize(4096, 4096)
	conn.SetNoDelay(1, 10, 2, 1)
	conn.SetDSCP(46)
	conn.SetMtu(1400)
	conn.SetACKNoDelay(false)
	conn.SetReadDeadline(time.Now().Add(time.Hour))
	conn.SetWriteDeadline(time.Now().Add(time.Hour))
	buf := make([]byte, 65536)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		conn.Write(buf[:n])
	}
}
/*
func main(){
	lis, err := kcp.ListenWithOptions(":10000", nil, 10, 3);
	if err != nil{
		log.Printf("%v\n", err);
		return;
	}
	for {
		if conn, err := lis.AcceptKCP(); err == nil{
			log.Printf("get connect from: %v\n", conn.RemoteAddr());
			go handleMux(conn, "127.0.0.1:9999");
		}else{
			log.Printf("%v\n", err);
		}
	}
}

//多路复用
func handleMux(conn io.ReadWriteCloser, target string){
	p1 := conn;
	p2, err := net.DialTimeout("tcp", "127.0.0.1:9999", 5*time.Second);
	if err != nil{
		log.Printf("%v\n", err);
		return;
	}
	go handleClient(p1, p2);
}

//客户端处理
func handleClient(p1 io.ReadWriteCloser, p2 io.ReadWriteCloser){
	defer p1.Close();
	defer p2.Close();

	// start tunnel
	p1die := make(chan struct{})
	go func() {
		io.Copy(p1, p2)
		close(p1die)
	}()

	p2die := make(chan struct{})
	go func() {
		io.Copy(p2, p1)
		close(p2die)
	}()

	// wait for tunnel termination
	select {
	case <-p1die:
	case <-p2die:
	}	
}
*/
