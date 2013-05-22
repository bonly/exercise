#include <nana/gui/wvl.hpp>
#include <nana/gui/widgets/listbox.hpp>

struct person{
    nana::string name;
    bool gender;
    unsigned age;
};

class resolver : public nana::gui::listbox::resolver_interface<person>{
    nana::string decode(std::size_t i, const person& pr) const {
        std::stringstream ss;
        switch(i) {
        case 0:
            return pr.name;
        case 1:
            return (pr.gender ? STR("Male") : STR("Female"));
        case 2:
            ss<<pr.age;
            return nana::charset(ss.str());
        }
        return nana::string();
    }

    void encode(person& pr, std::size_t i, const nana::string& s) {
        std::stringstream ss;
        switch(i) {
        case 0:
            pr.name = s;
            break;
        case 1:
            pr.gender = (s == STR("Male"));
            break;
        case 2:
            ss<<static_cast<std::string>(nana::charset(s));
            ss>>pr.age;
            break;
        }
    }
}

int main(){
    using namespace nana::gui;

    form fm;
    listbox  lb(fm, nana::rectangle(10, 10, 280, 120));
    lb.append_header(STR("Name"), 100);
    lb.append_header(STR("Gender"), 60);
    lb.append_header(STR("Age"), 60);

    lb.resolver(resolver()); //Pass the user defined resolver to the listbox

    person pr;
    pr.name = STR("Peter");
    pr.gender = true;
    pr.age = 10;
    lb.append(0, pr);  //person to item. resolver::decode() is required
    
    lb.item(0, 0, pr); //item to person. resolver::encode() is required

    fm.show();
    exec();
}
