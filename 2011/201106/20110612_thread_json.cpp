//因为read_json用了isstream,而且使用了spirit，如果要thread safe就需定义，并且recompiled using the -mt libraries 
#define BOOST_SPIRIT_THREADSAFE
#include <boost/property_tree/ptree.hpp>
#include <boost/property_tree/json_parser.hpp>
#include <boost/thread.hpp>
#include <iostream>
using namespace std;
using namespace boost::property_tree;

class testBind {
public:
    void testFunc() {
        cout<<"ok"<<endl;
        string str = "{\"51\":1,\"50\":1}";
        stringstream stream;
        stream<<str;
        ptree tree;
        read_json(stream, tree);
    }
};
/*
 * 
 */
int main(int argc, char** argv) {
    testBind tb;
    boost::thread_group tg;
    for (int i = 0; i < 20; ++i) {
        tg.add_thread(new boost::thread(boost::bind(&testBind::testFunc, &tb)));
    }
    int x;
    cin >> x;

    tg.join_all();


    return 0;
}

/*
g++ 20110612_thread_json.cpp -l boost_thread -lboost_system
*/
