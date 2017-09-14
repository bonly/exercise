// int changeCharset(char * src_buf, char * src_encoding, char * dst_buf, char * dst_encoding, int dst_size);
// jstring pcharToJstring(JNIEnv* env, const char* pat, const char* encoding);
// int  jstringToPchar(JNIEnv* env, jstring jstr, const char * encoding, char* outbuf, int outlen);
// jstring Java_com_huz_test_CharsetTest_InitJNIEnv(JNIEnv* env, jobject obj);

jstring charTojstring(JNIEnv* env, const char* pat);
char* jstringToChar(JNIEnv* env, jstring jstr);
