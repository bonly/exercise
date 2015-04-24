#include <boost/network/protocol/http/server.hpp>
#include <iostream>


namespace http = boost::network::http;


/*<< Defines the server. >>*/
struct hello_world;
typedef http::server<hello_world> server;

/*<< Defines the request handler.  It's a class that defines two
     functions, `operator()` and `log()` >>*/
struct hello_world {
    /*<< This is the function that handles the incoming request. >>*/
    void operator() (server::request const &request,
                     server::response &response) {
        server::string_type ip = source(request);
        unsigned int port = request.source_port;
        std::ostringstream data;
        data << "Hello, " << ip << ':' << port << '!';
        response = server::response::stock_reply(
            server::response::ok, data.str());
        std::clog << (std::string)body(request);
    }
    /*<< It's necessary to define a log function, but it's ignored in
         this example. >>*/
    void log(...) {
        std::clog << "a connect";
    }
};


int main(int argc, char * argv[]) {
    
    if (argc != 3) {
        std::cerr << "Usage: " << argv[0] << " address port" << std::endl;
        return 1;
    }

    try {
        /*<< Creates the request handler. >>*/
        hello_world handler;
        /*<< Creates the server. >>*/
        server::options options(handler);
        server server_(options.address(argv[1]).port(argv[2]));
        /*<< Runs the server. >>*/
        server_.run();
    }
    catch (std::exception &e) {
        std::cerr << e.what() << std::endl;
        return 1;
    }
    
    return 0;
}
//]

//g++ 20110708_srv_parse.cpp -I ~/opt/cpp-netlib/ -L ~/opt/cpp-netlib/./mybuild/libs/network/src/ -lcppnetlib-server-parsers -lpthread -lboost_system
