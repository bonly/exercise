//author mucjgm 16-02-24  jieshidonglin@163.com
package main

import (
	// "bufio"
	"fmt"
	"log"
	"net"
	// "time"
)

const (
	godMonster   = 1
	monsterKill  = 2
	godGuard     = 3
	guardGuard   = 4
	godWitch     = 5
	witchSave    = 6
	witchKill    = 7
	godSeer      = 8
	seerSeer     = 9
	godLight     = 10
	godChooseCop = 11
	godLoopSpeak = 12
)

type actor struct {
	c     net.Conn
	staus string
}

var Actor = make(map[int]actor)
var root = 0

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go activeActor(connection)
	}
}

func getID() int {
	root++
	return root
}

func activeActor(c net.Conn) {
	var newActor actor
	newActor.c = c
	var ID = getID()
	// input := bufio.NewScanner(c)
	fmt.Printf("ID %d player join!\n", ID)
	Actor[ID] = newActor
	// var judge bool
	for {
	  buf :=make([]byte,1024);
		len, _ := c.Read(buf);
	  cli_say := string(buf[:len]);
		fmt.Printf("%+v\n",cli_say);
		switch cli_say{
		case "exit":
				fmt.Println("cli say exit");
				c.Write([]byte("ok, exit now"));
				return;
			case "ready":
				fmt.Println("cli say: ", cli_say);
				c.Write([]byte("ok, copy that"));
				break;
			default:
				// fmt.Println("nothing");
				c.Close();
				return;
		}
	}
	/*
	for input.Scan() {
		switch input.Text() {
		case "exit":
			delete(Actor, ID)
			fmt.Printf("ID %d player exit!\n", ID)
			c.Write([]byte("exit\n"));
			judge = true
		case "ready":
			newActor.staus = "ready"
			fmt.Printf("ID %d player ready!\n", ID)
			c.Write([]byte("ready"))
			judge = true
		}
		time.Sleep(1 * time.Second)
		if judge {
			break
		}
	}
	*/
	Actor[ID] = newActor
	if judgeGame() {
		fmt.Printf("All %d players ready! Game start!", len(Actor))
	}
}

func (a actor) String() string {
	return a.staus
}

func judgeGame() bool {
	if len(Actor) > 0 {
		for _, act := range Actor {
			if act.staus != "ready" {
				return false
			}
		}
		return true
	}
	return false
}

func God() {

}
