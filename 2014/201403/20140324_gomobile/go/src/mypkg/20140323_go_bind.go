package mypkg  

import "fmt"

type Counter struct {  
	Value int
}

func (c *Counter) Inc() { 
    c.Value++;
}  

func New() *Counter { 
    return &Counter{ 5 };
}  

func Greetings(name string) string{
	return fmt.Sprintf("H %s", name);
}

/*
deven
GP
gomobild bind -target=android -o mypkg.aar mypkg 
*/
