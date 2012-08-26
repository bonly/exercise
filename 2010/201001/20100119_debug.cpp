#include <iostream>
#ifdef __BONLY_DEBUG__    
#define _$ \
    std::cerr<<__FUNCTION__ <<":"<< __FILE__ <<":"<<__LINE__<<std::endl;
#else
#define _$ 
#endif

int main()
{
    _$
    return 0;
}

