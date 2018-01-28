@echo off

echo %1
if {%1} == {} goto All
if {%1} == {env} goto Env
if {%1} == {pkg} goto Pkg
goto end

:Env
setx GOPATH %CD%;%GOPATH% /m
@echo GOPATH=%GOPATH%
goto end

:All
go build -buildmode=c-archive -o client.a src/app/client.go
gcc -m64 -shared -o libtechappen.so cli.def client.a -static -lstdc++ -lwinmm -lntdll -lWs2_32
goto end

:Pkg
mkdir src\vendor\google.golang.org\
mkdir src\vendor\golang.org\x
git clone http://github.com/grpc/grpc-go src/vendor/google.golang.org/grpc
git clone http://github.com/golang/text.git src/vendor/golan.org/x/text
git clone http://github.com/golang/net.git src/vendor/golang.org/x/net
git clone http://github.com/google/go-genproto src/vendor/google.golang.org/genproto
go get -u -v github.com/golang/protobuf/proto
go get -u -v github.com/golang/protobuf/protoc-gen-go
goto end

:end