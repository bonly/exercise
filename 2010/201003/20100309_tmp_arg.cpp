#include <iostream>
using namespace std;

template <typename WORK>
class handle
{
    public:
        typedef WORK work_type;
        void check(work_type *wk)
        {
            cout << "in handle" << endl;
            wk->finish_work();
            cout << "back handle" <<endl;
        }

};

class work
{
    public:
        typedef handle<work> handle_type;
        void on_check(handle_type &handle)
        {
            cout << "in work"<<endl;
            handle.check(this);
        }
        void finish_work()
        {
            cout << "finish work" <<endl;
        }
    
};

int main()
{
    work wk;
    handle<work> hd;
    wk.on_check(hd);

}

