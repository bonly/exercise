all:
	$(eval GOPATH=$(shell echo $${PWD%src*}):$(GOPATH))
	CGO_ENABLED=1 \
	go build -buildmode=c-shared -o libtechappen.so main.go 
	@#main_c.go
	@#GOARCH=386  \


android:
	#$(eval NDK_ROOT=$(HOME)/opt/andk)
	$(eval NDK_ROOT=$(HOME)/opt/andk13)
	$(eval GOPATH=$(PWD):$(GOPATH))
	CC=$(NDK_ROOT)/bin/arm-linux-androideabi-gcc \
	CXX=$(NDK_ROOT)/bin/arm-linux-androideabi-g++ \
	AR=$(NDK_ROOT)/bin/arm-linux-androideabi-ar \
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm \
	GOARM=7 \
	go build -gcflags="-shared" -buildmode=c-shared -o libtechappen.so main.go

clean:
	rm -rf libtechappen.so libtechappen.h


exe:
	#$(eval NDK_ROOT=$(HOME)/opt/andk)
	$(eval NDK_ROOT=$(HOME)/opt/andk13)
	$(eval GOPATH=$(PWD):$(GOPATH))
	CC=$(NDK_ROOT)/bin/arm-linux-androideabi-gcc \
	CXX=$(NDK_ROOT)/bin/arm-linux-androideabi-g++ \
	AR=$(NDK_ROOT)/bin/arm-linux-androideabi-ar \
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm \
	GOARM=7 \
	go build -ldflags="" -buildmode=pie -o libtec.so exe.go

aso:
	$(eval NDK_ROOT=$(HOME)/opt/andk13)
	$(eval GOPATH=$(PWD):$(GOPATH))
	CC=$(NDK_ROOT)/bin/arm-linux-androideabi-gcc \
	CXX=$(NDK_ROOT)/bin/arm-linux-androideabi-g++ \
	AR=$(NDK_ROOT)/bin/arm-linux-androideabi-ar \
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm \
	GOARM=7 \
	go build -ldflags="" -buildmode=pie -o libtec.so udp.go
