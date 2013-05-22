
#include <iostream>
#include <fstream>
#include <string>
using namespace std;

#include <cstdio>
#include <Python.h>
#include <boost/python.hpp>
using namespace boost::python;

//------------------------------------------------------------------------------
// Extend Python by creating a C/C++ function which will allow Python to GET an
// int value from the application.
//------------------------------------------------------------------------------

int g_someValue = 0;

PyObject* getApplicationValue(PyObject* self, PyObject* args)
{
  return Py_BuildValue("i", g_someValue);
}

//------------------------------------------------------------------------------
// Extend Python by creating a C/C++ function which will allow Python to SET an
// int value in the application.
//------------------------------------------------------------------------------

PyObject* setApplicationValue(PyObject* self, PyObject* args)
{
  int nValue;

  if (PyArg_ParseTuple(args, "i", &nValue))
  {
    g_someValue = nValue;
    cout << "Script called \"setApplicationValue\" and passed: " << nValue
          << endl;
  }

  Py_INCREF( Py_None );
  return Py_None;
}

//------------------------------------------------------------------------------
// Extend & Embed Python by defining a function which will allow a Python script
// to set a call-back for the application to use. This will allow C++ to call
// into a Python script to have work done.
//------------------------------------------------------------------------------

PyObject *g_pythonCallback = NULL;

static PyObject* setCallback(PyObject* self, PyObject* args)
{
  PyObject *pResult = NULL;
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

//------------------------------------------------------------------------------
// Make Python aware of our special C++ functions above.
//------------------------------------------------------------------------------
PyMethodDef g_methodDefinitions[] = { { "getApplicationValue",
  getApplicationValue, METH_VARARGS, "Returns an int value" },
  { "setApplicationValue", setApplicationValue, METH_VARARGS,
    "Sets an int value" }, { "setCallback", setCallback, METH_VARARGS,
    "Sets a call-back function in Python" }, { NULL, NULL } };
/*
BOOST_PYTHON_MODULE(extendAndEmbedTest)
{
  def("setCallback",setCallback);
  def("getApplicationValue", getApplicationValue);
}
*/
//-----------------------------------------------------------------------------
// Name: readPythonScript()
// Desc:
//-----------------------------------------------------------------------------
string *readPythonScript(string fileName)
{
  ifstream pythonFile;

  pythonFile.open(fileName.c_str());

  if (!pythonFile.is_open())
  {
    cout << "Cannot open Python script file, \"" << fileName << "\"!" << endl;
    return NULL;
  }
  else
  {
    // Get the length of the file
    pythonFile.seekg(0, ios::end);
    int nLength = pythonFile.tellg();
    pythonFile.seekg(0, ios::beg);

    // Allocate  a char buffer for the read.
    char *buffer = new char[nLength];
    memset(buffer, 0, nLength);

    // read data as a block:
    pythonFile.read(buffer, nLength);

    string *scriptString = new string;
    scriptString->assign(buffer);

    delete[] buffer;
    pythonFile.close();

    return scriptString;
  }
}

int main(void)
{
  Py_Initialize();

  object mod(handle<>(borrowed(PyImport_AddModule("extendAndEmbedTest"))));
  //PyObject *ret = PyImport_AddModule("pyobj");
  //PyObject *ret = Py_InitModule( "pyobj", g_methodDefinitions ); //不能通过BOOST_PYTHON_MODULE()代替
  object moth(handle<>(Py_InitModule( "extendAndEmbedTest", g_methodDefinitions ))); //不能通过BOOST_PYTHON_MODULE()代替

  //
  // Access the "__main__" module and its name-space dictionary.
  //

  object  main_mod(handle<>(borrowed(PyImport_AddModule("__main__"))));
  object  dict = main_mod.attr("__dict__");
  //PyObject *pMainModule = PyImport_AddModule("__main__");
  //PyObject *pMainDictionary = PyModule_GetDict(pMainModule);

  //
  // Exercise embedding by calling a Python script from a C++ application.
  //

  string *pythonScript = readPythonScript("test.py");

  if (pythonScript != NULL)
  {
    //ret = PyRun_String( pythonScript->c_str(), Py_file_input,
    //      pMainDictionary, pMainDictionary );
    //if (ret==0)
    //  PyErr_Print();

    delete pythonScript;
  }

  try{
    exec_file("test.py",dict);
  }
  catch(...)
  {
   PyErr_Print();
  }

  //FILE *fp=fopen("test.py","r");
  //int r = PyRun_SimpleFile(fp,"test.py");
  //if (r==0)
  //  PyErr_Print();

  //
  // Once the script has finished executing, extract one of it variables from
  // the name-space dictionary.
  //

  //PyObject *pResult = PyDict_GetItemString(pMainDictionary, "returnValue");
  //int nValue;
  //PyArg_Parse(pResult, "i", &nValue);

  int nValue = extract<int>(main_mod.attr("returnValue"));
  cout << "Script's \"returnValue\" variable was set to: " << nValue << endl;
  //
  // If the Python script set the call-back function, call it...
  //

  if (g_pythonCallback)
  {
    int nArg1 = 123;

    //PyObject *pArgList = Py_BuildValue("(i)", nArg1);
    //PyObject *pResult = PyEval_CallObject( g_pythonCallback, pArgList );

    try
    {
      call<void>(g_pythonCallback,nArg1);
    }
    catch(...)
    {
      PyErr_Print();
    }

    //Py_DECREF( pArgList );

    //if (pResult != NULL)
    //  Py_DECREF( pResult );
  }

  //
  // Cleanup after Python...
  //

  Py_Finalize();
  return 0;
}

/*
#------------------------------------------------------------------------------
#           Name: test.py
#         Author: Kevin Harris
#  Last Modified: 04/29/05
#    Description:
#------------------------------------------------------------------------------

import extendAndEmbedTest

#------------------------------------------------------------------------------
# Define a function and set it as a call-back for the application to use
#------------------------------------------------------------------------------

def somePythonCallBackFunction( n ):
    print "Application called \"somePythonCallBackFunction\" and passed: " + str( n )

extendAndEmbedTest.setCallback( somePythonCallBackFunction );

#------------------------------------------------------------------------------
# Test the getValue() function which gets an int value form the application.
#------------------------------------------------------------------------------

print "extendAndEmbedTest.getApplicationValue() returned: " + str( extendAndEmbedTest.getApplicationValue() )

#------------------------------------------------------------------------------
# Test the setApplicationValue() function which sets an int value in the
# application.
#------------------------------------------------------------------------------

extendAndEmbedTest.setApplicationValue( 55 )

# Let's check it again...
print "extendAndEmbedTest.getApplicationValue() returned: " + str( extendAndEmbedTest.getApplicationValue() )

#------------------------------------------------------------------------------
# Set a return value so we can extract it after the script has executed.
#------------------------------------------------------------------------------

returnValue = -1

 */
