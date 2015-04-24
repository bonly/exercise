#ifndef __POME_H__
#define __POME_H__

#include <pomelo.h>
extern void init();

typedef struct sct{
  int def;
}Tsct;

typedef struct Tmy{ //Tmy
  int abc;
  pc_client_t *client;
  Tsct kstc;  
} Tmy;

#endif
