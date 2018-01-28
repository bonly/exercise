package main 

func Add(x, y int) int{
	return x + y;
}

/*
go build -buildmode=plugin -o myplugin.so myplugin.go
*/