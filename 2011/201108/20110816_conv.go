package main

import (
    "flag"
    "fmt"
    "strconv"
)

func Sqrt(x float64) float64 {
    var z, d float64 = x, 1
    for d > 1e-15 {
        z0 := z
        z = z - (z*z-x)/(2*z)
        d = z - z0
        if d < 0 {
            d = -d
        }
    }
    return z
}

func main() {
    flag.Parse()
    for _, v := range flag.Args() {
        //f, err := strconv.Atof64(v)
        f, err := strconv.ParseFloat(v, 64);
        if err != nil {
            fmt.Printf("Couldn't convert %q: %v\n", v, err)
            continue
        }
        fmt.Printf("Sqrt(%v) = %v\n", f, Sqrt(f))
    }
}