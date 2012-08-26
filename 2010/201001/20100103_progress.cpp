#include <boost/progress.hpp>
#include <iostream>
using namespace std;
int main()
{
    clog << "program begin...\n";
    {
        boost::progress_timer t;
        clog << "progrees ...." << endl;
    }
    return 0;
}
