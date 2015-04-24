#pragma GCC diagnostic ignored "-Wunused-variable"
#pragma GCC diagnostic ignored "-Wunused-local-typedefs"
#include <boost/network/protocol/http/server.hpp>
#include <iostream>
#include <boost/bind.hpp>
#include <boost/thread.hpp>

namespace http = boost::network::http;


/*<< Defines the server. >>*/
struct hello_world;
typedef http::async_server<hello_world> server;

/*<< Defines the request handler.  It's a class that defines two
     functions, `operator()` and `log()` >>*/
struct hello_world {
    /*<< This is the function that handles the incoming request. >>*/
    void operator() (server::request const &req,
                     server::connection_ptr cont) {
        server::string_type ip = source(req);
        unsigned int port = req.source_port; //?
        //std::ostringstream data;
        //data << "Hello, " << ip << ':' << port << '!';
        char data[255]="";
        sprintf(data, "Hello %s:%d!", ip.c_str(), port);
        strcat(data, "Hello world from bonly");
        //cont->write(data.str());
        cont->write(data);
        std::clog << (std::string)body(req);
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
        //boost::thread t1(boost::bind(&server::run, &server_));
        //server_.run();
        //t1.join();
        
        boost::thread_group thr;
        thr.create_thread(boost::bind(&server::run, &server_));
        thr.join_all();
        server_.cloer();
    }
    catch (std::exception &e) {
        std::cerr << e.what() << std::endl;
        return 1;
    }
    
    return 0;
}
//]

//g++ 20110708_srv_parse.cpp -I ~/opt/cpp-netlib/ -L ~/opt/cpp-netlib/./mybuild/libs/network/src/ -lcppnetlib-server-parsers -lpthread -lboost_system
