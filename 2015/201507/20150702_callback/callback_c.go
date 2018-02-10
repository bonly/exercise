package main

/*
#include <stddef.h>
#include <stdio.h>
typedef void (*GOLOG_PROC) (char *szDscr);

void go_logcallback(char*);

GOLOG_PROC fn = go_logcallback;

void bridge(char *szDscr){
	return fn(szDscr);
}

void set_callback(void *cb){
	fn = cb;
}
*/
import "C"
