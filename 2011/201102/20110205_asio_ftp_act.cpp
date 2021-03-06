#include <iostream>
#include <string>
#include <vector>
#include <utility>
#include <cstdio>
#include <boost/asio.hpp>
#include <boost/asio/ip/tcp.hpp>
#include <boost/thread.hpp>


using std::cout;
using std::endl;
using std::vector;
using std::string;
using std::pair;
using namespace boost::asio;
using boost::asio::ip::tcp;
using boost::system::error_code;
using boost::system::system_error;
using boost::thread;


struct callable
{
private:
        
        string _hostname;
        
        string _port;
        
public:

        callable(const string& hostname, const string& port) : _hostname(hostname), _port(port)
        {
        }
        
        
        void operator()()
        {
                io_service ios;
                ip::tcp::resolver resolver(ios);
                ip::tcp::resolver::query query(_hostname, _port);
                ip::tcp::resolver::iterator it = resolver.resolve(query);
                ip::tcp::endpoint ep = *it;
                
                tcp::acceptor server(ios, ep);
                tcp::socket client(ios);
                server.accept(client);
                cout << "main(): client accepted from " << client.remote_endpoint().address() << endl;
                
                const int BUFLEN = 8192;
                vector<char> buf(BUFLEN);
                error_code error;
                int len = client.receive(buffer(buf, BUFLEN), 0, error);
                cout.write(buf.data(), len);
                cout << endl;
        }
};


int main()
{
        try
        {
                io_service ios;
                ip::tcp::resolver resolver(ios);
                ip::tcp::resolver::query query("183.60.126.26", "21"); ///ftp
                ip::tcp::resolver::iterator it = resolver.resolve(query);
                ip::tcp::endpoint endpoint = *it;
                
                ip::tcp::socket client(ios);
                client.connect(endpoint);
                const int BUFLEN = 1024;
                vector<char> buf(BUFLEN);
                
                error_code error;
                int len = client.receive(buffer(buf, BUFLEN), 0, error);
                cout.write(buf.data(), len);
                cout << endl;
        
                string request = "USER bonly\r\n";  ///
                cout << request;
                client.send(buffer(request, request.size()));
                len = client.receive(buffer(buf, BUFLEN), 0, error);
                cout.write(buf.data(), len);
                cout << endl;
                
                request = "PASS 111111\r\n";  ///
                cout << request;
                client.send(buffer(request, request.size()));
                len = client.receive(buffer(buf, BUFLEN), 0, error);
                cout.write(buf.data(), len);
                cout << endl;
                
                request = "PORT 0,0,0,0,195,80\r\n"; /// 195�256+80=50000
                cout << request;
                client.send(buffer(request, request.size()));
                len = client.receive(buffer(buf, BUFLEN), 0, error);
                cout.write(buf.data(), len);
                cout << endl;
                
                callable call("0.0.0.0", "50000"); /// 
                thread th(call);
                
                request = "RETR test.txt\r\n"; ///
                cout << request;
                client.send(buffer(request, request.size()));
                len = client.receive(buffer(buf, BUFLEN), 0, error);
                cout.write(buf.data(), len);
                cout << endl;
                th.join();
        }
        catch (system_error& exc)
        {
                cout << "main(): exc.what()=" << exc.what() << endl;
        }

        return EXIT_SUCCESS;
}
