#include <boost/asio/io_service.hpp>
#include "20130114_task.cpp"

class tasks_processor : private boost::noncopyable{
protected:
	boost::asio::io_service _ios;
	boost::asio::io_service::work _work;

	tasks_processor():_ios(), _work(_ios){
	}

public:
	static tasks_processor& get();

	template<typename T>
	inline void push_back(const T& task_unwrapped){
		_ios.post(detail::make_task_wrapped(task_unwrapped));
	}

	void start(){
		_ios.run();
	}

	void stop(){
		_ios.stop();
	}
};

tasks_processor& tasks_processor::get() {
    static tasks_processor proc;
    return proc;
}

int g_val = 0;
void func_test(){
	++g_val;
	if (g_val == 3){
		throw std::logic_error("Just checking");
	}

	boost::this_thread::interruption_point();
	if (g_val == 10){ //Emulation of thread interruption
		throw boost::thread_interrupted();  //Will be caught and won't stop execution
	}
	if (g_val == 90){
		tasks_processor::get().stop();
	}
}

int main(){
	static const std::size_t tasks_count = 100;
	// stop() is called at 90
	BOOST_STATIC_ASSERT(tasks_count > 90);
	for (std::size_t i=0; i<tasks_count; ++i){
		tasks_processor::get().push_back(&func_test);
	}

	// we can also use result of boost::bind call as a task
	tasks_processor::get().push_back(boost::bind(std::plus<int>(), 2, 2));

	//processing was not started.
	assert(g_val == 0);

	// will not throw, but blocks till one of the tasks it is owning call stop().
	tasks_processor::get().start();
	assert(g_val == 90);

	return 0;
}
