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
  
// Declare attribute keywords
BOOST_LOG_ATTRIBUTE_KEYWORD(severity, "Severity", severity_level)
BOOST_LOG_ATTRIBUTE_KEYWORD(timestamp, "TimeStamp", boost::posix_time::ptime)

void init_logging()
{
    boost::shared_ptr< sinks::synchronous_sink< sinks::text_file_backend > > sink = logging::add_file_log
    (
        "sample.log",
        keywords::format = expr::stream
            << expr::format_date_time(timestamp, "%Y-%m-%d, %H:%M:%S.%f")
            << " <" << severity.or_default(normal)
            << "> " << expr::message
    );

    // The sink will perform character code conversion as needed, according to the locale set with imbue()
    std::locale loc = boost::locale::generator()("en_US.UTF-8");
    sink->imbue(loc);

    // Let's add some commonly used attributes, like timestamp and record counter.
    logging::add_common_attributes();
}

void test_narrow_char_logging()
{
    // Narrow character logging still works
    src::logger lg;
    BOOST_LOG(lg) << "Hello, World! This is a narrow character message.";
}

void test_wide_char_logging()
{
    src::wlogger lg;
    BOOST_LOG(lg) << L"Hello, World! This is a wide character message.";

    // National characters are also supported
    const wchar_t national_chars[] = { 0x041f, 0x0440, 0x0438, 0x0432, 0x0435, 0x0442, L',', L' ', 0x043c, 0x0438, 0x0440, L'!', 0 };
    BOOST_LOG(lg) << national_chars;

    // Now, let's try logging with severity
    src::wseverity_logger< severity_level > slg;
    BOOST_LOG_SEV(slg, normal) << L"A normal severity message, will not pass to the file";
    BOOST_LOG_SEV(slg, warning) << L"A warning severity message, will pass to the file";
    BOOST_LOG_SEV(slg, error) << L"An error severity message, will pass to the file";
}

int main(){
    init_logging(); //调用InitLog之后，要设置属性

    using namespace logging::trivial;
    src::severity_logger< severity_level > lg;
    BOOST_LOG_SEV(lg, info) << "thread id: " << std::this_thread::get_id() << " Initialization succeeded";
    BOOST_LOG_TRIVIAL(trace) << "test log";
    BOOST_LOG_TRIVIAL(debug) << L"debug test log";
    BOOST_LOG_TRIVIAL(debug) << L"debug test log 测试";
    
    test_wide_char_logging();
    test_narrow_char_logging();
    return 0;
}      
/*
 在其他.cc文件中，只需要include一个头文件 <boost/log/trivial.hpp>  
并使用宏 BOOST_LOG_TRIVIAL()

g++ 20110728_log.cpp -std=c++11 -lboost_log -lboost_log_setup  -lboost_thread -lpthread -lboost_system -lboost_filesystem -DBOOST_LOG_DYN_LINK
*/

