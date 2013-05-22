#include <gtkmm.h>
#include <iostream>
 
Gtk::Window* pWindow = 0;
 
int main (int argc, char **argv)
{
    Gtk::Main kit(argc, argv);
    //Load the GtkBuilder file and instantiate its widgets:
    Glib::RefPtr refBuilder = Gtk::Builder::create();
    try
    {         //导入Glade生成的xml文件
        refBuilder->add_from_file("mygui.glade");
    }
    catch(const Glib::FileError & ex)
    {
        std::cerr << "FileError: " << ex.what() << std::endl;
        return 1;
    }
    catch(const Gtk::BuilderError & ex)
    {
        std::cerr << "BuilderError: " << ex.what() << std::endl;
        return 1;
    }
    //Get the GtkBuilder-instantiated Window:
    refBuilder->get_widget("window", pWindow);
    if(pWindow)
    {
        kit.run(*pWindow);
    }
    delete pWindow;
    return 0;
}

#pkg-config --cflags --libs gtkmm-2.4
