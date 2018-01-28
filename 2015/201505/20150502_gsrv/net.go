package main

import (
	"fmt"
	"log"
	kcp "github.com/xtaci/kcp-go"
	"time"
	"golang.org/x/crypto/pbkdf2"
	"crypto/sha1"
)

const addr = "127.0.0.1:4242";
var key = []byte("testkey");
var salt = []byte("techappen");

func Srv() error{
	log.Printf("srv begin\n");

	pass := pbkdf2.Key(key, salt, 4096, 32, sha1.New);
	block, _ := kcp.NewSalsa20BlockCrypt(pass);

	lst, err := kcp.ListenWithOptions(addr, block, 10, 3);
	if err != nil{
		return err;
	}

	lst.SetReadBuffer(4 * 1024 * 1024);
	lst.SetWriteBuffer(4 * 1024 * 1024);
	lst.SetDSCP(46);
	for {
		sess, err := lst.Accept();
		if err != nil{
			return err;
		}
		sess.(*kcp.UDPSession).SetReadBuffer(4 * 1024 * 1024);
		sess.(*kcp.UDPSession).SetWriteBuffer(4 * 1024 * 1024);

		fmt.Printf("get connect from %s\n", sess.RemoteAddr());
		go handleClient(sess.(*kcp.UDPSession));
	}

	return err;
}

func handleClient(conn *kcp.UDPSession){
	conn.SetStreamMode(true);
	conn.SetWindowSize(4096, 4096);
	conn.SetNoDelay(1, 10, 2, 1);
	conn.SetDSCP(46);
	conn.SetMtu(1400);
	conn.SetACKNoDelay(false);
	conn.SetReadDeadline(time.Now().Add(time.Hour));
	conn.SetWriteDeadline(time.Now().Add(time.Hour));

	buf := make([]byte, 65536);
	for {
		n, err := conn.Read(buf);
		if err != nil {
			fmt.Printf("perr err: %v\n", err);
			return;
		}
		conn.Write(buf[:n]);
	}
}
