#include <boost/asio.hpp>

#include <boost/exception_ptr.hpp>
#include <boost/lexical_cast.hpp>
#include "20130114_task.cpp"

namespace tp_base{
class tasks_processor : private boost::noncopyable{
protected:
    boost::asio::io_service _ios;
    boost::asio::io_service::work _work;

    tasks_processor():_ios(), _work(_ios){
    }

public:
    static tasks_processor& get();

    template<typename T>
    inline void push_task(const T& task_unwrapped){
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


using namespace tp_base;

// Part of tasks_processor class from
// tasks_processor_base.hpp, that must be defined
// Somewhere in source file
// tasks_processor& tasks_processor::get() {
//     static tasks_processor proc;
//     return proc;
// }

void func_test2(); // Forward declaration

void process_exception(const boost::exception_ptr& exc) {
    try {
        boost::rethrow_exception(exc);
    } catch (const boost::bad_lexical_cast& /*e*/) {
        std::cout << "Lexical cast exception detected\n" << std::endl;

        // Pushing another task to execute
        tasks_processor::get().push_task(&func_test2);
    } catch (...) {
        std::cout << "Can not handle such exceptions:\n" 
            << boost::current_exception_diagnostic_information() 
            << std::endl;

        // Stopping
        tasks_processor::get().stop();
    }
}

void func_test1() {
    try {
        boost::lexical_cast<int>("oops!");
    } catch (...) {
        tasks_processor::get().push_task(boost::bind(
            &process_exception, boost::current_exception()
        ));
    }
}

#include <stdexcept>
void func_test2() {
    try {
        // Some code goes here
        BOOST_THROW_EXCEPTION(std::logic_error("Some fatal logic error"));
        // Some code goes here
    } catch (...) {
        tasks_processor::get().push_task(boost::bind(
            &process_exception, boost::current_exception()
        ));
    }
}

void run_throw(boost::exception_ptr& ptr) {
    try {
        // A lot of code goes here
    } catch (...) {
        ptr = boost::current_exception();
    }
}

int main () {
    tasks_processor::get().push_task(&func_test1);
    tasks_processor::get().start();


    boost::exception_ptr ptr;
    // Do some work in parallel
    boost::thread t(boost::bind(
        &run_throw, 
        boost::ref(ptr)
    ));

    // Some code goes here
    // ...
    
    t.join();

    // Chacking for exception
    if (ptr) {
        // Exception occured in thread
        boost::rethrow_exception(ptr);
    }
}
