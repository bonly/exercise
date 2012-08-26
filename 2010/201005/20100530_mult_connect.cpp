/**
 * @file 20100530_mult_connect.cpp
 * @brief
 *
 * @author bonly
 * @date 2012-7-17 bonly created
 */
#include <cstdlib>
#include <cstring>
#include <iostream>
#include <boost/asio.hpp>
#include <boost/bind.hpp>
using boost::asio::ip::tcp;
using namespace std;
static int id = 1;
const char message[] = "test write string...";
class echo_session
{
    public:
        echo_session(boost::asio::io_service& io_service) :
                    socket_(io_service),io_(io_service)
        {
            id_ = id;
            ++id;
        }
        void start(const std::string& ip, const std::string& port)
        {
            //解析主机地址
            tcp::resolver resolver(io_);
            tcp::resolver::query query(tcp::v4(), ip, port);
            tcp::resolver::iterator iterator = resolver.resolve(query);
            //异步连接
            socket_.async_connect(
                        *iterator,
                        boost::bind(&echo_session::handle_connect, this,
                                    boost::asio::placeholders::error));
        }
    private:
        void handle_connect(const boost::system::error_code& error)
        {
            if (!error)
            {
                //连接成功，发送message中的数据
                boost::asio::async_write(
                            socket_,
                            boost::asio::buffer(message, sizeof(message)),
                            boost::bind(&echo_session::handle_write, this,
                                        boost::asio::placeholders::error));
            }
            else
                cout << error << endl;
        }
        void handle_write(const boost::system::error_code& error)
        {
            if (!error)
            {
                //写入完毕，接收服务器回射的消息
                boost::asio::async_read(
                            socket_,
                            boost::asio::buffer(buf_, sizeof(buf_)),
                            boost::bind(
                                        &echo_session::handle_read,
                                        this,
                                        boost::asio::placeholders::error,
                                        boost::asio::placeholders::bytes_transferred));
            }
            else
                cout << error << endl;
        }
        void handle_read(const boost::system::error_code& error,
                    size_t bytes_transferred)
        {
            if (!error)
            {
                //读取完毕，在终端显示
                cout << id_ << ":receive:" << bytes_transferred << "," << buf_
                            << endl;
                //周而复始...
                handle_connect(error);
            }
            else
                cout << error << endl;
        }
        int id_;
        tcp::socket socket_;
        boost::asio::io_service &io_;
        char buf_[sizeof(message)];
};
int main(int argc, char* argv[])
{
    const int session_num = 10000; //连接的数量
    echo_session* sessions[session_num];
    memset(sessions, 0, sizeof(sessions));
    try
    {
        if (argc != 3)
        {
            std::cerr << "Usage: blocking_tcp_echo_client <host> <port>/n";
            return 1;
        }
        boost::asio::io_service io_service;
        //创建session_num个连接
        for (int i = 0; i < session_num; ++i)
        {
            sessions[i] = new echo_session(io_service);
            sessions[i]->start(argv[1], argv[2]);
        }
        //io_service主循环
        io_service.run();
        for (int i = 0; i < session_num; ++i)
            if (sessions[i] != NULL)
                delete sessions[i];
    }
    catch (std::exception& e)
    {
        for (int i = 0; i < session_num; ++i)
            if (sessions[i] != NULL)
                delete sessions[i];
        std::cerr << "Exception: " << e.what() << "/n";
    }
    return 0;
}

