all:
	export CPATH=${JAVA_HOME}/include:${JAVA_HOME}/include/linux
	go build -buildmode=c-shared -o libhello.so hello.go
	javac Hello.java
	LD_LIBRARY_PATH=. java Hello