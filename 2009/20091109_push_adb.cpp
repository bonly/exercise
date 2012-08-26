#include <cstdio>
#include <cstdlib>

int main(int argc, char* argv[])
{
    int const begin = atoi(argv[1]);
    int const end = atoi(argv[2]);
    char fil[50]="";
    char comm[255]="";
    for (int i=begin; i<=end; ++i)
    {
      sprintf(fil,"%03d.mp3",i);
      sprintf(comm, "adb push %s /mnt/sdcard/media/story/ ",fil);
      printf ("copying file %s...\n",fil); 
      system(comm);
    }
    return 0;
}
