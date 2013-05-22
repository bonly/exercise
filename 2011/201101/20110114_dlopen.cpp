#include <android/log.h>
#include <dlfcn.h>
#include <stdio.h>
#include <string.h>

typedef int (*main_t)(int argc, char** argv);

static int help(const char* argv0)
{
    printf("%s: simple wrapper to work around LD_LIBRARY_PATH\n\n", argv0);
    printf("Args: executable, list all the libraries you need to load in dependency order, executable again, optional parameters\n");
    printf("example: %s /data/local/ttte /data/data/app/com.testwrapper/lib/ttt.so /data/local/ttte 12345\n", argv0);
    printf("Note: the executable should be built with CFLAGS=\"-fPIC -pie\", LDFLAGS=\"-rdynamic\"\n");

    return -1;
}

int main(int argc, char** argv)
{
    int rc, nlibs;
    void *dl_handle;

    if (argc < 2)
    {
        return help(argv[0]);
    }

    __android_log_print(ANDROID_LOG_DEBUG, "wrapper", "running '%s'", argv[1]);

    for (nlibs = 2; ; nlibs++)
    {
        if (nlibs >= argc)
        {
            return help(argv[0]);
        }

        __android_log_print(ANDROID_LOG_DEBUG, "wrapper", "loading '%s'", argv[nlibs]);
        dl_handle = dlopen(argv[nlibs], 0); // do not keep the handle, except for the last
        __android_log_print(ANDROID_LOG_DEBUG, "wrapper", "loaded '%s' -> %p", argv[nlibs], dl_handle);
        if (strcmp(argv[1], argv[nlibs]) == 0)
        {
            break;
        }
    }

    main_t pmain = (main_t)dlsym(dl_handle, "main");
    __android_log_print(ANDROID_LOG_DEBUG, "wrapper", "found '%s' -> %p", "main", pmain);
    rc = pmain(argc - nlibs, argv + nlibs);

//   we are exiting the process anyway, don't need to clean the handles actually

//   __android_log_print(3, "wrapper", "closing '%s'", argv[1]);
//   dlclose(dl_handle);

    return 0;
}

/*Android.mk
LOCAL_PATH := $(call my-dir)

include $(CLEAR_VARS)

LOCAL_MODULE    := wrapper
LOCAL_SRC_FILES := wrapper/main.c
LOCAL_LDLIBS    := -llog
LOCAL_CFLAGS    += -fPIC -fPIE
LOCAL_LDFLAGS   += -rdynamic 
include $(BUILD_EXECUTABLE)
*/

