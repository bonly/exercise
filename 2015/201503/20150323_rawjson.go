package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const input = `
{
	"type": 1,
	"msg": {
		"description": "dynamite",
		"authority": "the Bruce Dickinson"
	}
}`

// type Kind int

// const (
// 	sound Kind = iota
// 	cowbell
// )

type Envelope struct {
	Type int;//Kind
	Msg  interface{}
}

type App struct {
	// whatever your application state is
}

// Action is something that can operate on the application.
type Action interface {
	Run(app *App) error
}

type CowbellMsg struct {
	// ...
}

func (m *CowbellMsg) Run(app *App) error {
	// ...
	fmt.Println("in CowbellMsg");
	return nil;
}

type SoundMsg struct {
	// ...
}

func (m *SoundMsg) Run(app *App) error {
	// ...
	fmt.Println("in SoundMsg");
	return nil;
}

// var kindHandlers = map[Kind]func() Action{
// 	sound:   func() Action { return &SoundMsg{} },
// 	cowbell: func() Action { return &CowbellMsg{} },
// }

var kindHandlers = map[int]func() Action{
	1:   func() Action { return &SoundMsg{} },
	2: func() Action { return &CowbellMsg{} },
}

func main() {
	app := &App{
		// ...
	}

	// process an incoming message
	var raw json.RawMessage
	env := Envelope{
		Msg: &raw,
	}
	if err := json.Unmarshal([]byte(input), &env); err != nil {
		log.Fatal(err)
	}
	msg := kindHandlers[env.Type]()
	fmt.Printf("msg: %#v \n", msg);
	if err := json.Unmarshal(raw, msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("msg: %#v \n", msg);
	if err := msg.Run(app); err != nil {
		// ...
		fmt.Println("run");
	}
}