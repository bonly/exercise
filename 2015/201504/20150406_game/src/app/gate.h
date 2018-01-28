#ifndef __GATE_H__
#define __GATE_H__


typedef void (*Callback)(unsigned int sn, char *buffer);
extern Callback fn;
void bridge_callback(unsigned int sn, char *buffer);

#endif