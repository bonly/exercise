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
    //shm_remove(){ file_mapping::remove("MyShareMemory");}
    //~shm_remove(){ file_mapping::remove("MyShareMemory");}
  }remover;
  managed_mapped_file mfile(open_read_only, "MyShareMemory");
  
  enum{MAX=5};
  //managed_mapped_file::handle_t handle = mfile.get_address();
  void *ad = mfile.get_address();
  //for (int i=0; i<MAX; ++i){
  	int *aint = (int*)ad+(0*sizeof(int));
  	printf("%d\n", *aint);
  //}
  return 0;
}
