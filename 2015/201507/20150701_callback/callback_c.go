package main

/*
#include <stddef.h>
#include <stdio.h>
typedef void (*GOLOG_PROC) (char *szDscr);
GOLOG_PROC g_logcallback = NULL;

void test(){
	if (g_logcallback == NULL){
		printf("it is null\n");
	}
	g_logcallback("");
}

void setcallback(GOLOG_PROC cb){
	g_logcallback = cb;
}

void bridge(char *sz){
	return g_logcallback(sz);
}
*/
import "C"
