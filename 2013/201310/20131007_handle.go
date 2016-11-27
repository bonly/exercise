package main

import (
    "encoding/json"
    "io"
    "net/http"
)

const MaxBodySize = 1048576

func unmarshal(req *http.Request, dst interface{}) error {
    defer req.Body.Close()
    return json.NewDecoder(io.LimitReader(req.Body, MaxBodySize)).Decode(dst)
}

func handle(route string, handler func(http.ResponseWriter, *http.Request) interface{}) {
    http.HandleFunc(route, func(res http.ResponseWriter, req *http.Request) {
        obj := handler(res, req)
        if err, ok := obj.(error); ok {
            http.Error(res, err.Error(), 500)
            return
        }
        json.NewEncoder(res).Encode(obj)
    })
}

func main() {
    handle("/example", func(res http.ResponseWriter, req *http.Request) interface{} {
        var apiRequest struct {
            X, Y int
        }
        err := unmarshal(req, &apiRequest)
        if err != nil {
            return err
        }
        // do stuff
        return "whatever you want"
    })

    http.ListenAndServe(":8080", nil)
}