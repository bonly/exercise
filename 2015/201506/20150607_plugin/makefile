all: plugin
	go build -o main main.go

plugin:
	go build -o myplugin.so -buildmode=plugin myplugin.go
