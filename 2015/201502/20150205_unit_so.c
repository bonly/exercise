#ifdef __cplusplus
extern "C"{
#endif

const char* PrintHello(){
    return "hello";
}

int PrintANumber(){
    return 5;
}

int AddTwoInt(int a, int b){
    return a + b;
}

float AddTwoFloat(float a, float b){
    return a + b;
}

#ifdef __cplusplus
}
#endif

/*
gcc -fPIC -shared -o libfc.so fromc.c
*/