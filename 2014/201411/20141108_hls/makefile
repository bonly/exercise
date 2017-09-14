-include makefile.init

VerFlag='-X main.Version=${APP_VERSION} -X main.Code=${APP_CODE} -X main.App_time=${APP_TIME} ' #-s
Pre=hls
Bin_Path=bin
Bin=${Bin_Path}/${Pre}_${APP_TAG}
ReleasePkg=${Pre}-$(shell git describe master --tags).tar.xz 
OS=linux
#ARCH=386
#OS=windows
#ARCH=amd64

GCFLAGS= "-N -l"

.PHONY: clean pkg src release

all: pkg
	@echo "==== building ===="
	@#export GOPATH=`pwd`:${GOPATH}; go build -x -v -o ${Bin} -ldflags ${VerFlag} -gcflags ${GCFLAGS} src/main.go 
	@export GOPATH=`pwd`:${GOPATH}; \
		CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -o ${Bin} -ldflags ${VerFlag} -gcflags ${GCFLAGS} src/main.go 
	@echo "==== finish ===="
	@echo "output as " ${Bin}
	@cd ${Bin_Path};ln -fs ${Pre}_${APP_TAG} ${Pre}

pkg:
	@export GOPATH=`pwd`
	@#@go get github.com/Sirupsen/logrus

clean:
	@go clean
	rm -rf ${Bin_Path}/${Pre}_*

src:
	git archive master --format=zip > ${Pre}-`git describe master --tags`.zip

release: all
	cd ${Bin_Path};ln -sf ${Pre}_${APP_TAG} ${Pre}
	cd ${Bin_Path};tar -Jcvf ../${ReleasePkg} ${Pre}_${APP_TAG} config.json ${Pre}

