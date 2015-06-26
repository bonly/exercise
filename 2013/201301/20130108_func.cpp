#include <functional>
#include <cassert>

//typedef void (*func_t) (int);//不能将‘void process_integers(func_t)’的实参‘1’从‘int_processor’转换到‘func_t {aka void (*)(int)}’
typedef std::function<void(int)> func_t;

void process_integers(func_t f){
     f(100);
}	

class int_processor: public std::unary_function<int, void> {
    const int _min;
    const int _max;
    bool& _triggered;

public:
    int_processor(int min, int max, bool& triggered):_min(min),_max(max),_triggered(triggered){
    }

    void operator()(int i) const{
	    if(i > _min && i < _max){
		    _triggered = true;
	    }
    }
};


int main(){
  bool is_triggered = false;
  int_processor fo(0, 200, is_triggered);
  process_integers(fo);
  assert(is_triggered);
}

