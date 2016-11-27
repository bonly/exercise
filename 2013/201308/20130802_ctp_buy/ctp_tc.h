#ifndef __CTP_GO_FOR_C_H__
#define __CTP_GO_FOR_C_H__

#ifdef __cplusplus
extern "C"{
#endif

//from go
void data2hst(CThostFtdcDepthMarketDataField *);
void data2db(CThostFtdcDepthMarketDataField *);
char* gb2utf8(char*);

#ifdef __cplusplus
}
#endif

#define D() std::clog << __FILE__ << ":" << __FUNCTION__ << "():" << __LINE__ << std::endl;

#endif

/*
g++ -shared -fPIC market.cpp MarketSpi.cpp trader.cpp TraderSpi.cpp -o libbonly.so 
*/

