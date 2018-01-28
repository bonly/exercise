all:
	$(eval GOPATH=$(shell echo $${PWD%src*}):$(GOPATH)) 
	@#GOPATH=$(shell echo $${PWD%src*}):$(GOPATH) \ 
	gopherjs build -v cli.go


clean:
	rm -rf ../pkg cli.js*
