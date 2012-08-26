#include <iostream>
#include <boost/timer.hpp>
using namespace boost;
using namespace std;
int main()
{
    timer t;
    cout << "min timespan: " << t.elapsed_min() << endl;
    cout << "max timespan: " << t.elapsed_max() << endl;
    cout << "now time elapsed: " << t.elapsed() << endl;
    return 0;
}
