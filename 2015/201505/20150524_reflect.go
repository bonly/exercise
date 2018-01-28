package main

import (
    "errors"
    "fmt"
    "reflect"   
)

func InterfaceMap(i interface{}) (interface{}, error) {
    // Get type
    t := reflect.TypeOf(i)

    switch t.Kind() {
    case reflect.Map:
        // Get the value of the provided map
        v := reflect.ValueOf(i)

        // The "only" way of making a reflect.Type with interface{}
        it := reflect.TypeOf((*interface{})(nil)).Elem()

        // Create the map of the specific type. Key type is t.Key(), and element type is it
        m := reflect.MakeMap(reflect.MapOf(t.Key(), it))

        // Copy values to new map
        for _, mk := range v.MapKeys() {            
            m.SetMapIndex(mk, v.MapIndex(mk))
        }

        return m.Interface(), nil

    }

    return nil, errors.New("Unsupported type")
}

func main() {
    foo := make(map[string]int)
    foo["anisus"] = 42

    bar, err := InterfaceMap(foo)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%#v\n", bar.(map[string]interface{}))
}