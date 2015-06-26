#include <boost/asio/deadline_timer.hpp>
#include <boost/asio/io_service.hpp>
#include <boost/make_shared.hpp>

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

namespace detail{
typedef boost::asio::deadline_timer::duration_type duration_type;

template<typename Functor>
struct timer_task : public task_wrapped<Functor>{
private:
	typedef task_wrapped<Functor> base_t;
	boost::shared_ptr<boost::asio::deadline_timer> _timer;

public:
	template<typename Time>
	explicit timer_task(boost::asio::io_service& ios, 
		const Time& duration_or_time, 
		const Functor& task_unwrapped)
	: base_t(task_unwrapped), // call father's build
	  _timer(boost::make_shared<boost::asio::deadline_timer>(boost::ref(ios), duration_or_time)){
	}

	void push_task() const{
		_timer->async_wait(*this); // return immediately, but call (*this) after _timer's time up
	}

	void operator()(const boost::system::error_code& error) const{
		if(!error){
			base_t::operator()();
		}else{
			std::cerr << error << std::endl;
		}
	}
};
//}

//namespace detail{
template<typename Time, typename Functor>
inline timer_task<Functor> make_timer_task(
	boost::asio::io_service& ios, 
	const Time& duration_or_time, 
	const Functor& task_unwrapped){
	return timer_task<Functor>(ios, duration_or_time, task_unwrapped);
}
}

namespace tp_base{
class tasks_processor : private boost::noncopyable{
protected:
	boost::asio::io_service _ios;
	boost::asio::io_service::work _work;
	
	tasks_processor() : _ios(), _work(_ios){
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
}

namespace tp_timers{
class tasks_processor : public tp_base::tasks_processor{
public:
	static tasks_processor& get();

	typedef boost::asio::deadline_timer::duration_type duration_type;

	template<typename Functor>
	void run_after(duration_type duration, const Functor& task_unwrapped){
		detail::make_timer_task(_ios, duration, task_unwrapped).push_task();
	}

	typedef boost::asio::deadline_timer::time_type time_type;

	template<typename Functor>
	void run_at(time_type itime, const Functor& task_unwrapped){
		detail::make_timer_task(_ios, itime, task_unwrapped).push_task();
	}

};

tasks_processor& tasks_processor::get() {
    static tasks_processor proc;
    return proc;
}

}

