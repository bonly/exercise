#include <time.h>
#include <iostream>
#include <ctime>
#include <sys/time.h>
#include <unistd.h>
using namespace std;
//CLOCKS_PER_SEC

int main()
{
    /*
    clock_t tick;
    tick = clock();
    double t = (double)tick/CLK_TCK; 
    printf ("Total time: %f seconds\n",t);
    */

    /*
    pollfd pdf;
    pdf.fd = msqid;
    pdf.events = POLLIN;
    int ret = poll(&pdf, 1, 5000L);
    if (ret == 0) //超时
    {
        return -1;
    }
    else if (ret == 1) //有数据
    {
        int nLen = msgrcv(msqid, msgp, msgsz, msgtyp, msgflg);
        return (nLen>0)?nLen:0;
    }
    return 0;
    */
    
    clock_t c_begin,c_end;
    timeval t_begin,t_end;
    
    c_begin = clock();
    gettimeofday(&t_begin,0);

    sleep(5);

    c_end = clock();
    gettimeofday(&t_end,0);

    cout << "diffence clock: " << c_end - c_begin << endl;
    cout << "diffence time: " << t_end.tv_sec - t_begin.tv_sec << endl;


    return 0;


}

