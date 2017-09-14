all:
	export CPATH=${JAVA_HOME}/include:${JAVA_HOME}/include/linux
	export CGO_ENABLED=0
	gcc -c -fPIC jstring.c -o jstring.o 
	# ar rcs libjstring.a jstring.o
	gcc -shared -o libjstring.so jstring.o
	go build -buildmode=c-shared -o libqunar.so crypt.go
	javac Qunar.java
	LD_LIBRARY_PATH=. java Qunar

win:
	#windows下不行，暂无解决方案
	export CPATH=${JAVA_HOME}/include:${JAVA_HOME}/include/linux
	export CGO_ENABLED=0
	gcc -c -fPIC jstring.c -o jstring.o 
	ar rcs libjstring.a jstring.o
	GOOS=windows GOARCH=amd64 go build -buildmode=c-shared -o qunar.dll crypt.go

clean:
	rm -rf lib* *.a *.class *.o *.dll