#include<stdio.h>
#include<stdlib.h>
int main()
{
	        int *ptr = (int *)0;
	        *ptr = 100;
		return 0;
}


/*
 * ulimit -c unlimited
 * /proc/sys/kernel/core_uses_pid
 * /proc/sys/kernel/core_pattern
 * /corefile/core-%e-%p-%t
 */
