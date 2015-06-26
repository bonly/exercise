#include <functional>
#include <cassert>
#include <iostream>

//typedef void (*func_t) (int);//不能将‘void process_integers(func_t)’的实参‘1’从‘int_processor’转换到‘func_t {aka void (*)(int)}’
typedef std::function<void(int)> func_t;

void process_integers(func_t f){
     f(100);
}	

void my_ints_func(int i){
	std::clog << i << std::endl;
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


bool is_triggered = false;
void set_function_object(func_t& f){
     int_processor fo(0, 200, is_triggered);
     f = fo; //fo 出函数时已不在作用域,但仍然是对的, 因为function用了move转移对像,这就是std::any的技巧
}

int main(){
  func_t f;
  set_function_object(f);

  process_integers(f); //还是有用的
  assert(is_triggered);

  process_integers(my_ints_func&);
}

