#include <boost/interprocess/managed_shared_memory.hpp>

int main(){
   using namespace boost::interprocess;

   struct shm_remove{
      shm_remove() { shared_memory_object::remove("MySharedMemory");}
      ~shm_remove(){ shared_memory_object::remove("MySharedMemory");}
   } remover;

   managed_shared_memory managed_shm(create_only,"MySharedMemory", 65536);
   void *ptr = managed_shm.allocate(100); //100byte
   managed_shm.deallocate(ptr);
   ptr = managed_shm.allocate(100, std::nothrow);

   managed_shm.deallocate(ptr);
   return 0;
}

