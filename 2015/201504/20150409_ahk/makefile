Pre=ahk
Bin_Path=bin
Bin=${Bin_Path}/${Pre}${APP_TAG}
ReleasePkg=${Pre}-$(shell git describe master --tags).tar.xz 
#OS=linux
#ARCH=386
#ARCH=amd64

GCFLAGS= "-N -l"

.PHONY: clean pkg src release

all: pkg
	@echo "==== building ===="
	@export GOPATH=`pwd`:${GOPATH}; \
		CGO_ENABLED=1 GOOS=${OS} GOARCH=${ARCH} go build -o ${Bin} -gcflags ${GCFLAGS} src/main.go 
	@echo "==== finish ===="
	@echo "output as " ${Bin}

pkg:
	@echo "==== env prepare ===="
	@export GOPATH=`pwd`
	@go get github.com/go-sql-driver/mysql

clean:
	@echo "==== clean ===="
	@go clean
	rm -rf ${Bin_Path}/${Pre}_*
	@echo 

src:
	git archive master --format=zip > ${Pre}-`git describe master --tags`.zip

release: all
	cd ${Bin_Path};ln -sf ${Pre}_${APP_TAG} ${Pre}
	cd ${Bin_Path};tar -Jcvf ../${ReleasePkg} ${Pre}_${APP_TAG} config.json 

