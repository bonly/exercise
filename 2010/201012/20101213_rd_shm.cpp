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
  
  void *ad = mfile.get_address();
  
  std::pair<int*, managed_mapped_file::size_type> res;
  res = mfile.find<int> ("aInt");
  
  printf("%d\n", *(res.first));

  return 0;
}
