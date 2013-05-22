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
boost::barrier Bar(2);

class GUI{
public:
    GUI(){
        Frm.caption(L"main");
        Frm.show();
    }

    void run(){
        exec();
    }

public:
    form Frm;
};

struct Data{
    int aInt;
};

class GuiSync{
public:
    int InitFrm(form *frm){
        this->frm = frm;
        btn = new button(*frm, nana::rectangle(5,5,120,25));
        btn->caption (STR("开始"));
        return 0;
    }
    ~GuiSync(){
        delete btn;
    }

public:
    form*  frm;
    button* btn;
};

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

    GUI *gui = nullptr;

    boost::thread gui_run([&](){
        gui = new GUI;
        Bar.wait();
        gui->run();
    });

    GuiSync sync;
    boost::thread gui_sync([&](){
        Bar.wait();
        sync.InitFrm(&gui->Frm);
    });

    gui_sync.join();
    gui_run.join();

    delete gui;

    return 0;
}



