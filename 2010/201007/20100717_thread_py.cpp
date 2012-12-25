//
#include <Python.h>
#include <boost/thread.hpp>
#include <boost/date_time/posix_time/posix_time.hpp>

#include <stdio.h>

//#include "gil_lock.h"
//#define OUTSIDE

void Callback()
{
    PyObject *pModule, *pDict, *pFunc, *pValue;

#ifndef OUTSIDE
    //gil_lock glock;
#endif
    for(int i = 0; i < 20; i++)
    {
#ifdef OUTSIDE
        gil_lock glock;
#endif

        pModule = PyImport_ImportModule("test1");

        pDict = PyModule_GetDict(pModule);
        pFunc = PyDict_GetItemString(pDict, "test");

        if (PyCallable_Check(pFunc))
    {
      pValue = PyObject_CallObject(pFunc, NULL);
    }
    else
    {
      PyErr_Print();
    }

        //printf("Cleanup #%d\n", ldata);
    // Clean up

    Py_DECREF(pModule);
    //boost::this_thread::sleep(boost::posix_time::milliseconds(350));
    }
}


void start_python()
{
    PyThreadState *py_tstate = NULL;
    //...
    Py_Initialize();
    PyEval_InitThreads();
    //...
    py_tstate = PyGILState_GetThisThreadState();
    PyEval_ReleaseThread(py_tstate);
    //gil_lock glock;
    PyRun_SimpleString("import sys");
    PyRun_SimpleString("print sys.path");
}

void end_python()
{
    //...
    PyGILState_Ensure();
    Py_Finalize();
}

int main(int argc, char *argv[])
{
    start_python();

    const size_t numThreads = 30;
    boost::thread_group threads;
    for (size_t n=0; n<numThreads; ++n)
    {
        threads.create_thread(Callback);
    }

    printf("Waiting for threads to end...\n");
    threads.join_all();
    printf("Done\n");

    printf("Press any key to end this...\n");
    //getchar();
    end_python();
    return 0;
}
