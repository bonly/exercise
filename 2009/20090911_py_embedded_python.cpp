//Kevin Harris (kevin@codesampler.com)
//    Description: This sample demonstrates how to use the Boost.Python library 
//                 to embed Python into a C++ application. This will allow the 
//                 application to execute Python scripts. The sample also 
//                 demonstrates how to setup special C++ functions which will
//                 allow C++ and Python to communicate back and forth, 
//                 including how to setup a call-back function.
#include <sys/stat.h>
#include <iostream>
using namespace std;

#include <boost/python.hpp>
using namespace boost::python;

//-----------------------------------------------------------------------------
// Create a C++ function will allow Python to get an int value from the
// application.
//-----------------------------------------------------------------------------

int g_somevalue = 55;

PyObject* getvalue(PyObject* self, PyObject* args)
{
  return Py_BuildValue("i", g_somevalue);
}

//-----------------------------------------------------------------------------
// Create a C++ function will allow Python to set or pass four floats to
// the application.
//-----------------------------------------------------------------------------

PyObject* setXYZW(PyObject* self, PyObject* args)
{
  float x, y, z, w;

  if (PyArg_ParseTuple(args, "ffff", &x, &y, &z, &w))
  {
    cout << "Script called setXYZW and passed: " << " x = " << x << " y = "
          << y << " z = " << z << " w = " << w << endl;
  }

  Py_INCREF( Py_None );
  return Py_None;
}

//-----------------------------------------------------------------------------
// Define a function which will allow a Python script to set a call-back for
// the application to use.
//-----------------------------------------------------------------------------

static PyObject *g_pythonCallback = NULL;

static PyObject* setCallback(PyObject* self, PyObject* args)
{
  PyObject *result = NULL;
  PyObject *temp = NULL;

  if (PyArg_ParseTuple(args, "O", &temp))
  {
    if (!PyCallable_Check(temp))
    {
      PyErr_SetString(PyExc_TypeError, "parameter must be callable");
      Py_INCREF( Py_None );
      return Py_None;
    }

    Py_XINCREF( temp ); // Ref the new call-back
    Py_XDECREF( g_pythonCallback ); // Unref the previous call-back
    g_pythonCallback = temp; // Cache the new call-back
  }

  Py_INCREF( Py_None );
  return Py_None;
}

//-----------------------------------------------------------------------------
// Make Python aware of our special C++ functions above.
//-----------------------------------------------------------------------------

PyMethodDef moduleFuncs[] = { { "getvalue", getvalue, METH_VARARGS,
  "Returns an int value" }, { "setXYZW", setXYZW, METH_VARARGS,
  "Sets a 4 component vector" }, { "setCallback", setCallback, METH_VARARGS,
  "Sets a call-back function in Python" }, { NULL, NULL } };

//-----------------------------------------------------------------------------
// Name: readPythonScript()
// Desc:
//-----------------------------------------------------------------------------
char *readPythonScript(const char* fileName)
{
  FILE *pFile = fopen(fileName, "r");

  if (pFile == NULL)
  {
    cout << "Cannot open Python script file, \"" << fileName << "\"!" << endl;
    return 0;
  }

  struct stat fileStats;

  if (stat(fileName, &fileStats) != 0)
  {
    cout << "Cannot get file stats for Python script file, \"" << fileName
          << "\"!" << endl;
    return 0;
  }

  char *buffer = new char[fileStats.st_size];

  int bytes = fread(buffer, 1, fileStats.st_size, pFile);

  buffer[bytes] = 0; // Terminate the string

  fclose(pFile);

  return buffer;
}

//-----------------------------------------------------------------------------
// Name: main()
// Desc: Application's main entry point.
//-----------------------------------------------------------------------------
int main(void)
{
  //
  // Setup Python to be embedded...
  //

  Py_Initialize();

  // Define our custom module called "embedded_test"...
  handle<> embedded_test_module(borrowed(PyImport_AddModule("embedded_test")));
  handle<> embedded_test_init(borrowed(
        Py_InitModule( "embedded_test", moduleFuncs )));

  // Access Python's "__main__" module...
  handle<PyObject> main_module(borrowed(PyImport_AddModule("__main__")));
  // Create a dictionary object for the "__main__" module's namespace.
  dict main_namespace(handle<> (borrowed(PyModule_GetDict(main_module.get()))));

  char* pythonScript = readPythonScript("embedded.py");

  handle<> (PyRun_String( pythonScript,
        Py_file_input,
        main_namespace.ptr(),
        main_namespace.ptr() ));

  // Once the script has finished executing, extract one of it variables.
  int nRetvalue = extract<int> (main_namespace["retvalue"]);

  cout << "Script's retvalue variable was set to: " << nRetvalue << endl;

  delete pythonScript;

  //
  // If the script set the call-back function, call it...
  //

  if (g_pythonCallback)
  {
    int arg = 123;
    PyObject *arglist = NULL;
    PyObject *result = NULL;

    arglist = Py_BuildValue("(i)", arg);
    result = PyEval_CallObject( g_pythonCallback, arglist );
    Py_DECREF( arglist );

    if (result != NULL)
      Py_DECREF( result );
  }

  //
  // Cleanup after Python...
  //

  Py_Finalize();
  return 0;
}
