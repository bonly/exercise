#include "libpomelo.h"
#include <string.h>



pc_client_t* Connect(const char* ip, const int port){
	pc_client_t* ret = pc_client_new();
	struct sockaddr_in address;

	memset(&address, 0, sizeof(struct sockaddr_in));
	address.sin_family = AF_INET;

	address.sin_port = htons(port);
	address.sin_addr.s_addr = inet_addr(ip);

	if(pc_client_connect(ret, &address)) {
	    printf("fail to connect server.\n");
	    pc_client_destroy(ret);
	    return 0;
	}
	return ret;
}

void Notify(pc_client_t* cli, const char *route, json_t *msg){
	pc_notify_t *notify = pc_notify_new();
	pc_notify(cli, notify, route, msg, on_notifyed);
}

json_t* Add(const char* key, const char* value){
     json_t* ret = json_object();
     json_t *json_value = json_string(value);
     json_object_set(ret, key, json_value);
     json_decref(json_value);
     return ret;
}

void on_notifyed(pc_notify_t *req, int status){
	if(status == -1){
		printf("Fail to send notify to server.\n");
	}else{
		printf("Notify finished.\n");
	}

	json_t *msg = req->msg;
	json_decref(msg);
	pc_notify_destroy(req);
}

void wait_join(pc_client_t *cli){
	pc_client_join(cli);
	pc_client_destroy(cli);
}
