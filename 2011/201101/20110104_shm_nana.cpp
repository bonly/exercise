#include <nana/gui/wvl.hpp>
#include <nana/gui/widgets/button.hpp>
#include <nana/gui/widgets/textbox.hpp>
#include <boost/thread.hpp>
#include <boost/bind.hpp>
#include <boost/interprocess/managed_mapped_file.hpp>
#include <iostream>

using namespace nana::gui;
struct Data{
    int aInt;
};

int main(){
   using namespace boost::interprocess;
   using namespace std;
   try{
      managed_mapped_file mfile(open_only, "MyShareMemory");

	   auto ret = mfile.find<GUI>("aData");
	   if (ret.second != 1){
	     clog << "not found " << endl;
	     return 1;
	   }
	   Data* aData = ret.first;
	   aData->aInt = 30;
	   mfile.flush();
   }catch(exception &e){
   	  cerr << "open shred memory failed! " << e.what() << endl;
   }

   return 0;
}

   
