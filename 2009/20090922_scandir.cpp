#include <sys/types.h>
#include <stdio.h>
#include <dirent.h>
#include <stdlib.h>

extern int scandir();
extern int alphasort();

main()
{
	int num_entries, i;
	struct dirent **namelist, **list;

	if ((num_entries =
		scandir("/tmp", &namelist, NULL, alphasort)) < 0) {
	   fprintf(stderr, "Unexpected error\n");
	   exit(1);
	}
	printf("Number of entries is %d\n", num_entries);
	if (num_entries) {
	   printf("Entries are:");
	   for (i=0, list=namelist; i<num_entries; i++) {
	       printf(" %s", (*list)->d_name);
	       free(*list);
	       list++;
	   }
		      }
	free(namelist);
	printf("\n");
	exit(0);
}