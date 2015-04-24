package pomelo

/*
#cgo CFLAGS: -I/home/bonly/libpomelo/include -I/home/bonly/libpomelo -I/home/bonly/libpomelo/deps/uv/include -I/home/bonly/libpomelo/deps/jansson/src
#cgo linux CFLAGS: -DLINUX=1
#cgo LDFLAGS: -L/home/bonly/libpomelo -lpomelo -L/home/bonly/libpomelo/deps/jansson/src/.libs/ -ljansson -L/home/bonly/libpomelo/deps/uv -luv
#include <pomelo.h>

pc_client_t *client=0;

#include <stdio.h>

void Pa(){
  printf("%s\n", "hi world");
  client = pc_client_new();
  struct sockaddr_in address;

  memset(&address, 0, sizeof(struct sockaddr_in));
  address.sin_family = AF_INET;

  address.sin_port = htons(4020);
  address.sin_addr.s_addr = inet_addr("192.168.1.156");
  
  if(pc_client_connect(client, &address)) {
    printf("fail to connect server.\n");
    pc_client_destroy(client);
    return 1;
  }
  
  printf("%s\n", "new client");
}
*/
import "C"

import "fmt"

func Main(){
  fmt.Println("begin");
  
  C.Pa();
  
  fmt.Println("end");
}
