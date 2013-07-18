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

                ip::tcp::socket client(ios);
                client.connect(endpoint);
                string content = "Hello, World!";
                client.send(buffer(content, content.size()));
                cout << "Content sent" << endl;
        }
};


pair<string, string> parseHostPort(string s)
{
        size_t paramsPos = s.find('(');
        string params = s.substr(paramsPos + 1);
        size_t ip1Pos = params.find(',');
        string ip1 = params.substr(0, ip1Pos);
        size_t ip2Pos = params.find(',', ip1Pos + 1);
        string ip2 = params.substr(ip1Pos + 1, ip2Pos - ip1Pos - 1);
        size_t ip3Pos = params.find(',', ip2Pos + 1);
        string ip3 = params.substr(ip2Pos + 1, ip3Pos - ip2Pos - 1);
        size_t ip4Pos = params.find(',', ip3Pos + 1);
        string ip4 = params.substr(ip3Pos + 1, ip4Pos - ip3Pos - 1);
        size_t port1Pos = params.find(',', ip4Pos + 1);
        string port1 = params.substr(ip4Pos + 1, port1Pos - ip4Pos - 1);
        size_t port2Pos = params.find(')', port1Pos + 1);
        string port2 = params.substr(port1Pos + 1, port2Pos - port1Pos - 1);
        
        pair<string, string> hostPort;
        hostPort.first = ip1 + "." + ip2 + "." + ip3 + "." + ip4;
        int portVal = atoi(port1.c_str()) * 256 + atoi(port2.c_str());
        char portStr[10];
        sprintf(portStr, "%d", portVal);
        hostPort.second = string(portStr);
        return hostPort;
}


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
                
                request = "CWD alepho.com/public_html\r\n";
                cout << request;
                client.send(buffer(request, request.size()));
                len = client.receive(buffer(buf, BUFLEN), 0, error);
                cout.write(buf.data(), len);
                cout << endl;
                
                request = "PASV\r\n";
                cout << request;
                client.send(buffer(request, request.size()));
                len = client.receive(buffer(buf, BUFLEN), 0, error);
                cout.write(buf.data(), len);
                cout << endl;
                pair<string, string> portHost = parseHostPort(string(buf.data(), len));
                
                callable call(portHost.first, portHost.second);
                thread th(call);
                
                request = "STOR hello.txt\r\n";
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
