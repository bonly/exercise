//------------------------------------------------------------------------------
//           Name: pyc_extend_embed.cpp
//         Author: Kevin Harris
//  Last Modified: 01/19/07
//    Description: This sample demonstrates how to integrate C/C++ and
//                 Python together by both extending and embedding Python
//                 into a C/C++ application. This is done by using the C API of
//                 Python to set up special functions which will allow C/C++
//                 and Python to communicate back and forth.
//
// Note:
//
// To compile this project, you'll need to install Python 2.5. I downloaded and
// used "python-2.5.msi", which I downloaded from here:
//
// http://www.python.org/download/
//
// Note:
//
// With certain versions of Python, the Debug build may fail to link because it
// can't find "python25_d.lib". In that case, you'll need to stick with the
// Release build, or download and build a debug version of the Python library
// from the source, so you can link it in.
//------------------------------------------------------------------------------

#include <iostream>
#include <fstream>
#include <string>
using namespace std;

#include <Python.h>

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

//------------------------------------------------------------------------------
// Name: main()
// Desc: Application's main entry point.
//------------------------------------------------------------------------------
int main(void)
{
  //
  // Setup Python to be embedded...
  //

  Py_Initialize();

  //
  // Define our custom module called "extendAndEmbedTest"...
  //

  PyImport_AddModule("extendAndEmbedTest"); //此模块名只需与脚本中导入的名一致即可,无命名限制
  Py_InitModule( "extendAndEmbedTest", g_methodDefinitions );

  //
  // Access the "__main__" module and its name-space dictionary.
  //

  PyObject *pMainModule = PyImport_AddModule("__main__");
  PyObject *pMainDictionary = PyModule_GetDict(pMainModule);

  //
  // Exercise embedding by calling a Python script from a C++ application.
  //

  string *pythonScript = readPythonScript("test.py");

  if (pythonScript != NULL)
  {
    PyObject *prun = PyRun_String( pythonScript->c_str(), Py_file_input,
          pMainDictionary, pMainDictionary );

    delete pythonScript;
  }

  //
  // Once the script has finished executing, extract one of it variables from
  // the name-space dictionary.
  //

  PyObject *pResult = PyDict_GetItemString(pMainDictionary, "returnValue");

  int nValue;
  PyArg_Parse(pResult, "i", &nValue);
  cout << "Script's \"returnValue\" variable was set to: " << nValue << endl;

  //
  // If the Python script set the call-back function, call it...
  //

  if (g_pythonCallback)
  {
    int nArg1 = 123;

    PyObject *pArgList = Py_BuildValue("(i)", nArg1);
    PyObject *pResult = PyEval_CallObject( g_pythonCallback, pArgList );

    Py_DECREF( pArgList );

    if (pResult != NULL)
      Py_DECREF( pResult );
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
#此模块名只需与Ｃ中建的模块名一致即可
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
