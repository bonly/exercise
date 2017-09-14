/*
 * auth: bonly
 * create: 2016.5.27
 * http://thiemonge.org/getting-started-with-uinput
 * http://www.einfochips.com/download/dash_jan_tip.pdf
 *
 */

#include "uinputwrapper.h"

//初始化设备
int initVKeyboardDevice(char* uinputPath) {
    int     i;
    int     deviceHandle = -1;
    struct  uinput_user_dev uidev; //设备数据结构

    deviceHandle = open(uinputPath, O_WRONLY | O_NONBLOCK | O_NDELAY);

    // 设置可处理的标识位
    if(deviceHandle > 0) {
        if(ioctl(deviceHandle, UI_SET_EVBIT, EV_KEY) < 0 &&
           ioctl(deviceHandle, UI_SET_EVBIT, EV_REL) < 0 &&
           ioctl(deviceHandle, UI_SET_RELBIT, REL_X) < 0 &&
           ioctl(deviceHandle, UI_SET_RELBIT, REL_Y) < 0 ){
            if(releaseDevice(deviceHandle) < 0) {
                exit(EXIT_FAILURE);
            } else {
                deviceHandle = -1;
            }
        } else {
			// 注册关注的事件
            for(i=1; i<256; i++) { //ASIC码
                ioctl(deviceHandle, UI_SET_KEYBIT, i);
            }

            //鼠标相关事件
            ioctl(deviceHandle, UI_SET_KEYBIT, BTN_MOUSE); 
            ioctl(deviceHandle, UI_SET_KEYBIT, BTN_TOUCH); 
            ioctl(deviceHandle, UI_SET_KEYBIT, BTN_LEFT); 
            ioctl(deviceHandle, UI_SET_KEYBIT, BTN_MIDDLE); 
            ioctl(deviceHandle, UI_SET_KEYBIT, BTN_RIGHT); 
            ioctl(deviceHandle, UI_SET_KEYBIT, BTN_FORWARD); 
            ioctl(deviceHandle, UI_SET_KEYBIT, BTN_BACK); 
            
            memset(&uidev, 0, sizeof (uidev));
            snprintf(uidev.name, UINPUT_MAX_NAME_SIZE, "uinput_vkeyboard");
            uidev.id.bustype = BUS_USB;
            uidev.id.vendor  = 0x4711;
            uidev.id.product = 0x0815;
            uidev.id.version = 1;

            //创建设备到子系统中
            if (write(deviceHandle, &uidev, sizeof (uidev)) < 0) {
                exit(EXIT_FAILURE);
            }

            //重新请求获取此设备，以便验证已创建成功
            if (ioctl(deviceHandle, UI_DEV_CREATE) < 0) {
                exit(EXIT_FAILURE);
            }

            // sleep(2);

        }
    }

    return deviceHandle;
}

int sendBtnEvent(int deviceHandle, int key, int btnState) {
	// check whether the keycode is valid and return -1 on error
	if (key < 1 || key > 255) {
		return -1;
	}

	// btnState should only be either 0 or 1
	if (btnState < 0 || btnState > 1) {
		return -1;
	}

    struct input_event ev;
    memset(&ev, 0, sizeof(ev));
    
    ev.type = EV_KEY;
    ev.code = key;  
    ev.value= btnState;
    int ret = write(deviceHandle, &ev, sizeof(ev));

	// in case of any error return the error and skip syncing events
	if (ret < 0) {
		return ret;
	}

	ret = syncEvents(deviceHandle, &ev);

	return ret;
}

int syncEvents(int deviceHandle, struct input_event *ev) {
	memset(ev, 0, sizeof(*ev));

	ev->type  = EV_SYN;
	ev->code  = 0;
	ev->value = SYN_REPORT;

	return write(deviceHandle, ev, sizeof(*ev));
}

int releaseDevice(int deviceHandle) {
	return ioctl(deviceHandle, UI_DEV_DESTROY);
}

void send_click_events(int uinp_fd) { 
    struct input_event event;
    // Move pointer to (0,0) location 
    memset(&event, 0, sizeof(event)); 
    gettimeofday(&event.time, NULL);         
    
    event.type = EV_REL; 
    event.code = REL_X; 
    event.value = 100;         
    write(uinp_fd, &event, sizeof(event));         
    
    event.type = EV_REL;         
    event.code = REL_Y;         
    event.value = 100;         
    write(uinp_fd, &event, sizeof(event)); 
    
    event.type = EV_SYN;         
    event.code = SYN_REPORT;         
    event.value = 0;         
    write(uinp_fd, &event, sizeof(event)); 

    // Report BUTTON CLICK - PRESS event  
    memset(&event, 0, sizeof(event)); 
    gettimeofday(&event.time, NULL); 
    
    event.type = EV_KEY; 
    event.code = BTN_LEFT; 
    event.value = 1; 
    write(uinp_fd, &event, sizeof(event)); 
    
    event.type = EV_SYN; 
    event.code = SYN_REPORT; 
    event.value = 0; 
    write(uinp_fd, &event, sizeof(event)); 

    // Report BUTTON CLICK - RELEASE event  
    memset(&event, 0, sizeof(event)); 
    gettimeofday(&event.time, NULL); 
    
    event.type = EV_KEY; 
    event.code = BTN_LEFT; 
    event.value = 0; 
    write(uinp_fd, &event, sizeof(event)); 
    
    event.type = EV_SYN; 
    event.code = SYN_REPORT; 
    event.value = 0; 
    write(uinp_fd, &event, sizeof(event)); 
}