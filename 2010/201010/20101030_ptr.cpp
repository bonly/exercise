#include <memory>
#include <thread>

struct AC{
   int  myint;
};

std::shared_ptr<AC> getAC(){
   return std::make_shared<AC>();
}

void run(){
   for (int i=0; i<20; ++i){
      std::shared_ptr<AC> ac = getAC();
      ac->myint = 20;
   }
}

int main(){
  std::thread tg(std::bind(&run));

  tg.join();
  return 0;
}

/**
  C++11中并没有加入thread_group
  只能自己制作容器了
*/
