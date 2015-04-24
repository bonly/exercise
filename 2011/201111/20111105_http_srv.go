package main

import (
    "fmt"
    "net/http"
    "sync"
)

type state struct {
    *sync.Mutex // inherits locking methods
    Vals map[string]string // map ids to values
}

var State = &state{&sync.Mutex{}, map[string]string{}}

func get(rw http.ResponseWriter, req *http.Request) {
    State.Lock()
    defer State.Unlock() // ensure the lock is removed after leaving the the function
    id := req.URL.Query().Get("id") // if you need other types, take a look at strconv package
    val := State.Vals[id]
    delete(State.Vals, id)
    rw.Write([]byte("got: " + val))
}

func post(rw http.ResponseWriter, req *http.Request) {
    State.Lock()
    defer State.Unlock()
    id := req.FormValue("id")
    State.Vals[id] = req.FormValue("val")
    rw.Write([]byte("go to http://localhost:8080/?id=42"))
}

var form = `<html>
    <body>
        <form action="/" method="POST">
            ID: <input name="id" value="42" /><br />
            Val: <input name="val" /><br />
            <input type="submit" value="submit"/>
        </form>
    </body>
</html>`

func formHandler(rw http.ResponseWriter, req *http.Request) {
    rw.Write([]byte(form))
}

// for real routing take a look at gorilla/mux package
func handler(rw http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case "POST":
        post(rw, req)
    case "GET":
        if req.URL.String() == "/form" {
            formHandler(rw, req)
            return
        }
        get(rw, req)
    }
}

func main() {
    fmt.Println("go to http://localhost:8080/form")
    // thats the default webserver of the net/http package, but you may
    // create custom servers as well
    err := http.ListenAndServe("localhost:8080", http.HandlerFunc(handler))
    if err != nil {
        fmt.Println(err)
    }
}