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
                                                                                                                                                                                 en(choice, length):

   passwd = ""

   alphanum = ('0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ')
   alpha = ('abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ')
   alphacap = ('ABCDEFGHIJKLMNOPQRSTUVWXYZ')
   all = ('abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&amp;*()-_+=~[]{}|\:;"\'&lt;&gt;,.?/')

   if str(choice).lower() == "alphanum":
      choice = alphanum

   elif str(choice).lower() == "alpha":
      choice = alpha

   elif str(choice).lower() == "alphacap":
      choice = alphacap

   elif str(choice).lower() == "all":
      choice = all

   else:
      print "Type doesn't match\n"
      sys.exit(1)

   return passwd.join(random.sample(choice, int(length)))

title()
if len(sys.argv) &lt;= 3 or len(sys.argv) == 5:
   print "\nUsage: ./passgen.py   "
   print "\t[options]"
   print "\t   -w/-write  : Writes passwords to file\n"
   print "There are 4 types to use: alphanum, alpha, alphacap, all\n"
   sys.exit(1)

for arg in sys.argv[1:]:
   if arg.lower() == "-w" or arg.lower() == "-write":
      txt = sys.argv[int(sys.argv[1:].index(arg))+2]

if sys.argv[3].isdigit() == False:
   print sys.argv[3],"must be a number\n"
   sys.exit(1)
if sys.argv[2].isdigit() == False:
   print sys.argv[2],"must be a number\n"
   sys.exit(