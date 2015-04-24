#include "20110615_pome.h"
#include "_cgo_export.h"
#include <pomelo.h>
#include <stdio.h>

void Pa(){
  printf("%s\n", "hi world");
  Gfun4C();
  
  pc_client_t *client = pc_client_new();
  struct sockaddr_in address;

  memset(&address, 0, sizeof(struct sockaddr_in));
  address.sin_family = AF_INET;

  address.sin_port = htons(4010);
  address.sin_addr.s_addr = inet_addr("192.168.1.111");
  
  if(pc_client_connect(client, &address)) {
    printf("fail to connect server.\n");
    pc_client_destroy(client);
    return;
  }
  
  printf("%s\n", "new client");
}
