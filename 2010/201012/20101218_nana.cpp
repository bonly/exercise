#include <nana/gui/wvl.hpp>
#include <nana/gui/widgets/label.hpp>
#include <nana/deploy.hpp>
int main()
{
    using namespace nana::gui;
    form fm;
    label lb(fm, fm.size());
    lb.caption(STR("Hello, World"));
    fm.show();
    exec();
}