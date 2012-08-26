// Compile this with:
// g++ -Wall -o boosttest -lboost_python -lpython2.3 -lreadline -I /usr/include/python2.3/ boosttest.cc

/** Sample test script:

 class test(boosttest.ABC):
 def say(self):
 return "hi"
 def say_else(self, fact):
 return "hi " + str(fact.get_count())

 t = test()
 c.add_abc(a)
 c.add_abc(t)
 c.abc_says()
 c.abc_says_else()

 */

#include <Python.h>
#include <boost/python.hpp>
#include <boost/shared_ptr.hpp>
using namespace boost;
using namespace boost::python;

#include <iostream>
#include <sstream>
#include <string>

using namespace std;

#include <unistd.h>
#include <stdio.h>

#undef HAVE_LIBREADLINE
#ifdef HAVE_LIBREADLINE
#include <readline/readline.h>
#include <readline/history.h>
#endif

#include <boost/python/numeric.hpp>
#include <boost/python/tuple.hpp>

#include <string>
#include <vector>

// This ties us to Numeric...
#ifdef HAVE_NUMERIC
#include <Numeric/arrayobject.h>
#endif

#ifdef HAVE_LIBREADLINE
char *rl(const char *prompt)
{
  char *line_read = readline(prompt);

  /* If the line has any text in it,
   save it on the history. */
  if (line_read && *line_read)
    add_history(line_read);

  return (line_read);
}
#else
char *rl(const char *prompt)
{
  return 0;
}
#endif

class Fiction
{
  public:
    unsigned int count_;
    Fiction() :
      count_(0)
    {
      cout << "In Fiction::Fiction()" << endl;
    }

    ~Fiction()
    {
      cout << "In Fiction::~Fiction()" << endl;
    }

    Fiction(const Fiction &rhs)
    {
      cout << "In Fiction::Fiction(const Fiction&)" << endl;
      count_ = rhs.count_;
    }

    unsigned int get_count()
    {
      return count_++;
    }
};

class Hmm
{
  public:
    Hmm()
    {
    }

    virtual ~Hmm()
    {
    }

    virtual string say()
    {
      return "HMMMMMMMmmmm!";
    }
};

class ABC
{
  public:
    Hmm hmm_;

    ABC()
    {
      cout << "In ABC::ABC()." << endl;
    }

    virtual ~ABC() // DESTRUCTORS CANNOT BE DEFINED = 0! That won't work with python.
    {
      cout << "In ABC::~ABC()." << endl;
    }

    virtual std::string say() = 0;
    virtual int bam(int a, int b)
    {
      return a & b;
    }
    std::string say_what()
    {
      return say();
    }

    virtual const Hmm &get_hmm()
    {
      return hmm_;
    }
    virtual std::string say_hmm()
    {
      return hmm_.say();
    }
    virtual std::string say_else(Fiction& hmm)
    {
      ostringstream out;
      out << "ABC says hmm " << hmm.get_count() << " times. " << endl;
      return out.str();
    }
};

class derived: public ABC
{
  public:
    virtual ~derived()
    {
    }

    derived()
    {
    }

    std::string say()
    {
      return "hello world. I'm a talking monkey!";
    }
};

std::string call_say(ABC& b)
{
  return b.say();
}
int call_bam(ABC& abc, int a, int b)
{
  return abc.bam(a, b);
}
std::string call_say_what(ABC& b)
{
  return b.say_what();
}
const Hmm & call_get_hmm(ABC &b)
{
  return b.get_hmm();
}
std::string call_say_hmm(ABC &b)
{
  return b.say_hmm();
}
std::string call_say_else(ABC &b, Fiction &hmm)
{
  return b.say_else(hmm);
}

class ABCWrap: public ABC
{
  public:
    ABCWrap(PyObject* self_) :
      self(self_)
    {
    }
    std::string say()
    {
      return call_method<std::string> (self, "say");
    }
    std::string say_what()
    {
      return call_method<std::string> (self, "say_what");
    }
    int bam(int a, int b)
    {
      return call_method<int> (self, "bam");
    }
    const Hmm &get_hmm()
    {
      return call_method<const Hmm&> (self, "get_hmm");
    }
    std::string say_else(Fiction &hmm)
    {
      return call_method<std::string> (self, "say_else", boost::ref(hmm));
    }

    std::string default_say_else(Fiction &hmm)
    {
      return ABC::say_else(hmm);
    }

    PyObject* self;
};

class abc_container
{
  private:
    vector<shared_ptr<ABC> > abcs_;
    Fiction fact_;
  public:
    abc_container()
    {
    }

    virtual ~abc_container()
    {
    }

    void add_abc(shared_ptr<ABC> abc)
    {
      abcs_.push_back(abc);
    }

    void abc_says()
    {
      vector<shared_ptr<ABC> >::iterator iter = abcs_.begin();
      vector<shared_ptr<ABC> >::iterator iter_e = abcs_.end();

      int i = 0;
      while (iter != iter_e)
      {
        cout << "Item " << i << " says: '" << (*iter)->say() << "'" << endl;
        ++iter;
        ++i;
      }
    }

    void abc_says_else()
    {
      vector<shared_ptr<ABC> >::iterator iter = abcs_.begin();
      vector<shared_ptr<ABC> >::iterator iter_e = abcs_.end();

      int i = 0;
      while (iter != iter_e)
      {
        cout << "Item " << i << " says else: '" << (*iter)->say_else(fact_)
              << "'" << endl;
        ++iter;
        ++i;
      }
    }
};

// class AbcContWrap : public abc_container
// {
// public:
//   PyObject *self;
//   vector<object> objs_;
//   object slots[4];

//   AbcContWrap(PyObject *self_)
//     : self(self_) {}

//   void slot_abc(unsigned int s, object obj)
//   {
//     if (s < 4)
//     {
//       slots[s] = obj;
//     }
//   }

//   void add_abc(object obj)
//   {
//     objs_.push_back(obj);
//     abc_container::add_abc(extract<ABC&>(obj));
//   }

//   void abc_says()
//   {
//     abc_container::abc_says();
//   }

//   void abc_says_else()
//   {
//     abc_container::abc_says_else();
//   }
// };


BOOST_PYTHON_MODULE(boosttest)
{
  class_<Fiction> ("Fiction", no_init) .def("get_count", &Fiction::get_count);

  class_<Hmm> ("Hmm") .def("say", &Hmm::say);

  class_<ABC, ABCWrap, boost::noncopyable> ("ABC") .def("say", &ABCWrap::say) .def(
        "say_what", &ABCWrap::say_what) .def("say_else", &ABCWrap::say_else) //, &ABCWrap::default_say_else)
  .def("get_hmm", &ABCWrap::get_hmm, return_internal_reference<> ());

  class_<derived, bases<ABC> > ("derived") .def("say", &derived::say) .def(
        "bam", &derived::bam) .def("say_else", &ABC::say_else);

  //  class_<abc_container, AbcContWrap, boost::noncopyable>("abc_container")
  //     .def("add_abc", &AbcContWrap::add_abc)
  //     .def("slot_abc", &AbcContWrap::slot_abc)
  //     .def("abc_says", &AbcContWrap::abc_says)
  //     .def("abc_says_else", &AbcContWrap::abc_says_else)

  class_<abc_container, boost::noncopyable> ("abc_container") .def("add_abc",
        &abc_container::add_abc) .def("abc_says", &abc_container::abc_says) .def(
        "abc_says_else", &abc_container::abc_says_else);
}

struct _inittab modules_[] = { { (char*)"boosttest", &initboosttest }, { 0, 0 } };

int main(int argc, char **argv)
{
  int ret = PyImport_ExtendInittab(modules_);
  if (ret == -1)
  {
    cout << "Unable to extend python with built in modules." << endl;
    return 1;
  }

  Py_Initialize();

#ifdef HAVE_LIBREADLINE
  rl_bind_key('\t', rl_insert);
#endif

  try
  {
    handle<> main_module(borrowed(PyImport_AddModule("__main__")));
    handle<> main_namespace(borrowed(PyModule_GetDict(main_module.get())));

    handle<> btname(PyString_FromString("boosttest"));
    handle<> bt(PyImport_Import(btname.get()));
    PyDict_SetItemString(main_namespace.get(), "boosttest", bt.get());

    handle<> sys(PyImport_ImportModule("sys"));
    PyDict_SetItemString(main_namespace.get(), "sys", sys.get());

    handle<> resasdas(PyRun_String("from Numeric import *", Py_single_input,
          main_namespace.get(),
          main_namespace.get()));

    handle<> res(PyRun_String("dir()", Py_file_input, main_namespace.get(),
          main_namespace.get()));

    handle<> res2(PyRun_String("a = boosttest.derived()\n"
          "a.say()", Py_file_input,
          main_namespace.get(),
          main_namespace.get()));

    handle<> res3(PyRun_String("c = boosttest.abc_container()\n",
          Py_file_input,
          main_namespace.get(),
          main_namespace.get()));

    //     numeric::array arr(make_tuple(1, 2, 3));
    //     PyDict_SetItemString(main_namespace.get(), "garr", arr.ptr());
    //     handle<> bonk(PyRun_String("zeros((3,3))", Py_single_input,
    //                              main_namespace.get(),
    //                              main_namespace.get()));

    //     handle<> zeros (PyRun_String("zeros", Py_single_input,
    //                                  main_namespace.get(),
    //                                  main_namespace.get()));

    //numeric::array bk = extract<numeric::array>(bonk.get() );

    //     numeric::array obj(make_tuple(make_tuple(1, 2), make_tuple(3, 4)));
    //     //obj.resize(make_tuple(2,2));
    //     PyDict_SetItemString(main_namespace.get(), "blargh", obj.ptr());

    //char* data_ptr=((PyArrayObject*)obj.ptr())->data;
    //char src_ptr[] = {2, 3, 6, 7};
    //memcpy(data_ptr, src_ptr, 4);

    //     list muffins;
    //     muffins.append(1);
    //     muffins.append(2);
    //     muffins.append(3);
    //     muffins.append(4);

    //     PyDict_SetItemString(main_namespace.get(), "muffins", muffins.ptr());

    //Py_Main(argc, argv);
    string buffer; // for multiline commands
    char *ln = 0;
    const char *prompt, *p1 = ">>> ", *p2 = "... ";
    prompt = p1;
    bool multiline = false;
    int slen = 0;
    while ((ln = rl(prompt)))
    {
      buffer += ln;
      buffer.push_back('\n');
      slen = strlen(ln);

      // Ugg, need to check if the last non white space char is a ':',
      // if so, go multiline.
      //if (pos != npos &&
      if (strlen(ln) > 0 && ln[strlen(ln) - 1] == ':')
      {
        prompt = p2;
        multiline = true;
      }

      if ((multiline && strlen(ln) == 0) || (!multiline && strlen(ln) != 0))
      {
        try
        {
          handle<> res(PyRun_String(buffer.c_str(), Py_single_input,
                main_namespace.get(),
                main_namespace.get()));
          //extract<str> strres (res.get());
          //if (strres.check()) {
          //cout << PyString_AsString(res.get());
          //}
        }
        catch (error_already_set)
        {
          PyErr_Print();
        }
        multiline = false;
        prompt = p1;
        buffer.clear();
      }
    }
    cout << endl;
  }
  catch (error_already_set)
  {
    PyErr_Print();
  }

  Py_Finalize();
  return 0;
}
