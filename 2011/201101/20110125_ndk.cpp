//BEGIN_INCLUDE(all)
#include <jni.h>
#include <errno.h>

#include <android/log.h>
#include <android_native_app_glue.h>

#define LOGI(...) ((void)__android_log_print(ANDROID_LOG_INFO, "native-activity", __VA_ARGS__))
#define LOGW(...) ((void)__android_log_print(ANDROID_LOG_WARN, "native-activity", __VA_ARGS__))

/**
 * Shared state for our app.
 */
struct my_struct {
    struct android_app* app;
};

/**
 * Process the next input event.
 */
static int32_t engine_handle_input(struct android_app* app, AInputEvent* event) {
        struct my_struct* str = (struct my_struct*)app->userData;

        if (AInputEvent_getType(event) == AINPUT_EVENT_TYPE_KEY) {
                LOGI("AINPUT_EVENT_TYPE_KEY struct ptr=%p", str);
                return 1;
        }
        return 0;
}

/**
 * Process the next main command.
 */
static void engine_handle_cmd(struct android_app* app, int32_t cmd) {
        struct my_struct* str = (struct my_struct*)app->userData;

        switch (cmd) {
        case APP_CMD_SAVE_STATE:
                // The system has asked us to save our current state.  Do so.
                LOGI("APP_CMD_SAVE_STATE struct ptr=%p", str);
                break;
        case APP_CMD_INIT_WINDOW:
                // The window is being shown, get it ready.
                LOGI("APP_CMD_INIT_WINDOW struct ptr=%p", str);
                break;

        break;
        case APP_CMD_TERM_WINDOW:
                // The window is being hidden or closed, clean it up.
                LOGI("APP_CMD_TERM_WINDOW struct ptr=%p", str);
                break;
        case APP_CMD_GAINED_FOCUS:
        // When our app gains focus, we start monitoring the accelerometer.
                break;
        case APP_CMD_LOST_FOCUS:
        // When our app loses focus, we stop monitoring the accelerometer.
        // This is to avoid consuming battery while not being used.
                break;
        }
}

/**
 * This is the main entry point of a native application that is using
 * android_native_app_glue.  It runs in its own thread, with its own
 * event loop for receiving input events and doing other things.
 */

void android_main(struct android_app* state) {
        struct my_struct str;

        // Make sure glue isn't stripped.
        app_dummy();

        /* Call Java class */   
        JNIEnv* env_custom;

        LOGI("pointer vm=%p env=%p", state->activity->vm, state->activity->env);
        (*state->activity->vm)->AttachCurrentThread(state->activity->vm, &env_custom, NULL);

        jclass activityClass = (*env_custom)->FindClass(env_custom, "android/app/NativeActivity");
        LOGI("activityClass=%p", activityClass);

        jmethodID getClassLoader = (*env_custom)->GetMethodID(env_custom, activityClass, "getClassLoader", "()Ljava/lang/ClassLoader;");

        jobject cls = (*env_custom)->CallObjectMethod(env_custom, state->activity->clazz, getClassLoader);

        jclass classLoader = (*env_custom)->FindClass(env_custom, "java/lang/ClassLoader");
        LOGI("classLoader=%p", classLoader);

        jmethodID findClass = (*env_custom)->GetMethodID(env_custom, classLoader, "loadClass", "(Ljava/lang/String;)Ljava/lang/Class;");        
        LOGI("findClass=%p", findClass);

        jstring strClassName = (*env_custom)->NewStringUTF(env_custom, "com/MyAct/MyActClass");
        jclass classIWant = (jclass)(*env_custom)->CallObjectMethod(env_custom, cls, findClass, strClassName);
        LOGI("classIWnant=%p", classIWant);

        /* constructor */
        jmethodID constructor = (*env_custom)->GetMethodID(env_custom, classIWant, "<init>","()V");
        LOGI("constructor=%p", constructor);
        jobject my_obj = (*env_custom)->CallObjectMethod(env_custom, classIWant, constructor);

        /* method */
        jmethodID slog = (*env_custom)->GetMethodID(env_custom, classIWant, "simpleLog","()V");
        LOGI("slog=%p", slog);
        (*env_custom)->CallStaticObjectMethod(env_custom, my_obj, slog);

        /* static method */
        jmethodID dump = (*env_custom)->GetStaticMethodID(env_custom, classIWant, "dumpCameraInfo","()V");
        LOGI("dump=%p", dump);
        (*env_custom)->CallStaticObjectMethod(env_custom, classIWant, dump);

        jint res = (*state->activity->vm)->DetachCurrentThread(state->activity->vm);

        memset(&str, 0, sizeof(str));
        state->userData = &str;
        state->onAppCmd = engine_handle_cmd;
        state->onInputEvent = engine_handle_input;
        str.app = state;


    while (1) {
        // Read all pending events.
        int ident;
        int events;
        struct android_poll_source* source;

        // If not animating, we will block forever waiting for events.
        // If animating, we loop until all events are read, then continue
        // to draw the next frame of animation.
        while ((ident=ALooper_pollAll(0, NULL, &events, (void**)&source)) >= 0) {

            // Process this event.
            if (source != NULL) {
                source->process(state, source);
            }

            // Check if we are exiting.
            if (state->destroyRequested != 0) {
                return;
            }
        }

    }
}
