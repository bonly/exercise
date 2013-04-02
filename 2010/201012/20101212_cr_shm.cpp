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
  mfile.flush();
  return 0;
}
  
