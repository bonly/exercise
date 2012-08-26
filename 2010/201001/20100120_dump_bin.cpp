#include <cstdlib>
#include <string.h>
#include <cstdio>
#include <ctype.h>

void _hex(unsigned char c, char* s)
{
    sprintf(s, "%02X", c);
    s[2] = ' ';
}

char* _hexdump16bytes(char* p, size_t& nBegin, size_t& nLeft)
{
    if(nLeft <= 0)
    {
        return NULL;
    }

    enum{line_size = 67};
    static char s[line_size + 1];
    s[line_size] = '\0';

    char* p2 = p + nBegin;
    for(size_t i=0; i<16; ++i)
    {
        if(nLeft > 0)
        {
            unsigned char c = p2[i];
            _hex(c, s+i*3);

            s[51+i] = isprint(c) ? c : '.';
            --nLeft;
        }
        else
        {
            size_t j=i*3;
            s[j] = ' ';
            s[j+1] = ' ';
            s[j+2] = ' ';
            s[51+i] = ' ';
        }
    }

    s[48] = ' ';
    s[49] = ';';
    s[50] = ' ';

    nBegin += 16;

    return s;
}

void hexdump(char* buf, const size_t len)
{
    size_t i=0;
    size_t j=len;

    char* s = 0;
    while((s=_hexdump16bytes(buf,i,j))!=0)
    {
        //clog << s << "\n";
        //clog.write(p, sizeof(mml::MML));
        printf("%s\n", s);
    }
}


int main()
{
    char p[]="this is a test\n";
    hexdump(p, strlen(p));
    return 0;
}
