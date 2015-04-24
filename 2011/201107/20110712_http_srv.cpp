#include <boost/bind.hpp>
#include <boost/thread.hpp>
#include <boost/network/utils/thread_pool.hpp>
#include <boost/network/include/http/server.hpp>
#include <boost/scoped_ptr.hpp>
#include <string>

class Handler;

typedef boost::network::http::async_server< Handler > Server;

class Handler
{
public:
    void operator()(const Server::request& request, Server::connection_ptr connection) { /*stuff*/ }
};

class ServerFacade
{
    typedef boost::scoped_ptr< boost::network::utils::thread_pool > ThreadPoolPtr;
    typedef boost::scoped_ptr< Server > ServerPtr;
    typedef boost::scoped_ptr< boost::thread > ThreadPtr;

public:
    ServerFacade(const std::string& ip_address, const std::string& port_number, unsigned int thread_count)
    :   server_thread_pool_(new boost::network::utils::thread_pool(thread_count))
    ,   server_(new Server(
            boost::network::http::address = ip_address,
            boost::network::http::port = port_number,
            boost::network::http::handler = handler_,
            boost::network::http::thread_pool = *server_thread_pool_,
            boost::network::http::reuse_address = true
        ))
    ,   server_thread_(new boost::thread(
            boost::bind(&Server::run, server_.get())
        ))
    {
        std::cout << "Server established at " << ip_address << ":" << port_number << " with " << thread_count << " threads." << std::endl;
    }

    ~ServerFacade()
    {
        server_->stop();

        server_thread_->join();

        server_thread_pool_.reset();
        server_.reset();
        server_thread_.reset();
    }

private:
    Handler handler_;
    ThreadPoolPtr server_thread_pool_;
    ServerPtr server_;
    ThreadPtr server_thread_;
};

int main(int argc, char** argv)
{
    {
        ServerFacade("127.0.0.1", "12345", 1); // <-- here io_service::run is invoked in another thread
    } // <-- here we close the acceptor, call stop on async_server, and join the io_service::run thread (and get stuck)
    return 0; // <-- we never get to here
}
