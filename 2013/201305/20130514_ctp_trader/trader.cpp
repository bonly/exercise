#include "trader.h"
// testTraderApi.cpp : 定义控制台应用程序的入口点。
//
#include "ThostFtdcTraderApi.h"
#include "TraderSpi.h"

// UserApi对象
CThostFtdcTraderApi* pUserApi;

// 配置参数
// char  FRONT_ADDR[] = "tcp://asp-sim2-front1.financial-trading-platform.com:26205";		// 前置地址
char  FRONT_ADDR[] = "tcp://ztqh-md2.china-invf.com:41213";

TThostFtdcBrokerIDType	BROKER_ID = "2030";				// 经纪公司代码
TThostFtdcInvestorIDType INVESTOR_ID = "354079";			// 投资者代码
TThostFtdcPasswordType  PASSWORD = "481531";			// 用户密码
TThostFtdcInstrumentIDType INSTRUMENT_ID = "cu1408";	// 合约代码
TThostFtdcDirectionType	DIRECTION = THOST_FTDC_D_Sell;	// 买卖方向
TThostFtdcPriceType	LIMIT_PRICE = 60000; // 38850;				// 价格

// 请求编号
int iRequestID = 0;

void cnt(){
	// 初始化UserApi
	pUserApi = CThostFtdcTraderApi::CreateFtdcTraderApi();			// 创建UserApi
	CTraderSpi* pUserSpi = new CTraderSpi();
	pUserApi->RegisterSpi((CThostFtdcTraderSpi*)pUserSpi);			// 注册事件类
	pUserApi->SubscribePublicTopic(THOST_TERT_QUICK);				// 注册公有流
	pUserApi->SubscribePrivateTopic(THOST_TERT_QUICK);				// 注册私有流
	pUserApi->RegisterFront(FRONT_ADDR);							// connect
	pUserApi->Init();

	pUserApi->Join();
//	pUserApi->Release();
}

/*
g++ -shared -fPIC trader.cpp TraderSpi.cpp -o libbonly.so 
*/