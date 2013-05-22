//============================================================================
// Name        : 4.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#include <omp.h>
#include <stdio.h>
using namespace std;


int main(int argc, char* argv[])
{
#pragma omp parallel for
     for (int i = 0; i < 10; i++ )
     {
         printf("i = %d\n", i);
     }
     return 0;
}

/*
int main (int argc, char *argv[])
{
	int id, nthreads;
  #pragma omp parallel private (id)
	{
			id = omp_get_thread_num ();
			printf ("Hello World from thread %d\n", id);
      #pragma omp barrier
			if (id == 0) {
					nthreads = omp_get_num_threads ();
					printf ("There are %d threads\n", nthreads);
			}
	}
	return 0;
}
*/

