#include <boost/thread/thread.hpp>
#include <iostream>

namespace detail{

template <typename T>
struct task_wrapped{
private:
	T _task_unwrapped;

public:
	explicit task_wrapped(const T& task_unwrapped)
	:_task_unwrapped(task_unwrapped){}

	void operator()() const{
		try{//resetting interruption
			boost::this_thread::interruption_point();
		}catch(const boost::thread_interrupted&){}

		try{
			_task_unwrapped(); //executing task
		}catch(const std::exception& e){
			std::cerr << "Exception: " << e.what() << std::endl;
		}catch(const boost::thread_interrupted& e){
			std::cerr << "Thread interruption" << std::endl;
		}catch(...){
			std::cerr << "Unknow exception" << std::endl;
		}
	}
};

template<typename T>
task_wrapped<T> make_task_wrapped(const T& task_unwrapped){
	return task_wrapped<T>(task_unwrapped);
}
}
