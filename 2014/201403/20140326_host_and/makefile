all:
	GOARCH=arm \
	go build -v \
	-o update_hosts 20140326_host_and.go


## build standalone gcc
#cd ~/android-ndk && $NDK/build/tools/make-standalone-toolchain.sh \
--arch=arm --platform=android-21 --install-dir=/tmp/my-android-toolchain
# build go for arm
# CC_FOR_TARGET=$ANDK_CC GOOS=android GOARCH=arm GOARM=7 ./make.bash
#CC_FOR_TARGET=/home/opt/andk/bin/arm-linux-androideabi-gcc GOOS=android GOARCH=arm GOARM=7 CGO_ENABLED=1 ./make.bash
#CC_FOR_TARGET=/home/opt/andk/bin/arm-linux-androideabi-gcc GOOS=android GOARM=7 


#	go build -v -p=8 -pkgdir=${GOMOBILE}/pkg_android_arm -tags="" -ldflags="-extldflags=-pie -pthread" \
#	GOROOT=/home/opt/go.git \
	GOMOBILE="/home/opt/go.my/pkg/gomobile" \
	CC_FOR_TARGET=/home/opt/andk/bin/arm-linux-androideabi-gcc \
	CC=/home/bonly/opt/andk/bin/arm-linux-androideabi-gcc \
	CXX=/home/bonly/opt/andk/bin/arm-linux-androideabi-g++ \
	GOARM=7 GOOS=linux GOARCH=arm \
	CGO_ENABLED=1 GOARM=7 \
