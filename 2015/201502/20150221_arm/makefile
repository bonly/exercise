NDK_ROOT=/home/opt/andk

all:
	CC=${NDK_ROOT}/bin/arm-linux-androideabi-gcc \
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm \
	go build -x -o hello hello.go


#CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -x -o libabc.so -ldflags="-shared -pthread" dll.go

