#include <iostream>

using namespace std;

#define FUN_var(f) {cerr<<#f<<endl;}
#define FUN_val(f) {cerr<<f<<endl;}
#define FUN(f)   {int f=100; cerr << #f << ":" << f << endl;}
int main()
{
    char fname[]="myfun";
    FUN_var(fname);  ///< fname
    FUN_val(fname); ///< myfun
    FUN_var("myfun");  ///< "myfun"
    FUN_val("myfun");  ///< myfun

    const char* mf="myfun";
    FUN_var(mf);   ///< mf
    FUN_val(mf);  ///< myfun
    FUN(mf);  ///< mf:100

    return 0;
}

