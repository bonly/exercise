#ifdef __cplusplus
extern "C" {
#endif
 
char* gidToUid(unsigned long long ullGID);
 
unsigned long long uidToGid(const char* uid);
 
#ifdef __cplusplus
}
#endif