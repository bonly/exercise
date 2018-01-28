NDK_ROOT=/home/opt/andk

.PHONY: and linux proto cs gsrv get clean

all: proto gsrv cs

clean:
	rm -rf src/cs src/He
	rm -rf gsrv client.exe 

and:
	@echo '构建android版本动态库'
	CC=${NDK_ROOT}/bin/arm-linux-androideabi-gcc \
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm \
	go build -buildmode=c-shared -x -o libtechappen.so


#CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -x -o libabc.so -ldflags="-shared -pthread" dll.go

linux:
	@echo "构建linux版本动态库"
	export GOPATH=`pwd`:${GOPATH}; \
	GOOS=linux \
	CGO_ENABLED=1 \
	go build -v -buildmode=c-shared -x -o libtechappen.so srv.go

proto:
	@echo "构建golang及c#版本协议"
	mkdir -p src/He
	mkdir -p src/cs
	export GOPATH=`pwd`; \
	protoc -I src --go_out=plugins=grpc:src/He \
		src/he.proto
	protoc -I=src \
		--grpc_out=src/cs \
		--csharp_out=src/cs \
		--plugin=protoc-gen-grpc=./Grpc.Tools.1.3.6/tools/linux_x64/grpc_csharp_plugin \
		src/he.proto

cs:
	@echo "构建c#客户端"
	mcs -lib:Grpc.Core.1.3.6/lib/net45/,Google.Protobuf.3.3.0/lib/net45 \
		-r:Grpc.Core.dll,Google.Protobuf.dll \
		-recurse:'*.cs' \
		-out:client.exe
	@#mcs -t:library 创建dll


gsrv:
	@echo "构建golang服务端"
	export GOPATH=`pwd`:${GOPATH}; \
		go build -v -o gsrv gsrv.go

get:
	@echo "设置golang环境"
	#go get -u -v github.com/golang/protobuf/proto
	go get -u -v github.com/golang/protobuf/{proto,protoc-gen-go}

cs_pre:
	@echo "设置c#环境"
	nuget install Grpc.Core
	nuget install Grpc
	nuget install Grpc.Tools
	nuget install Google.Protobuf

run:
	@echo "运行客户端"
	MONO_PATH=Google.Protobuf.3.3.0/lib/net45:\
	Grpc.Core.1.3.6/lib/net45:\
	System.Interactive.Async.3.1.1/lib/net45 \
	mono client.exe
