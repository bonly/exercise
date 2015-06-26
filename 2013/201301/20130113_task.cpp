#include <deque>
#include <boost/function.hpp>
#include <boost/thread.hpp>
#include <boost/thread/mutex.hpp>
#include <boost/thread/locks.hpp>
#include <boost/thread/condition_variable.hpp>
#include <boost/bind.hpp>

class work_queue{
public:
	typedef boost::function<void()> task_type;

private:
	std::deque<task_type> _tasks;
	boost::mutex          _tasks_mutex;
	boost::condition_variable _cond;

public:
	void push_task(const task_type& task){
		boost::unique_lock<boost::mutex> lock(_tasks_mutex);
		_tasks.push_back(task);
		lock.unlock();
		_cond.notify_one();  //notify the blocking lock
	}
	task_type try_pop_task(){  //non-block
		task_type ret;
		boost::lock_guard<boost::mutex> lock(_tasks_mutex);
		if (!_tasks.empty()){
			ret = _tasks.front();
			_tasks.pop_front();
		}
		return ret;
	}
	task_type pop_task(){ //block
		boost::unique_lock<boost::mutex> lock(_tasks_mutex);
		while (_tasks.empty()){ //must use while to check the condition
			_cond.wait(lock); //must use unique_lock, wait will unlock the "lock" and wait for condition
		}
		task_type ret = _tasks.front();
		_tasks.pop_front();
		return ret;
	}
};

work_queue g_queue;

void do_nothing(){
}

const std::size_t tests_tasks_count = 3000;

void pusher(){
	for (std::size_t i=0; i<tests_tasks_count; ++i){
		g_queue.push_task(&do_nothing);
	}
} 

void popper_sync(){
	for (std::size_t i=0; i<tests_tasks_count; ++i){
		g_queue.pop_task() //get task
		();//exec 
	}
}

int main(){
	boost::thread pop_sync1(&popper_sync);
	boost::thread pop_sync2(&popper_sync);
	boost::thread pop_sync3(&popper_sync);

	boost::thread push1(&pusher);
	boost::thread push2(&pusher);
	boost::thread push3(&pusher);

	pop_sync1.join();
	pop_sync2.join();
	pop_sync3.join();

	push1.join();
	push2.join();
	push3.join();

	assert(!g_queue.try_pop_task());//check if empty with non-block...will failed

	g_queue.push_task(&do_nothing);

	assert(g_queue.try_pop_task());
}
