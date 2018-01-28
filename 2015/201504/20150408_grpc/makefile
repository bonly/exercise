OPT=/home/opt
ROOT=${OPT}/grpc.git

CFLAGS=-D_GLIBCXX_USE_CXX11_ABI=0 -std=c99 \
-fsanitize-coverage=edge -fsanitize=undefined -fno-omit-frame-pointer \
-Wno-unused-command-line-argument -Wvarargs \
-pthread -I ${ROOT}/include

LDFLAGS=-D_GLIBCXX_USE_CXX11_ABI=0

LIBRARY_PATH=-L ${ROOT}/lib64 -lprotobuf -L${ROOT}/lib  \
-L{ROOT}/build/third_party/boringssl/ssl/ \
-L/home/bonly/opt/grpc.git/build/third_party/boringssl/crypto/ \
-L/home/bonly/opt/grpc.git/build/third_party/cares/ \
-L/home/bonly/opt/grpc.git/build/third_party/boringssl/decrepit/ \
-L /usr/lib/ \

LIBS=-lz -lrt -lm -lpthread -pthread -ldl -lssl -lgrpc++ -lprotobuf-lite \

CC=gcc
LD=ld

TARGET  = he

STLIB = 

SOURCES = $(shell echo *.c)
#SOURCES = $(notdir $(FILES))

OBJECTS = $(SOURCES:.c=.o)
#OBJECTS = $(patsubst %.cc,%.o,$(SOURCES))


all: clean $(TARGET)

$(TARGET): $(OBJECTS) $(STLIB)
	$(CC) -v $(LDFLAGS) $(LIBRARY_PATH) $(LIBS)  $^ -o $@

.cc.o:
	$(CC) $(CFLAGS) -c $< -o $@

.PHONY: proto clean test

proto:
	protoc -I . --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` helloworld.proto
	protoc -I . --cpp_out=. helloworld.proto

clean:
	rm -rf *.o he

test:
	echo $(STLIB)
