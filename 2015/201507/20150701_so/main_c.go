package main

/*
typedef void (*Callback)(unsigned int sn, char *buffer);
Callback fn = 0;
void SoCallCs(unsigned int sn, char *buffer){
	return fn(sn, buffer);
}
*/
import "C"
