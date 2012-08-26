#include <string.h>
#include <cstdio>
int get_string(const char* str, char delimit, char* outstr, int* outlen)
{
    char* p = strchr((char*) str, delimit);
    if (NULL == p)
    {
        return -1;
    }

    memcpy(outstr, str, p - str);
    *outlen = p - str + 1;
    return 0;
}

int proto(char *buf)
{
    char tmparr[20][32];
    memset(tmparr, 0, 20 * 32);
    char* p = buf;

    //protocol->result_code = 1;

    /*
     if('1' != protocol->header.encrypt_flag){
     protocol->result_code = 1;
     snprintf(protocol->errmsg, sizeof(protocol->errmsg), "Count=0&MSG=NotEncrypt");
     create_response_helper(protocol, outmsg, outlen);
     return -1;
     }
     */

    /*
     char* p = (char*)inmsg + PROTOCOL_HEAD_LEN;


     PRO_DEBUG("random_code:%s datalen = %d\n", protocol->random_code, datalen);
     //if(strcmp(protocol->random_code, "1234567890") != 0)
     //    return -1;

     Encrypt::crypt(p, datalen, protocol->random_code, strlen(protocol->random_code));

     if(LOG_CHECK_DEBUG(g_r5_plog)){
     char tmp[1024] = {0};
     memcpy(tmp, p, datalen);
     PRO_DEBUG("dump query_request original string:\n");
     PRO_DEBUG("%s\n", p);
     }

     */
    int datalen = 255;
    int tmplen = 0;
    char delimit = '=';

    for (int i = 0; i <= 16; ++i)
    {
        if (get_string(p, '=', tmparr[i], &tmplen) < 0)
        {
            //snprintf(protocol->errmsg, sizeof(protocol->errmsg), "Count=0&MSG=RequestError");
            //create_response_helper(protocol, outmsg, outlen);
            //PRO_WARN("split 1 faield, str = %s\n", p);
            printf("error\n");
            return -1;
        }

        p += tmplen;
        datalen -= tmplen;

        i++;

        if (get_string(p, '&', tmparr[i], &tmplen) < 0)
        {
            break;
        }

        p += tmplen;
        datalen -= tmplen;
    }

    if (get_string(p, '=', tmparr[18], &tmplen) < 0)
    {
        //snprintf(protocol->errmsg, sizeof(protocol->errmsg), "Count=0&MSG=RequestError");
        //create_response_helper(protocol, outmsg, outlen);
        //PRO_DEBUG("split 3 faield, str = %s\n", p);
        printf("=err\n");
        return -1;
    }

    p += tmplen;
    datalen -= tmplen;

    memcpy(tmparr[19], p, datalen);
    return 0;
}

int main()
{
    char buf[]= "seqno=000000000000000001&SubNo=13700414106&switch_flag=04&Brand=A&BeginDate=20110513091720&EndDate=20110724191720&QueryTime=20110829142730"; 
    return proto(buf);

}

