#include <iostream>
#include <memory>

class CA {
public:
   CA(){ 
     std::clog << "CA()" << std::endl;
   }
   virtual ~CA(){
     std::clog << "~CA()" << std::endl;
   }
};

class CB : public CA {
public: 
   CB(){
      std::clog << "CB()" << std::endl;
   }
   ~CB(){
      std::clog << "~CB()" << std::endl;
   }
};

int main()
{
   std::shared_ptr<CA> cb =std::make_shared<CB>();
   return 0;
}

