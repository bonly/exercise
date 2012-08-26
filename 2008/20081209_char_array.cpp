#include <stdio.h>
#include <stdlib.h>
int main(void)
{

        char str[][10]={
                        "hello",
                        "world",
                        0
                        };
        /*bbs上的解法
        char str[][10]={
                        "hello",
                        "world",
                        0
                        };
        char **p = malloc(sizeof(char *)*2);
        p[0] = str[0];
        p[1] = str[1];

        printf("%s\n%s\n", p[0], p[1]);
        */
        
        for (int i=0; i<3; ++i)
        {
          printf("%s ",str[i]);
        }
        return 0;
}

