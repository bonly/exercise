all:
	export GOPATH=`pwd`:${GOPATH}; \
	GOOS=linux \
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	go build -x -v app/main.go
