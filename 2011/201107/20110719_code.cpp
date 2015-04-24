#include <iostream>
#include <string>
#include <map>
#include <boost/system/error_code.hpp>

#define ALL_CODE \
  CODE_OBJ_V(SUCCES,0) \
  CODE_OBJ(DB_ERR) \
  CODE_OBJ(CFG_ERR) \
  CODE_OBJ_V(NET_ERR,20)
  
#define CODE_OBJ(a) a,
#define CODE_OBJ_V(a,v) a=v,
enum class Code { ALL_CODE };
#undef CODE_OBJ
#undef CODE_OBJ_V

template <typename Enumeration>
auto asInt(Enumeration const value)
-> typename std::underlying_type<Enumeration>::type
{
   return static_cast<typename std::underlying_type<Enumeration>::type>(value);
}

#define TypeName(x) #x
#define MakeCode(x, y) \
  std::make_pair<int, std::string>(asInt(x), #y)
    
#define CODE_OBJ(a) {code.push_back(MakeCode(a, TypeName(a))}
#define CODE_OBJ_V(a, v) {code.push_back(MakeCode(v, TypeName(a))}
#define CODE_OBJ_S(a, v, s) {code.push_back(MakeCode(v, s)}
    
class Code : public boost::system::error_category {
public:
    virtual const char *name() const BOOST_SYSTEM_NOEXCEPT{ 
      return "APP"; 
    }
    std::string message(int ev) const {
        //return TypeName(ev);
        //return typeid(ECode::OK).name();
    }
    std::map code;
      
    Code(){
      code
    }
};

Code mycode;

int main(){
  boost::system::error_code ec(14, mycode); 
  std::cout << ec.value() << std::endl; 
  std::cout << ec.category().name() << std::endl; 
  std::cout << ec.message() << std::endl; 
   return 0;
}



