static void engine_save_buffer(struct engine *engine)
{
        JavaVM *vm;
        JNIEnv *env;
        jclass clazz;
        jmethodID method;
        jobject obj;
        jstring str;
        char fname[256];

        // first attach our thread to the VM so we can get a JNI environment pointer
        vm = engine->app->activity->vm;
        (*vm)->AttachCurrentThread(vm, &env, NULL);

        // lookup the external storage directory
        clazz = (*env)->FindClass(env, "android/os/Environment");
        method = (*env)->GetStaticMethodID(env, clazz, "getExternalStorageDirectory", "()Ljava/io/File;");
        obj = (*env)->CallStaticObjectMethod(env, clazz, method);

        // convert it to a string
        clazz = (*env)->GetObjectClass(env, obj);
        method = (*env)->GetMethodID(env, clazz, "getAbsolutePath", "()Ljava/lang/String;");
        str = (jstring)(*env)->CallObjectMethod(env, obj, method);

        // get a pointer to the string and create the output .ppm filename
        const char *path = (*env)->GetStringUTFChars(env, str, NULL);
        snprintf(fname, sizeof(fname), "%s/aobench-android.ppm", path);
        (*env)->ReleaseStringUTFChars(env, str, path);

        // save the .ppm
        aobench_saveppm(fname, engine->width, engine->height, engine->img);
}
