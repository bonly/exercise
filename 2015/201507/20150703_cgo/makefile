NDK_ROOT=/home/opt/andk
CC=${NDK_ROOT}/bin/arm-linux-androideabi-gcc
AR=${NDK_ROOT}/bin/arm-linux-androideabi-ar 

all:
	CC=${NDK_ROOT}/bin/arm-linux-androideabi-gcc \
	CGO_ENABLED=1 GOARCH=arm GOOS=darwin go build -buildmode=c-archive cs.go