@echo on

@echo %1
@if {%1} == {} goto All
@if {%1} == {env} goto Env
@if {%1} == {pkg} goto Pkg
@if {%1} == {cs} goto Cs
@if {%1} == {gsrv} goto Gsrv
@goto end

:Env
@setx GOPATH %CD%;%GOPATH% /m
@echo GOPATH=%GOPATH%
goto end

:All
@set GOPATH=%CD%;%GOPATH%
@rem gcc -c -o src/app/gate.o src/app/gate.c
@rem @ar rcs src/app/libgate.a src/app/gate.o
@rem @set LIBRARY_PATH=%CD%/src/app
@rem @set CGO_LDFLAGS=-L %CD%\src\app -lgate
@rem @echo %CGO_LDFLAGS%
@set CGO_ENABLED=1
go build -x -v -buildmode=c-archive -o client.a src/app/client.go
@gcc -m64 -shared -o libtechappen.so cli.def client.a -static -lstdc++ -lwinmm -lntdll -lWs2_32
@goto end

:Pkg
@mkdir src\vendor\google.golang.org\
@mkdir src\vendor\golang.org\x
@git clone http://github.com/grpc/grpc-go src/vendor/google.golang.org/grpc
@git clone http://github.com/golang/text.git src/vendor/golang.org/x/text
@git clone http://github.com/golang/net.git src/vendor/golang.org/x/net
@git clone http://github.com/google/go-genproto src/vendor/google.golang.org/genproto
@go get -u -v github.com/golang/protobuf/proto
@go get -u -v github.com/golang/protobuf/protoc-gen-go
@goto end

:Cs
rem http://download.microsoft.com/download/9/3/F/93FCF1E7-E6A4-478B-96E7-D4B285925B00/vc_redist.x64.exe
rem https://download.mono-project.com/archive/5.0.1/windows-installer/mono-5.0.1.1-x64-0.msi
@mcs -out:client.exe Main.cs
@goto end

:Gsrv
@go build -v -o gsrv src/app/gsrv.go

:SrvReq
@go get github.com/go-sql-driver/mysql
@go get github.com/jmoiron/sqlx

:end
@echo 完成
