#include <iostream>
#include <boost/signals2.hpp>

struct HelloWorld{
   void operator()() const {
     std::cout << "Hello, world!" << std::endl;
   }
};


int main(){
  boost::signals2::signal<void ()> sig;
  HelloWorld hello;
  sig.connect (hello);
  sig();
}

