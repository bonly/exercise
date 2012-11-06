#ifndef _PYSIDE_HPP_
#define _PYSIDE_HPP_
#include <boost/python.hpp>

// Macro to declare void intiXXX() function
#define PY_MODULE_DECLARE(MODULENAME) extern "C" void init##MODULENAME();
#define PY_MODULE_IMPLEMENT(x) BOOST_PYTHON_MODULE(x)
#define PY_DEF boost::python::def

void hello();

PY_MODULE_DECLARE(emb)
#endif
