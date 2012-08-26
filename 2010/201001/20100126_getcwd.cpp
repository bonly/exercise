#include <iostream>
using namespace std;
int main()
{
    char buf[255];

    cout << "PWD: " << getenv("PWD") << endl;
    getcwd(buf,255);
    cout << "getcwd: " << buf << endl;

    clog << chdir("/home/") << endl;
    cout << getenv("PWD") << endl;
    cout << getcwd(buf,255) << endl;

    return 0;
}
