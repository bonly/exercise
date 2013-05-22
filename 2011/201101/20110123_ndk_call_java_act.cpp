/**
Android : How to call a java activity using an intent from a purely native NDK application?
Many android programmers think that it is impossible to call a java activity using intents from a purely native NDK application (native activity as the MAIN launcher activity).

It is NOT true.

But it is not a straight forward though and is little tricky.
People think that it is impossible because they think that context will not be there for pure native NDK applications and hence we cannot launch any java activity.
Following is the code snippet which can be used to get the context in a purely native NDK application code.

JNIEnv *env;
nativestate->activity->vm->AttachCurrentThread(&env, NULL);
jclass activityClass = env->FindClass("android/app/NativeActivity");
jmethodID getClassLoader = env->GetMethodID(activityClass,"getClassLoader", "()Ljava/lang/ClassLoader;");
jobject cls = env->CallObjectMethod(nativestate->activity->clazz, getClassLoader);
jclass classLoader = env->FindClass("java/lang/ClassLoader");
jmethodID findClass = env->GetMethodID(classLoader, "loadClass", "(Ljava/lang/String;)Ljava/lang/Class;");
jmethodID contextMethod = env->GetMethodID(activityClass, "getApplicationContext", "()Landroid/content/Context;");
jobject contextObj = env->CallObjectMethod(nativestate->activity->clazz, contextMethod);

We can use this contextObj value, which is actually the context for launching any java activity using an Intent. I have tested this and it works perfectly fine.

http://vkswtips.blogspot.com/2012/01/android-how-to-call-java-activity-using.html

http://stackoverflow.com/questions/9990830/sending-intents-from-an-android-ndk-application
ending intents from an Android-NDK application
There is a "am" command that you can run that will send Intents to Activities or Services.

const char *cmd="am startservice -a %s --ei ars_flag 2 --ei invitationType %d --ei mode 1 --es ars_gadget_uuid \"%s\" --ei ars_conn_handle %d --es ars_user_uuid \"%s\" --es ars_username \"%s\"";

//sprintf (cmdbuffer,cmd,...)

//system (cmdbuffer)

like this:
 char* args = {"/system/bin/am", "start","-a", "android.intent.action.VIEW", "http://www.google.com/"};
 if(fork() == 0) {
     execvp("/system/bin/am", args);
 }
http://journals.ecs.soton.ac.uk/java/tutorial/native1.1/implementing/index.html
*/
