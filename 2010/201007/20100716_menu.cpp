/**
 * @file 20100716_menu.cpp
 * @brief 
 * @author bonly
 * @date 2012-11-6 bonly Created
 */

#include <boost/python.hpp>
#include <boost/thread.hpp>
#include <iostream>

using namespace boost::python;

class PyInf
{
    public:
        PyInf()
        {
            Py_Initialize();
            //PyEval_InitThreads(); ///多线程支持
            main_mod = import("__main__"); ///如果在python语句中使用了变量,就须要指定场景参数,则导入这个
            main_spa = main_mod.attr("__dict__");
        }

        void init()
        {
            //PyEval_AcquireLock();
            //Py_BEGIN_ALLOW_THREADS;
            exec_file ("20100715_menu.py",main_spa);
            exec("app = wx.PySimpleApp()\n"
                 "frame = ToolbarFrame(parent=None, id=-1)\n"
                 "frame.Show()\n", main_spa);
            //PyEval_ReleaseLock();
            //Py_END_ALLOW_THREADS;

        }
        void ClickButton()
        {
            //PyEval_AcquireLock();
            //Py_BEGIN_ALLOW_THREADS;

            std::clog << "u click the button\n";
            exec("import wx", main_spa);
            //exec("print 10", main_spa);
            /*
            exec("dlg=wx.MessageDialog(None, 'c直接建对话框','MessageDialog', wx.YES_NO|wx.ICON_QUESTION);\n"
                 "result=dlg.ShowModal();\n"
                 "dlg.Destroy();"
                 ,main_spa);
            //*/
            //Py_END_ALLOW_THREADS;
            //PyEval_AcquireLock();
        }

        void run()
        {
            //PyEval_AcquireLock();
            //Py_BEGIN_ALLOW_THREADS;
            exec("app.MainLoop()\n", main_spa);
            //PyEval_ReleaseLock();
            //Py_END_ALLOW_THREADS;
        }

        ~PyInf()
        {
            //PyEval_ReleaseLock(); ///多线程释放
            //Py_END_ALLOW_THREADS;
            Py_Finalize();
        }

    private:
        object main_mod;
        object main_spa;
};

BOOST_PYTHON_MODULE(inf)
{
    class_<PyInf>("PyInf").def("ClickButton", &PyInf::ClickButton);
}

void other_call(PyInf &py)
{
    boost::this_thread::sleep(boost::posix_time::milliseconds(5000));
    py.ClickButton();
}
int main(int argc, char *argv[])
{
    try
    {
        PyInf myc;
        myc.init();
        boost::thread thr(&PyInf::run, myc);

        /*
        PyInf otc;
        otc.init();
        boost::thread cl(&other_call, otc);
        */
        //boost::thread cl(&other_call, myc);


        thr.join();
        //cl.join();
    }
    catch (...)
    {
        PyErr_Print();
    }

    return 0;
}

