package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

type Rw struct{}

func main() {
	client("bonly", "xbed111", "198.11.177.244:22")
}
func client(user, passwd, ip string) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd),
		},
	}
	client, err := ssh.Dial("tcp", ip, config)
	if err != nil {
		log.Println("建立连接：", err)
		return
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		log.Println("建立会话：", err)
		return
	}
	defer session.Close()
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Printf("创建终端: %s", err)
		return
	}
	if err := session.Shell(); err != nil {
		log.Printf("执行Shell: %s", err)
		return
	}
	err = session.Wait()
	if err != nil {
		log.Println(err)
	}
}