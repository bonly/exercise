#include <cstdio>

struct Pack
{
    int len;
    int cmd;
};

int main()
{
    printf("short int: %d\n", sizeof(short int));
    printf("int: %d\n", sizeof(int));
    printf("long: %d\n", sizeof(long));
    printf("usinged int: %d\n", sizeof(unsigned int));
    printf("usinged long: %d\n", sizeof(unsigned long));
    printf("char: %d\n", sizeof(char));
    printf("Pack: %d\n", sizeof(Pack));
    return 0;
}
