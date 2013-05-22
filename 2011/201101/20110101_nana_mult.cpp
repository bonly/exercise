/**
 * @file main.cpp
 * @brief 
 * @author bonly
 * @date 2013-4-1 bonly Created
 */

#include <nana/gui/wvl.hpp>
#include <nana/gui/widgets/button.hpp>
#include <nana/gui/widgets/textbox.hpp>
#include <boost/thread.hpp>
#include <boost/bind.hpp>
#include <boost/interprocess/managed_mapped_file.hpp>

using namespace nana::gui;

class GUI_O{
public:
    void run(){
        frm = new form;
        button btn(*frm, nana::rectangle(5, 5, 94, 23));
        btn.caption(STR("开始"));

        bx = new textbox(*frm, nana::rectangle(5, 40, 100, 100));

        btn.make_event<events::click>([&](){
              bx->append(nana::string(STR("atext")),false);
              }
            );
        frm->show();
    }
public:
    form *frm;
    textbox *bx;
    bool rx;
};

class GUI{
public:
    void run(){
        rx = false;
        form frm;
        Frm = &frm;
        button btn(frm, nana::rectangle(5, 5, 94, 23));
        btn.caption(STR("开始"));

        button oth(frm, nana::rectangle(120, 5, 94, 23));
        oth.make_event<events::click>([](){
            GUI_O o;
            o.run();
        });

        textbox bx(frm, nana::rectangle(5, 40, 100, 100));
        Bx = &bx;

        btn.make_event<events::click>([&](){
              rx = !rx;
              }
            );
        frm.show();
    }
public:
    form *Frm;
    textbox *Bx;
    bool rx;
};



int main(){
    using namespace boost::interprocess;
    struct shm_remove{
       shm_remove(){ file_mapping::remove("MyShareMemory");}
      ~shm_remove(){ file_mapping::remove("MyShareMemory");}
    }remover;

    managed_mapped_file mfile(create_only, "MyShareMemory", 500*sizeof(GUI));
    GUI *gui = mfile.construct<GUI>("agui")();

    //GUI gui;
    boost::thread th(boost::bind(&GUI::run, boost::ref(*gui)));
    boost::thread ot([&](){
        while(true){
            if(gui->rx)
              gui->Bx->append(nana::string(STR("atext\n")),false);
        }
    }
    );
    
    boost::thread ex(boost::bind(&exec));
    
    th.join();
    ot.join();
    ex.join();


    return 0;
}



