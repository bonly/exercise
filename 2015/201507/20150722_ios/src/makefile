all: ios

clean:
	rm -rf lib* ios *.bundle


pc:
	$(eval GOPATH=$(shell echo $${PWD%src*}):${GOPATH}) 
	go build -o libtechappen.so cli_pc.go

ios: tarm64 tarm7 tx86_64 ti386
	xcrun lipo -arch arm64 ios/arm64/libcs.a \
	-arch armv7 ios/armv7/libcs.a \
	-arch x86_64 ios/x86_64/libcs.a \
	-arch i386 ios/i386/libcs.a \
	-create -output techappen.bundle
	file techappen.bundle

tarm64:
	mkdir -p ios/arm64
	$(eval GOPATH=$(shell echo $${PWD%src*}):${GOPATH}) 
	CGO_ENABLED=1 \
	GOARCH=arm64 \
	CC=`pwd`/ios.sh \
	CXX=`pwd`/ios.sh \
	go build -buildmode=c-archive -o ios/arm64/libcs.a cli_pc.go

tarm7:
	mkdir -p ios/armv7
	$(eval GOPATH=$(shell echo $${PWD%src*}):${GOPATH}) 
	CGO_ENABLED=1 \
	GOARCH=arm \
	GOARM=7 \
	CC=`pwd`/ios.sh  \
	CXX=`pwd`/ios.sh \
	go build -buildmode=c-archive -o ios/armv7/libcs.a cli_pc.go

tx86_64:
	mkdir -p ios/x86_64
	$(eval GOPATH=$(shell echo $${PWD%src*}):${GOPATH}) 
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	GOOS=darwin \
	go build -buildmode=c-archive -o ios/x86_64/libcs.a cli_pc.go

ti386:
	mkdir -p ios/i386
	$(eval GOPATH=$(shell echo $${PWD%src*}):${GOPATH}) 
	CGO_ENABLED=1 \
	GOARCH=386 \
	GOOS=darwin \
	go build -buildmode=c-archive -o ios/i386/libcs.a cli_pc.go
