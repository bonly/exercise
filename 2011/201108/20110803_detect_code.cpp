#include "detect_code.h"
#include <chardet.h>
//char out_encode[CHARDET_MAX_ENCODING_NAME]
#include <iostream>

char* GetLocalEncoding(const char* in_str, unsigned int str_len, char* out_encode){
        chardet_t chardect=NULL;
        if(chardet_create(&chardect)==CHARDET_RESULT_OK){
                if(chardet_handle_data(chardect, in_str, (unsigned int)str_len) == CHARDET_RESULT_OK)
                        if(chardet_data_end(chardect) == CHARDET_RESULT_OK)
                                chardet_get_charset(chardect, out_encode, CHARDET_MAX_ENCODING_NAME);
        }
        if(chardect)
                chardet_destroy(chardect);
        return out_encode;
}


/*
int main(){
   char test[]="这是一个测试";
   char out[255]="";
   std::clog << GetLocalEncoding(test, sizeof(test), out); 
}
*/
/*
g++ 20110802_detect_code.cpp -I ~/opt/libchardet-0.0.4/src/ -L ~/opt/libchardet-0.0.4/src/.libs/ -lchardet
*/

