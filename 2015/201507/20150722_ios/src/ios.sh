#!/bin/sh
SDK=iphoneos
SDK_PATH=`xcrun --sdk $SDK --show-sdk-path`
export IPHONEOS_DEPLOYMENT_TARGET=5.1
CLANG=`xcrun --sdk $SDK --find clang`

if [ "$GOARCH" == "arm" ]; then
	CLANGARCH="armv7"
elif [ "$GOARCH" == "arm64" ]; then
	CLANGARCH="arm64"
else
	echo "unknown GOARCH=$GOARCH" >&2
	exit 1
fi

exec $CLANG -arch $CLANGARCH -isysroot $SDK_PATH -mios-version-min=6.0 "$@"
