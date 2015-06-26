#include "20130118_net_multi.h"

const std::size_t threads_count = 5;

#include <boost/thread/barrier.hpp>
boost::barrier g_barrier(threads_count);

using namespace tp_multi;
void multythread_test() {
    g_barrier.wait();
    tasks_processor::get().stop();
}

int main() {
    for (std::size_t i = 0; i < threads_count; ++i) {
        tasks_processor::get().push_back(&multythread_test);
    }
    time_t t1 = time(NULL);
    tasks_processor::get().start_multiple(threads_count);
    time_t t2 = time(NULL);
    // One additional second for some io_service and OS delays
    assert(t2 - t1 < 1); 
}