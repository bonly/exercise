NDK_ROOT=/home/opt/andk

.PHONY: and linux proto cs gsrv get

and:
	CC=${NDK_ROOT}/bin/arm-linux-androideabi-gcc \
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm \
	go build -buildmode=c-shared -x -o libtechappen.so


#CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -x -o libabc.so -ldflags="-shared -pthread" dll.go

linux:
	export GOPATH=`pwd`:${GOPATH}; \
	GOOS=linux \
	CGO_ENABLED=1 \
	go build -v -buildmode=c-shared -x -o libtechappen.so srv.go

proto:
	export GOPATH=`pwd`; \
		protoc -I src/he src/he/he.proto --go_out=plugins=grpc:src/he

cs:
	mcs Main.cs


gsrv:
	export GOPATH=`pwd`:${GOPATH}; \
		go build -v -o gsrv gsrv.go

get:
	#go get -u -v github.com/golang/protobuf/proto
	go get -u -v github.com/golang/protobuf/{proto,protoc-gen-go}
