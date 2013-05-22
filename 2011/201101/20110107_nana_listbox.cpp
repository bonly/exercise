#include <nana/gui/wvl.hpp>
#include <nana/gui/widgets/listbox.hpp>

int main(){
    using namespace nana::gui;
    form fm;
    listbox lb(fm, nana::rectangle(10, 10, 280, 120));
    lb.append_header(STR("Header"), 200);
    lb.append_item(STR("int"));
    lb.append_item(STR("double"));

    lb.append_header(STR("Value"), 200);
    lb.append_item(STR("13"));
    lb.append_item(STR("3.14"));
    
    lb.anyobj(0, 0, 10);
    lb.anyobj(0, 1, 0.1);

    int * pi = lb.anyobj<int>(0, 0);        //it returns a nullptr if there is not an int object is specified.
    double * pd = lb.anyobj<double>(0, 1);  //it returns a nullptr if there is not an double object is specified.

    fm.show();
    exec();
}