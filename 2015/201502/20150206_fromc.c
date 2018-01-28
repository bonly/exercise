#ifdef __cplusplus
extern "C" {
#endif
 
char* gidToUid(unsigned long long ullGID){
	return "hello";
}
 
unsigned long long uidToGid(const char* uid){
	return 0;
}
 
#ifdef __cplusplus
}
#endif
/*
 gcc -fPIC -shared -o libfc.so fromc.c
*/
 
