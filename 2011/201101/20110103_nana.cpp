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
bool not_end = true;

class GUI{
public:
    void arun(){

        rx = false;
        frm = new form; //form的生成show 及exec需在一个线程中
        button btn(*frm, nana::rectangle(5, 5, 94, 23));
        btn.caption(STR("开始"));

        button oth(*frm, nana::rectangle(120, 5, 94, 23));
        button *bn2=nullptr;
        oth.make_event<events::click>([&](){
            bn2 = new button(*frm, nana::rectangle(250, 5, 94, 23));
        });
        oth.caption(STR("新建"));

        bx = new textbox(*frm, nana::rectangle(5, 40, 100, 100));

        btn.make_event<events::click>([&](){
              rx = !rx;
              }
            );

        frm->make_event<events::destroy>([&](){ not_end = false;});
        frm->show();

        exec(); //需在这里
        //delete bn2;bn2=nullptr;
        delete frm;frm=nullptr;
        delete bx;bx=nullptr;
    }
    void run(){  //单独函数中也不行
        frm->show();
        exec();
    }
    ~GUI(){
        if(add) delete add;
        if(frm) delete frm;
    }
public:
    form *frm=nullptr;
    textbox *bx=nullptr;
    button *add=nullptr;
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
    //gui->arun();

    //GUI gui;
    boost::thread th(boost::bind(&GUI::arun, boost::ref(*gui)));
    boost::thread ot([&](){
        while(not_end){
            if(gui->rx){
              gui->bx->append(nana::string(STR("atext\n")),false);
              if(gui->add == nullptr)
                  gui->add = new button(*(gui->frm), nana::rectangle(360,5,94,23));
            }
        }
    }
    );

    //boost::thread ex(boost::bind(&exec)); ///单线程跑也不行
    th.join();
    ot.join();


    return 0;
}



