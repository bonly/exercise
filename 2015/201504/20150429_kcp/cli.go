package main

import(
	kcp "github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"
	"crypto/sha1"
	"time"
	"fmt"
)

var key = []byte("testkey");

func main(){
	pass := pbkdf2.Key(key, []byte("mysalt"), 4096, 32, sha1.New);	
	block, _ := kcp.NewSalsa20BlockCrypt(pass);
	sess, err := kcp.DialWithOptions("127.0.0.1:9999", block, 10, 3)
	// sess, err := kcp.Dial("127.0.0.1:9999");
	if err != nil {
		panic(err)
	}	

	sess.SetStreamMode(true)
	sess.SetStreamMode(false)
	sess.SetStreamMode(true)
	sess.SetWindowSize(4096, 4096)
	sess.SetReadBuffer(4 * 1024 * 1024)
	sess.SetWriteBuffer(4 * 1024 * 1024)
	sess.SetStreamMode(true)
	sess.SetNoDelay(1, 10, 2, 1)
	sess.SetMtu(1400)
	sess.SetMtu(1600)
	sess.SetMtu(1400)
	sess.SetACKNoDelay(true)
	sess.SetDeadline(time.Now().Add(time.Minute))	

	sess.SetWriteDelay(true);
	sess.SetDUP(1);
	const N = 100;
	buf := make([]byte, 10);
	for i := 0; i<N; i++{
		msg := fmt.Sprintf("hello%v", i);
		sess.Write([]byte(msg));
		if n, err := sess.Read(buf); err == nil{
			if string(buf[:n]) != msg{
				fmt.Printf("not the same");
			}else{
				fmt.Printf("recv: %s\n", buf[:n]);
			}
		}else{
			fmt.Printf("%v\n", err);
		}
	}
	sess.Close();
}
/*
func main(){
	lis, err := net.ListenTCP("tcp", "127.0.0.1:9999");
	if err != nil{
		log.Printf("%v\n", err);
		return;
	}

	// block, _ = kcp.NewNoneBlockCrypt()
	conn, err := kcp.DialWithOptions("127.0.0.1:10000", nil, 10, 3);

	//处理复用

	p1, err := list.AcceptTCP();
	// p2 = 
	go handleClient(p1, p2);
}

func handleClient(p1, p2 io.ReadWriteCloser) {
	log.Println("stream opened")
	defer log.Println("stream closed")
	defer p1.Close()
	defer p2.Close()

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