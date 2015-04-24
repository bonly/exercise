#ifndef __LIBPOMELO_H__
#define __LIBPOMELO_H__
#include <pomelo.h>

extern json_t* Add(const char* key, const char* value);

extern pc_client_t* Connect(const char* ip, const int port);
extern void Notify(pc_client_t* cli, const char *route, json_t *msg);

extern void on_notifyed(pc_notify_t *req, int status);
extern void wait_join(pc_client_t *cli);
#endif
