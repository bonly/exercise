package main

/*
typedef void (*Callback) (char* buffer);
void bridge_callback(Callback fn, char *buffer){
  return fn(buffer);
}
*/
import "C"
