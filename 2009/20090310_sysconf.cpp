#include <cstdio>
#include <cstdlib>
#include <unistd.h> //_SC_MQ_OPEN_MAX

int
main(int argc,char* argv[])
{
        printf("MQ_OPEN_MAX = %ld, MQ_PRIO_MAX = %ld\n", 
                        sysconf(_SC_MQ_OPEN_MAX),sysconf(_SC_MQ_PRIO_MAX));
        exit(0);
}

/*
 * MQ_OPEN_MAX һ��������ͬʱӵ�д򿪵���Ϣ���е������Ŀ
 * MQ_PRIO_MAX ������Ϣ��������ȼ�ֵ��1
 */
