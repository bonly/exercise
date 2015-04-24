package main

import("github.com/kellydunn/golang-geo"
       "fmt")

func main() {
     // Make a few points
     p := geo.NewPoint(42.25, 120.2)
     p2 := geo.NewPoint(30.25, 112.2)

     // find the great circle distance between them
     dist := p.GreatCircleDistance(p2)
     fmt.Printf("great circle distance: %d\n", dist)
}