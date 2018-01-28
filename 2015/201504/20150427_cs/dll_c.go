package main

/*
//两个go都分别做了typedef却不冲突？！
typedef void (*Callback)(unsigned int sn, char *buffer);
Callback fn;
void bridge_callback(unsigned int sn, char *buffer){
	return fn(sn, buffer);
}
*/
import "C"