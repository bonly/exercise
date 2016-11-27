#include "ofMain.h"
#include "ofApp.h"
#include "libgf.h"
#include <dlfcn.h>
#include <cstdio>
//========================================================================
int main( ){
	void *handle = dlopen("libgf", RTLD_LAZY); //打开
    if (!handle) {
        printf("dlopen failed\n");
        return 1;
    }

    void (*some_func)();
	some_func = (void (*)()) dlsym(handle,"Gf");
	if(some_func==NULL) {
        printf("dlsym failed\n");
        return 1;
	}
	some_func();
	ofSetupOpenGL(1024,768, OF_WINDOW);			// <-------- setup the GL context

	// this kicks off the running of my app
	// can be OF_WINDOW or OF_FULLSCREEN
	// pass in width and height too:
	ofRunApp( new ofApp());

    dlclose(handle); //关闭
}
