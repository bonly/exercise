go build -buildmode=c-archive -o client.a src/app/client.go
gcc -m64 -shared -o client.dll cli.def client.a -static -lstdc++ -lwinmm -lntdll -lWs2_32