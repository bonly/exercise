#include <algorithm>
#include <string>
#include <sstream>

#include <boost/property_tree/ptree.hpp>
#include <boost/property_tree/json_parser.hpp> 

using namespace std;
using namespace boost::property_tree;

int _tmain(int argc, _TCHAR* argv[])
{
    try
    {
        std::string j("{ \"Object1\" : { \"param1\" : 10.0, \"initPos\" : { \"\":1.0, \"\":2.0, \"\":5.0 }, \"initVel\" : [ 0.0, 0.0, 0.0 ] } }");
        std::istringstream iss(j);

        ptree pt;
        json_parser::read_json(iss, pt);

        auto s = pt.get<std::string>("Object1.param1");
        cout << s << endl; // 10

        ptree& pos = pt.get_child("Object1.initPos");
        std::for_each(std::begin(pos), std::end(pos), [](ptree::value_type& kv) { 
            cout << "K: " << kv.first << endl;
            cout << "V: " << kv.second.get<std::string>("") << endl;
        });
    }
    catch(std::exception& ex)
    {
        std::cout << "ERR:" << ex.what() << endl;
    }

    return 0;
}

