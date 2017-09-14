package main

// #include <jni.h>
import "C"

//export Java_Hello_print
func Java_Hello_print(*C.JNIEnv, C.jobject) {
        println("hello from Go")
}

func main() {} // dummy

/*
// Hello.java
class Hello {
        private native void print();
        public static void main(String[] args) {
                new Hello().print();
                System.err.println("Hello from Java");
        }
        static {
                System.loadLibrary("hello");
        }
}
*/

/*
//-I"${JAVA_HOME}/include" -I"${JAVA_HOME}/include/linux"
export CPATH=${JAVA_HOME}/include:${JAVA_HOME}/include/linux
go build -buildmode=c-shared -o libhello.so hello.go 
javac Hello.java
LD_LIBRARY_PATH=. java Hello
*/