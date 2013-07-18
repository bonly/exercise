#include <boost/asio.hpp>
#include <iostream>
bool bexit = false;
void signal_handler(const boost::system::error_code & err, int signal) {
    std::clog << "recv signal" << std::endl;
    bexit = true;
    // log this, and terminate application
}

int main(){
  boost::asio::io_service service;
  boost::asio::signal_set sig(service, SIGINT, SIGTERM);
  while(!bexit){
      sig.async_wait(signal_handler);
      std::clog << "waiting ... " << std::endl;
      sleep(1);
  }
  
}

