#include <iostream>
enum class RestCode { OK, ADD_CPU_ERROR=3 };  

template <typename Enumeration>
auto as_integer(Enumeration const value)
-> typename std::underlying_type<Enumeration>::type
{
   return static_cast<typename std::underlying_type<Enumeration>::type>(value);
}

#define quote(x) #x

int main(){
  std::clog << as_integer(RestCode::ADD_CPU_ERROR);
  std::clog << quote(ADD_CPU_ERROR);
  return 0;
}

