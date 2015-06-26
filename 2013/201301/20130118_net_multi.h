#include "20130117_net_task.h"

namespace tp_multi{
class tasks_processor : public tp_network::tasks_processor{
public:
	static tasks_processor& get(){
		static tasks_processor proc;
		return proc;
	}

	void start_multiple(std::size_t thread_count=0){
		if (!thread_count){
			thread_count = (std::max)(static_cast<int>(boost::thread::hardware_concurrency()), 1); //min is 1
		}

		--thread_count;

		boost::thread_group tg;
		for (std::size_t i=0; i<thread_count; ++i){
			tg.create_thread(boost::bind(&boost::asio::io_service::run, boost::ref(_ios)));
		}
		_ios.run();
		tg.join_all();
	}
};
}