all:
	ANDROID_HOME=${OPT}/adt/sdk \
	GOPATH=${PWD}:${GOPATH} \
	gomobile bind -v -target=android \
	-o mypkg.aar mypkg


## build standalone gcc
#cd ~/android-ndk && $NDK/build/tools/make-standalone-toolchain.sh \
    --arch=arm --platform=android-21 --install-dir=/tmp/my-android-toolchain
# build go for arm
# CC_FOR_TARGET=$ANDK_CC GOOS=android GOARCH=arm GOARM=7 ./make.bash
# build arm program
# GOARCH=arm go build -v -o update_hosts host_and.go
