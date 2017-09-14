// +build darwin linux

package main

import "log"

import "golang.org/x/mobile/app" 
import "golang.org/x/mobile/event/lifecycle" 
import "golang.org/x/mobile/event/paint"

func main() { 
    app.Main(func(a app.App) { 
        // each posible lyfecycle stage 
        for e := range a.Events() { 
            // when an events come's 
            switch e := a.Filter(e).(type) { 
            case lifecycle.Event: 
                log.Println(e.From, "->", e.To) 
            case paint.Event: 
                // if it's a paint event, log it 
                if e.External { 
                    log.Println("paint external") 
                } else { 
                    log.Println("paint internal") 
                } 
                a.Publish() 
            } 
        } 
    }) 
} 

