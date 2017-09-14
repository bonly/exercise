package main

import (
"bytes"
"golang.org/x/crypto/ssh"
"fmt"
)

func main() {
	// An SSH client is represented with a ClientConn. Currently only
	// the "password" authentication method is supported.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig.
	config := &ssh.ClientConfig{
	    User: "root",
	    Auth: []ssh.AuthMethod{
	        ssh.Password("hayes"),
	    },
	}
	client, err := ssh.Dial("tcp", "192.168.1.1:22", config)
	if err != nil {
	    panic("Failed to dial: " + err.Error())
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
	    panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/bin/uname"); err != nil {
	    panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
	}
