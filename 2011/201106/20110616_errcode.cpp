#define ALL_CODE \
  CODE_OBJ_V(SUCCES,0) \
  CODE_OBJ(DB_ERR) \
  CODE_OBJ(CFG_ERR) \
  CODE_OBJ_V(NET_ERR,20)

#define CODE_OBJ(a) a,
#define CODE_OBJ_V(a,v) a=v,
enum Code { ALL_CODE };
#undef CODE_OBJ
#undef CODE_OBJ_V

#define CODE_OBJ(a) {a, #a},
#define CODE_OBJ_V(a,v) {a, #a},
struct {
  Code val;
  const char * msg;
} static all_code[] = { ALL_CODE };  ///不能中间隔空代码啊
#undef CODE_OBJ
#undef CODE_OBJ_V

#include <iostream>
int main(){
  std::clog << all_code[CFG_ERR].msg << std::endl;
  std::clog << all_code[DB_ERR].msg << std::endl;
  std::clog << all_code[NET_ERR].msg << std::endl;
  return 0;
}

