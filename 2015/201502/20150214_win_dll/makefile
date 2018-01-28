st:
	GOGCCFLAGS="-fPIC" go build -buildmode=c-archive -o libhe.a 


lib2dllerr:
	gcc -m64 -shared -o he.dll mydll.def libhe.a -Wl,--allow-multiple-definition -static -lstdc++ -lwinmm -lntdll -lWs2_32

lib2dll:
	gcc -m64 -shared -o he.dll -Wl,--whole-archive libhe.a  -Wl,--allow-multiple-definition

all:
	go build -buildmode=c-archive dll.go
	gcc -shared -pthread -o abc.dll dll.a
