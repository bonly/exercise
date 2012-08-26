#include <iostream>
#include <string>
using namespace std;

#define dis(N,k) \
    clog << string(#N##k) << endl

int main(int argc, char* argv[])
{
    dis((Hello),(World));
    return 0;
}

