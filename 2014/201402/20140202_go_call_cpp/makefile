all:
	g++ -o cplus.o -c cplus.cpp
	g++ -o c.o -c c.c
	ar r libc_test.so c.o cplus.o
	go build  call_cpp.go
