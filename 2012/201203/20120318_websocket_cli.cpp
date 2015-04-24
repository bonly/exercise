#define NOMINMAX
#define _WEBSOCKETPP_CPP11_STL_
// #define _WEBSOCKETPP_CPP11_FUNCTIONAL_
// #define _WEBSOCKETPP_CPP11_MEMORY_


#include <iostream>
#include <string>
 
#include <boost/thread.hpp>
 
#include <websocketpp/config/asio_no_tls_client.hpp>
#include <websocketpp/client.hpp>
 
using namespace std;
 
typedef websocketpp::client<websocketpp::config::asio_client> WSClient;
 
int main( int argc, char** argv )
{
  // string sUrl = "ws://echo.websocket.org";
  // string sUrl = "ws://174.129.224.73"; //此ip是上面网页的,但实现证明不能如此用
  string sUrl = "ws://127.0.0.1";
 
  // cretae endpoint
  WSClient  mEndPoint;
 
  // initial endpoind
  mEndPoint.init_asio();
  // mEndPoint.start_perpetual(); //新版本没有这个方法,所以后面的mEndPoint.run()需要调到connect后才执行
  // boost::thread wsThread( [&mEndPoint](){ mEndPoint.run(); } ); //此句需调后
  
  // set handler
  mEndPoint.set_message_handler(
  []( websocketpp::connection_hdl hdl, WSClient::message_ptr msg ){
      cout << msg->get_payload() << endl;
    }
  );
 
  // get connection and connect
  websocketpp::lib::error_code ec;
  WSClient::connection_ptr wsCon = mEndPoint.get_connection( sUrl, ec);
  if (ec){
    std::cerr << "get_connection err: " << ec.message() << std::endl;
  }
  mEndPoint.connect( wsCon );
 
  boost::thread wsThread( [&mEndPoint](){ mEndPoint.run(); } );//调后到这里connect之后执行

  string sCmd;
  while( true )
  {
    cout << "Enter Command: " << flush;
    getline(cin, sCmd);
 
    if( sCmd == "quit" )
    {
      wsCon->close( 0, "close" );
      break;
    }
 
    auto ec = wsCon->send( sCmd );
    cout << ec.message() << endl;
  }
 
  // mEndPoint.stop_perpetual();  //新版去掉了这个函数
  // mEndPoint.end_perpetual(); //新版有这个函数,但不是这样用
  wsThread.join();
}

/*
g++ -I ~/websocketpp/ -std=c++11 20120318_websocket_cli.cpp -lboost_system -lpthread -lboost_thread
需要:-D_WEBSOCKETPP_CPP11_STL_
*/