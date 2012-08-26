//============================================================================
// Name        : save.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#include <iostream>
#include <time.h>

struct MyData
{
	public:
		int number;
		char check[20];

	 MyData():number(0)
	 {
		 memset (check, 0, 20);
	 }
};
int
main()
{
	clock_t tick;
	tick = clock();
	double t = (double)tick/CLK_TCK; //精确到毫秒
	printf ("Total time: %f seconds\n",t);

	time_t start,stop;
	time (&start);
	double timeused;
	time (&stop);
	timeused = difftime (stop, start);//精确到秒
	printf ("Total time: %f seconds\n",timeused/CLK_TCK);

	clock_t mstart=0,mstop=0;
	mstart = clock ();
  long i=10000000L;
  while (i--);
	mstop = clock ();
	double use = (double)(mstop - mstart)/CLOCKS_PER_SEC;
	printf ("\nm second: %f \n",use);

	return 0;
}

