OPT=/home/opt
CFLAGS=-D_GLIBCXX_USE_CXX11_ABI=0 -std=c99 \
-fsanitize-coverage=edge -fsanitize=undefined -fno-omit-frame-pointer \
-Wno-unused-command-line-argument -Wvarargs \
-pthread -I ${OPT}/grpc/include
#
LDFLAGS=-D_GLIBCXX_USE_CXX11_ABI=0
#-static-libstdc++ -static-libgcc --enable-static-nss

LIBRARY_PATH=-L ${OPT}/grpc/lib64 -lprotobuf -L${OPT}/grpc/lib  \
-L/home/bonly/opt/grpc.git/build/third_party/boringssl/ssl/ \
-L/home/bonly/opt/grpc.git/build/third_party/boringssl/crypto/ \
-L/home/bonly/opt/grpc.git/build/third_party/cares/ \
-L/home/bonly/opt/grpc.git/build/third_party/boringssl/decrepit/ \
-L /usr/lib/ \

LIBS=-lz -lrt -lm -lpthread -pthread -ldl -lssl -lgrpc++ -lprotobuf-lite \

#-lgrpc++ -lgrpc++_reflection  -lprotobuf \
-lgrpc++_error_details  \
-lgrpc++_cronet  -lgrpc++_unsecure  

#-static -lssl
CC=gcc
LD=ld

TARGET  = he
SOURCES = $(shell echo *.cc)
OBJECTS = $(SOURCES:.cc=.o)
#SOURCES = $(notdir $(FILES))
#OBJECTS = $(patsubst %.cc,%.o,$(SOURCES))
STLIB = #$(shell echo /home/opt/grpc/lib64/*c++*.a) \
$(shell echo /home/opt/grpc/lib/*.a) \
/home/opt/grpc/lib64/libprotobuf.a \
/home/opt/grpc/lib64/libgpr.a \
/home/opt/grpc/lib64/libgrpc.a \
/home/opt/grpc/lib64/libgrpc_cronet.a \



#/home/bonly/opt/grpc.git/build/third_party/boringssl/ssl/libssl.a \
/home/bonly/opt/grpc.git/build/third_party/cares/libcares.a \
/home/bonly/opt/grpc.git/build/third_party/boringssl/crypto/libcrypto.a \
/home/bonly/opt/grpc.git/build/third_party/boringssl/decrepit/libdecrepit.a \

#$(shell echo /home/opt/grpc/lib64/*.a) \
$(shell echo /home/opt/grpc/lib/*.a) \
/home/bonly/opt/grpc.git/build/third_party/cares/libcares.a \



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
