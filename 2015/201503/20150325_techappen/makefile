NDK_ROOT=/home/opt/andk

.PHONY: and linux proto cs gsrv get clean

all: proto gsrv cs

clean:
	rm -rf src/cs src/He
	rm -rf gsrv client.exe 
	rm -rf libtechappen.so libtechappen.h

soand:
	@echo '构建android版本动态库'
	CC=${NDK_ROOT}/bin/arm-linux-androideabi-gcc \
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm \
	go build -buildmode=c-shared -x -o libtechappen.so src/app/client.go

#CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -x -o libabc.so -ldflags="-shared -pthread" dll.go

solnx:
	@echo "构建linux版本动态库"
	export GOPATH=`pwd`:${GOPATH}; \
	GOOS=linux \
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	go build -x -buildmode=c-shared -o libtechappen.so src/app/client.go

dll:
	@echo "构建windows版本动态库"
	export GOPATH=`pwd`:${GOPATH}; \
	GOOS=windows \
	CGO_ENABLED=1 \
	go build -buildmode=c-archive -o client.a src/app/client.go
	gcc -m64 -shared -o client.dll cli.def client.a -static -lstdc++ -lwinmm -lntdll -lWs2_32
#gcc -shared -pthread -o Hello2.dll Hello2.c ExportHello.a -lWinMM -lntdll -lWS2_32

proto:
	@echo "构建golang及c#版本协议"
	mkdir -p src/proto
	mkdir -p src/cs
	export GOPATH=`pwd`; \
	protoc -I src --go_out=plugins=grpc:src/proto \
		src/Proto.proto
#	protoc -I=src \
#		--grpc_out=src/cs \
#		--csharp_out=src/cs \
#		--plugin=protoc-gen-grpc=./Grpc.Tools.1.3.6/tools/linux_x64/grpc_csharp_plugin \
#		src/Proto.proto

cs:
	@echo "构建c#客户端"
	mcs -lib:. \
		-recurse:'*.cs' \
		-out:client.exe
#mcs -t:library 创建dll
#		-r:libtechappen.so \


gsrv:
	@echo "构建golang服务端"
	export GOPATH=`pwd`:${GOPATH}; \
		go build -v -o gsrv src/app/gsrv.go

dll_pre:
	@echo "设置golang在windows下的环境"
	export GOPATH=`pwd`:${GOPATH}
	mkdir -p src/vendor/github.com/golang/
	mkdir -p src/vendor/google.golang.org/
	git clone http://github.com/grpc/grpc-go src/vendor/google.golang.org/grpc
	git clone http://github.com/golang/text.git src/vendor/github.com/golang/text
	git clone http://github.com/golang/net.git src/vendor/github.com/golang/net
	git clone http://github.com/google/go-genproto src/vendor/google.golang.org/genproto
	#go get -u -v github.com/golang/protobuf/proto
	go get -u -v github.com/golang/protobuf/{proto,protoc-gen-go}



cs_pre:
	@echo "设置c#环境"
	nuget install Grpc.Core
	nuget install Grpc
	nuget install Grpc.Tools
	nuget install Google.Protobuf
#	nuget install System.Interactive.Async
#nuget install System.Runtime.Loader
#nuget install System.Runtime

run:
	@echo "运行客户端"
	MONO_PATH=Google.Protobuf.3.3.0/lib/net45:\
	Grpc.Core.1.3.6/lib/net45 \
	mono client.exe

#System.Runtime.4.3.0 \
#System.Runtime.Loader.4.3.0/lib/netstandard1.5:\
#System.Interactive.Async.3.1.1/lib/net45:\
