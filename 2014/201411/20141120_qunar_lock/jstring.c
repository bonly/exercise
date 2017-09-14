#include <string.h>
#include <jni.h>
#include <pthread.h>
 
// //初始JNI虚拟机环境和线程
JavaVM*   gJavaVM;
JNIEnv*   gJniEnv;
pthread_t gJvmThread;
 
// //Java字符串的类和获取字节的方法ID
// jclass    gStringClass;
// jmethodID gmidStringInit;
// jmethodID gmidStringGetBytes;
 
// //初始化JNI环境，此函数由JNI调用
// jstring Java_com_huz_test_CharsetTest_InitJNIEnv(JNIEnv* env, jobject obj){
//     (*env)->GetJavaVM(env, &gJavaVM);
//     gJniEnv=env;
//     gJvmThread=pthread_self();//记住当前JNI环境的线程
 
//     //获取Java String类和回调方法ID信息，由于每次转换都需要，因此用全局变量记下来，免得浪费时间重复执行
//     gStringClass= (*env)->FindClass(env,"java/lang/String");
//     gmidStringGetBytes= (*env)->GetMethodID(env,gStringClass, "getBytes", "(Ljava/lang/String;)[B");
//     gmidStringInit= (*env)->GetMethodID(env,gStringClass, "<init>", "([BLjava/lang/String;)V");

//     return (*env)->NewStringUTF(env, "OK");
// }
 
// //由Java String转为指定编码的char
// int  jstringToPchar(JNIEnv* env, jstring jstr, const char * encoding, char* outbuf, int outlen){
//     char* rtn = NULL;
//     jstring jencoding;
//     if (encoding==NULL)
//         jencoding= (*env)->NewStringUTF(env,"utf-8");
//     else
//         jencoding=(*env)->NewStringUTF(env,encoding);
//     jbyteArray barr= (jbyteArray)(*env)->CallObjectMethod(env,jstr, gmidStringGetBytes, jencoding);
//     jsize alen = (*env)->GetArrayLength(env,barr);
//     jbyte* ba = (*env)->GetByteArrayElements(env,barr, JNI_FALSE);
//     if (alen > 0)
//     {
//         if(outlen==0)
//             return alen;
//         if(outlen<=alen)
//             return -1;
//         rtn=outbuf;
//         memcpy(rtn, ba, alen);
//         rtn[alen] = 0;
//     }
//     (*env)->ReleaseByteArrayElements(env,barr, ba, 0);
 
//     return alen;
// }
 
// //由指定编码以零结束的char转为Java String
// jstring pcharToJstring(JNIEnv* env, const char* pat, const char* encoding){
//     jstring jencoding;
//     jbyteArray bytes = (*env)->NewByteArray(env,strlen(pat));
//     (*env)->SetByteArrayRegion(env,bytes, 0, strlen(pat), (jbyte*)pat);
//     if (encoding==NULL)
//         jencoding= (*env)->NewStringUTF(env,"utf-8");
//     else
//         jencoding=(*env)->NewStringUTF(env,encoding);
 
//     return (jstring)(*env)->NewObject(env,gStringClass, gmidStringInit, bytes, jencoding);
// }
 
// //在C代码中执行字符串编码转换
// //参数分别为：原始字符，原始编码，目标字符缓冲区，目标编码，目标缓冲区大小，返加转换结果的长度
// int changeCharset(char * src_buf, char * src_encoding, char * dst_buf, char * dst_encoding, int dst_size){
//     JNIEnv *env;
//     jstring jtemp;
//     int res;
 
//     //由于初始化只执行了一次，本函数与初始JNI调用可能不在同一线程，因此需要判断当前线程
//     if(gJvmThread==pthread_self())
//     {
//         //如果是同一个线程，直接转
//         env=gJniEnv;
//         jtemp=pcharToJstring(env, src_buf, src_encoding);
//         res=jstringToPchar(env, jtemp,dst_encoding, dst_buf, dst_size);
//     }
//     else
//     {
//         //如果不是同一个线程，先Attach再转
//         env=gJniEnv;
//         (*gJavaVM)->AttachCurrentThread(gJavaVM,&env,NULL);
//         jtemp=pcharToJstring(env, src_buf, src_encoding);
//         res=jstringToPchar(env, jtemp,dst_encoding, dst_buf, dst_size);
//         (*gJavaVM)->DetachCurrentThread(gJavaVM);
//     }
//     return res;
// }

jstring charTojstring(JNIEnv* env, const char* pat) {
    (*env)->GetJavaVM(env, &gJavaVM);
    gJniEnv=env;
    gJvmThread=pthread_self();//记住当前JNI环境的线程    
    
    //定义java String类 strClass
    jclass strClass = (*env)->FindClass(env, "Ljava/lang/String;");
    //获取String(byte[],String)的构造器,用于将本地byte[]数组转换为一个新String
    jmethodID ctorID = (*env)->GetMethodID(env, strClass, "<init>", "([BLjava/lang/String;)V");
    //建立byte数组
    jbyteArray bytes = (*env)->NewByteArray(env, strlen(pat));
    //将char* 转换为byte数组
    (*env)->SetByteArrayRegion(env, bytes, 0, strlen(pat), (jbyte*) pat);
    // 设置String, 保存语言类型,用于byte数组转换至String时的参数
    jstring encoding = (*env)->NewStringUTF(env, "GB2312");
    //将byte数组转换为java String,并输出
    return (jstring) (*env)->NewObject(env, strClass, ctorID, bytes, encoding);
}

char* jstringToChar(JNIEnv* env, jstring jstr) {
    (*env)->GetJavaVM(env, &gJavaVM);
    gJniEnv=env;
    gJvmThread=pthread_self();//记住当前JNI环境的线程        

    char* rtn = NULL;
    jclass clsstring = (*env)->FindClass(env, "java/lang/String");
    jstring strencode = (*env)->NewStringUTF(env, "GB2312");
    jmethodID mid = (*env)->GetMethodID(env, clsstring, "getBytes", "(Ljava/lang/String;)[B");
    jbyteArray barr = (jbyteArray) (*env)->CallObjectMethod(env, jstr, mid, strencode);
    jsize alen = (*env)->GetArrayLength(env, barr);
    jbyte* ba = (*env)->GetByteArrayElements(env, barr, JNI_FALSE);
    if (alen > 0) {
        rtn = (char*) malloc(alen + 1);
        memcpy(rtn, ba, alen);
        rtn[alen] = 0;
    }
    (*env)->ReleaseByteArrayElements(env, barr, ba, 0);
    return rtn;
}
