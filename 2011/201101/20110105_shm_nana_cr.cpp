/**
 * @file main.cpp
 * @brief 
 * @author bonly
 * @date 2013-4-1 bonly Created
 */

#include <nana/gui/wvl.hpp>
#include <nana/gui/widgets/button.hpp>
#include <nana/gui/widgets/textbox.hpp>
#include <nana/gui/widgets/label.hpp>
#include <boost/thread.hpp>
#include <boost/bind.hpp>
#include <boost/interprocess/managed_mapped_file.hpp>

using namespace nana::gui;
bool not_end = true;

class GUI{
public:
    void run(){
        frm = new form();
        label lb(*frm, nana::rectangle(5, 5, 100, 20));
        lb.caption(STR("abc"));
        frm->show();
        exec();
    }
    ~GUI(){
        if(frm) delete frm;
    }
public:
    form *frm=nullptr;
};

struct Data{
    int aInt;
};

//typedef boost::interprocess::basic_managed_mapped_file < \
//   char, \
//   boost::interprocess::rbtree_best_fit<boost::interprocess::mutex_family>, \
//   boost::interprocess::map_index \
//   >  my_mapped_file;

//typedef boost::interprocess::basic_managed_mapped_file
//< char,  boost::interprocess::rbtree_seq_fit< boost::interprocess::null_mutex_family>,  boost::interprocess::map_index >
//my_mapped_file;

int main(){
    using namespace boost::interprocess;
    struct shm_remove{
       shm_remove(){ file_mapping::remove("MyShareMemory");}
      ~shm_remove(){ file_mapping::remove("MyShareMemory");}
    }remover;

    managed_mapped_file mfile(create_only, "MyShareMemory", 500*sizeof(Data));
    Data *dt = mfile.construct<Data>("aData")();
    dt->aInt = 10;
    mfile.flush();

    GUI *gui = new GUI;

    boost::thread gui_run(boost::bind(&GUI::run, boost::ref(*gui)));
    gui_run.join();

    delete gui;

    return 0;
}



