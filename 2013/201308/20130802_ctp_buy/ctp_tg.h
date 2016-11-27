#ifndef __CTP_C_FOR_GO_H__
#define __CTP_C_FOR_GO_H__

#ifdef __cplusplus
extern "C"{
#endif

//from c
void trader();
void market();

#ifdef __cplusplus
}
#endif

#endif

/*
g++ -shared -fPIC market.cpp MarketSpi.cpp trader.cpp TraderSpi.cpp -o libbonly.so 
*/
