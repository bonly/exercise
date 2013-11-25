/**
 * @file 20110514_obj.cpp
 * @brief 
 * @author bonly
 * @date 2013年10月30日 bonly Created
 */
#pragma GCC diagnostic ignored "-Wunused-local-typedefs"
#pragma GCC diagnostic ignored "-Wunused-variable"
#include <boost/thread.hpp>
#include <iostream>

namespace Bas{
using boost::thread;
using namespace std;
class TObj{
public:
    virtual void Run()=0;
    virtual void Start(){
        _thr = thread(&TObj::Run, this);
        _thr.join();
    }
    virtual ~TObj(){
        _thr.join();
    }
protected:
    thread _thr;
};

class Worker : public TObj{
public:
    virtual void Run(){
        for (int i = 0; i < 10; ++i){
            clog << "Running...\n";

        }
    }
};
}
using Bas::Worker;

int main(int argc, char* argv[]){
    Worker wk;
    wk.Start();
    return 0;
}
