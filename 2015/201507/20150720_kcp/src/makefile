all:
	$(eval GOPATH=$(shell echo $${PWD%src*}:$(GOPATH)))
	go build -race app.go
