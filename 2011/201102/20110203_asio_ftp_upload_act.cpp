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
                ip::tcp::endpoint endpoint = *it;
                
                tcp::acceptor server(ios, endpoint);
                tcp::socket client(ios);
                server.accept(client);
                cout << "main(): client accepted from " << client.remote_endpoint().address() << endl;

                string content = "Hello, World!";
                client.send(buffer(content, content.size()));
                cout << "Content sent" << endl;
        }
};


int main()
{
        try
        {
                io_service ios;
                ip::tcp::resolver resolver(ios);
                ip::tcp::resolver::query query("ftp.alepho.com", "21");
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
        
                string request = "USER ***\r\n";
                cout << request;
                client.send(buffer(request, request.size()));
                len = client.receive(buffer(buf, BUFLEN), 0, error);
                cout.write(buf.data(), len);
                cout << endl;
                
                request = "PASS ***\r\n";
                cout << request;
                client.send(buffer(request, request.size()));
                len = client.receive(buffer(buf, BUFLEN), 0, error);
                cout.write(buf.data(), len);
                cout << endl;
                
                request = "PORT 192,168,1,102,195,80\r\n";
                cout << request;
                client.send(buffer(request, request.size()));
                len = client.receive(buffer(buf, BUFLEN), 0, error);
                cout.write(buf.data(), len);
                cout << endl;
                
                callable call("192.168.1.102", "50000");
                thread th(call);
                
                request = "STOR alepho.com/public_html/hello.txt\r\n";
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
