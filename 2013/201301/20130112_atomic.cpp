#include <boost/atomic.hpp>
//#include <typeinfo>
#include <boost/typeof/typeof.hpp>
#include <boost/thread.hpp>
#include <cassert>
#include <cstddef>

boost::atomic<int> shared_i(0);

void do_inc(){
	for (std::size_t i=0; i<3000; ++i){
		const int i_snapshot = ++shared_i;
	}
}

void do_dec(){
	for (std::size_t i=0; i<3000; ++i){
		const int i_snapshot = --shared_i;
	}
}

int main(){
	boost::thread t1(&do_inc);
	boost::thread t2(&do_dec);
	t1.join();
	t2.join();
	assert(shared_i==0);
	std::clog << "shared_i = " << shared_i << std::endl;
	std::clog << typeid(shared_i).name() << std::endl;
	return 0;
}

