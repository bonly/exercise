#include <iostream>
using namespace std;

#define DECL_VAR(type, name, tn)      \
    namespace FLAG_##tn##_inst {      \
        type FLAGS_##name;    \
    }                                 \
    using FLAG_##tn##_inst::FLAGS_##name

DECL_VAR(bool, test, bool);
int main()
{
   FLAGS_test = true;
   return 0;
}

