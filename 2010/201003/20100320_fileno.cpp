#include <cstdio>

int main()
{
    FILE* file = fopen("/tmp/bac", "wr");
    int intfile = fileno(file); //把文件stream转换为fileno
    fclose(file);
    return 0;
}
