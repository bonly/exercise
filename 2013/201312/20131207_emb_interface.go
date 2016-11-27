package main

import (
    "fmt"
    "sort"
    "strconv"
)

// An Interface "composed" of 5 methods: sort (3 methods) and 2 more
type MinMax interface {
    sort.Interface //embedding the sort.Interface
    Copy() MinMax
    Get(i int) interface{}
}
type Person struct {
    name  string
    age   int
    phone string
}

// Person with the following methods become a Stringer! see usage below
func (h Person) String() string {
    return "[name: " + h.name + " - age: " + strconv.Itoa(h.age) + " years]"
}

// Our collection
type People []Person // People is a type of slices that contain Persons
// is sortable as it implements the 3 following methods.
func (g People) Len() int {
    return len(g)
}
func (g People) Less(i, j int) bool {
    // sortable on age, and from younger to older
    if g[i].age < g[j].age {
        return true
    }
    return false
}
func (g People) Swap(i, j int) {
    g[i], g[j] = g[j], g[i]
}

// People satisfies the MinMax Interface too as it implements also these 3
// methods
func (g People) Get(i int) interface{} { return g[i] }
func (g People) Copy() MinMax {
    c := make(People, len(g))
    copy(c, g)
    return c
}
func main() {
    // initializing a group of People
    group := People{
        Person{name: "Bart", age: 24},
        Person{name: "Bob", age: 23},
        Person{name: "Gertrude", age: 104},
        Person{name: "Paul", age: 44},
        Person{name: "Sam", age: 34},
        Person{name: "Jack", age: 54},
        Person{name: "Martha", age: 74},
        Person{name: "Leo", age: 4},
    }
    // Let's print each member of this group
    fmt.Println("The unsorted group is:")
    for _, value := range group {
        // each value is a person... a Stringer
        fmt.Println(value)
    }
    // Now let's get the older and the younger
    younger, older := GetMinMax(group)
    fmt.Println("\n➞ Younger is", younger)
    fmt.Println("➞ Older is ", older)
    // uncomment to sort the group
    // sort.Sort(group)
    // Let's print this group again, uncomment the following 4 lines
    // fmt.Println("\nThe original group is still:")
    // for _, value := range group {
    // fmt.Println(value)
    // }
}

// Where we take advantage of the Interface MinMax
func GetMinMax(C MinMax) (min, max interface{}) {
    K := C.Copy() //returns a MinMax
    // a MinMax type is sortable and has a Get method, so K is sortable.
    sort.Sort(K)
    // youngest should be the first, oldest the last
    min, max = K.Get(0), K.Get(K.Len()-1)
    return
}

/* Expected Result:
The unsorted group is:
[name: Bart - age: 24 years]
[name: Bob - age: 23 years]
[name: Gertrude - age: 104 years]
[name: Paul - age: 44 years]
[name: Sam - age: 34 years]
[name: Jack - age: 54 years]
[name: Martha - age: 74 years]
[name: Leo - age: 4 years]
➞ Younger is [name: Leo - age: 4 years]
➞ Older is  [name: Gertrude - age: 104 years]
*/