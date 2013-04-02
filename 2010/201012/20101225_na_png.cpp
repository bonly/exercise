/**
在0.2版的发布中，Nana提供了对PNG的支持，但是在默认状态下，为了方便快捷地配置程序库，Nana关闭了支持PNG的特性。

对PNG的支持，Nana C++ Library是使用的libpng(www.libpng.org)。因此，有两个不同的策略来支持这一特性。
1，使用Nana库附带的libpng
2，使用操作系统中提供的libpng，这意味着我们需要自行地安装libpng。
打开位于Nana包含目录中的config.hpp头文件。我们可以找到一行类似这样的注释：
//#define NANA_ENABLE_PNG  
通过撤消这个注释来开启对PNG的支持
保留#define NANA_LIBPNG，则使用Nana库附带的libpng。注释这句就是用操作系统中提供的libpng。
默认情况下，Win32包的Nana库是使用附带的libpng，而Linux(X11)包则使用操作系统提供的。

配置完成之后，重新编译Nana库
在Windows上，链接位于“%nana%/extrlib” 文件夹中的libpng。
在Linux(X11)上，修改你的makefile，为编译器添加一条"-lpng"
*/
#include <nana/gui/wvl.hpp>
#include <nana/gui/widgets/picture.hpp>
#include <nana/gui/layout.hpp>

int main()
{
    using namespace nana::gui;
    form fm;
    picture pic(fm);

    gird    gd(fm);
    gd.push(pic, 0, 0);

    pic.load(STR("a_png_file.png"));
    fm.show();
    exec();
}
