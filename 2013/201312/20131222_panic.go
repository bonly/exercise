package main

func main() {
	panic("kerboom")
}

/*
% env GOTRACEBACK=0 ./crash 
panic: kerboom
% echo $?
2
*/
