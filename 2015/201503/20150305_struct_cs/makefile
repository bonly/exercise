NDK_ROOT=/home/opt/andk

and:
	CC=${NDK_ROOT}/bin/arm-linux-androideabi-gcc \
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm \
	go build -buildmode=c-shared -x -o libtechappen.so


#CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -x -o libabc.so -ldflags="-shared -pthread" dll.go

linux:
	GOOS=linux \
	CGO_ENABLED=1 \
	go build -buildmode=c-shared -x -o libtechappen.so
