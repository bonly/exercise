package mypkg  

type Counter struct {  
	Value int
}

func (c *Counter) Inc() { 
    c.Value++;
}  

func New() *Counter { 
    return &Counter{ 5 };
}  
