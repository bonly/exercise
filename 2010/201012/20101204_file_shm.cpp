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
  //void *array[MAX];
  //for (int i=0; i<MAX; ++i){
    /// get_handle_from_address(ptr),ptr
  	int *array=(int*)mfile.allocate(sizeof(int)); 
  	*array=445;
  //}
  
  mfile.flush();
  return 0;
}
  
