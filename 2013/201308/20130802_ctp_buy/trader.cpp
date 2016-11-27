// testTraderApi.cpp : 定义控制台应用程序的入口点。
//
#include <iostream>
#include "ThostFtdcTraderApi.h"
#include "TraderSpi.h"
#include "ctp_tc.h"
#include "ctp_tg.h"

// UserApi对象
CThostFtdcTraderApi* ptUserApi;

// 配置参数
// char  T_FRONT_ADDR[] = "tcp://asp-sim2-front1.financial-trading-platform.com:26205";		// 前置地址
char  T_FRONT_ADDR[] = "tcp://180.168.146.187:10000";
// TThostFtdcBrokerIDType	T_BROKER_ID = "2030";				// 经纪公司代码
TThostFtdcBrokerIDType	T_BROKER_ID = "9999";				// 经纪公司代码
// TThostFtdcInvestorIDType T_INVESTOR_ID = "354079";			// 投资者代码
TThostFtdcInvestorIDType T_INVESTOR_ID = "017436";			// 投资者代码
// TThostFtdcPasswordType  T_PASSWORD = "481531";			// 用户密码
TThostFtdcPasswordType  T_PASSWORD = "hay111";			// 用户密码
TThostFtdcInstrumentIDType INSTRUMENT_ID = "TF1509";	// 合约代码
TThostFtdcDirectionType	DIRECTION = THOST_FTDC_D_Sell;	// 买卖方向
TThostFtdcPriceType	LIMIT_PRICE = 40810; // 38850;				// 价格

// 请求编号
int itRequestID = 0;

void trader(){
	std::cerr << "======== 启动交易 ========" << std::endl;
	// 初始化UserApi
	ptUserApi = CThostFtdcTraderApi::CreateFtdcTraderApi();			// 创建UserApi
	CTraderSpi* ptUserSpi = new CTraderSpi();
	ptUserApi->RegisterSpi((CThostFtdcTraderSpi*)ptUserSpi);			// 注册事件类
	ptUserApi->SubscribePublicTopic(THOST_TERT_QUICK);				// 注册公有流
	ptUserApi->SubscribePrivateTopic(THOST_TERT_QUICK);				// 注册私有流
	ptUserApi->RegisterFront(T_FRONT_ADDR);							// connect
	ptUserApi->Init();

	ptUserApi->Join();
//	ptUserApi->Release();
	std::cerr << "======== 结束交易 ========" << std::endl;
}