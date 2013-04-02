#include <boost/interprocess/managed_mapped_file.hpp>
#include <cassert>
#include <cstring>

class My_class{
  int aInt;
  char cName[12];
};

int main(){
  using namespace boost::interprocess;
  struct shm_remove{
     shm_remove(){ file_mapping::remove("MyShareMemory");}
    //~shm_remove(){ file_mapping::remove("MyShareMemory");}
  }remover;
  managed_mapped_file mfile(create_only, "MyShareMemory", 100*sizeof(My_class));
  
  enum {MAX=5};
  int *pint = mfile.construct<int>("aInt")();
  *pint = 6;

  //pint = mfile.construct<int>("aInt")(); /// 不能创建同名，会运行时异常
  //*pint = 8;

  pint = mfile.find_or_construct<int>("aInt")(); 
  *pint = 9;

  mfile.flush();
  return 0;
}
  
