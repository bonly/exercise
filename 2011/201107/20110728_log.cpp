#include <thread>
#include <boost/log/core.hpp>
#include <boost/log/trivial.hpp>
#include <boost/log/expressions.hpp>
#include <boost/log/sinks/text_file_backend.hpp>
#include <boost/log/utility/setup/file.hpp>
#include <boost/log/utility/setup/common_attributes.hpp>
#include <boost/log/sources/severity_logger.hpp>
#include <boost/log/sources/record_ostream.hpp>
namespace logging = boost::log;
namespace src = boost::log::sources;
namespace sinks = boost::log::sinks;
namespace keywords = boost::log::keywords;
  
void InitLog() {
  // 必须注册格式，否则你看不到severity字段
  boost::log::register_simple_formatter_factory< boost::log::trivial::severity_level, char >("Severity");
  logging::add_file_log(
   keywords::file_name = "./sign_%Y-%m-%d_%H-%M-%S.%N.log",
   keywords::rotation_size = 10 * 1024 * 1024,
   keywords::time_based_rotation = sinks::file::rotation_at_time_point(0, 0, 0),
   keywords::format = "[%TimeStamp%] (%Severity%) : %Message%",
   keywords::min_free_space=3 * 1024 * 1024
   );
  logging::core::get()->set_filter(logging::trivial::severity >= logging::trivial::debug);
}

int main(){
    InitLog(); //调用InitLog之后，要设置属性
    logging::add_common_attributes();
    using namespace logging::trivial;
    src::severity_logger< severity_level > lg;
    BOOST_LOG_SEV(lg, info) << "thread id: " << std::this_thread::get_id() << " Initialization succeeded";
    BOOST_LOG_TRIVIAL(trace) << "test log";
    BOOST_LOG_TRIVIAL(debug) << L"debug test log";
    BOOST_LOG_TRIVIAL(debug) << L"debug test log 测试";
    return 0;
}      
/*
 在其他.cc文件中，只需要include一个头文件 <boost/log/trivial.hpp>  
并使用宏 BOOST_LOG_TRIVIAL()

g++ 20110728_log.cpp -std=c++11 -lboost_log -lboost_log_setup  -lboost_thread -lpthread -lboost_system -lboost_filesystem -DBOOST_LOG_DYN_LINK
*/

