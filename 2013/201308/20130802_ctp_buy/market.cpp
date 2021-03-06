#include <iostream>
#include "ThostFtdcMdApi.h"
#include "MarketSpi.h"
#include "ctp_tc.h"
#include "ctp_tg.h"

// UserApi对象
CThostFtdcMdApi* pUserApi;

// 配置参数
/*
模拟境对接 
BrokerID="0292" 
交易：222.240.130.22:41205 
          
行情：222.240.130.22:41213 
   
正式环境对接 
BrokerID="0292" 
交易：180.166.0.225:18993 
          
行情：180.166.0.225:18883 
*/
// char FRONT_ADDR[] = "tcp://zjzx-md1.ctp.shcifco.com:41213";
//char FRONT_ADDR[] = "tcp://210.22.25.180:41213";
//char FRONT_ADDR[] = "tcp://210.22.25.184:41213";
char FRONT_ADDR[] = "tcp://180.168.146.187:10010";

TThostFtdcBrokerIDType	BROKER_ID = "2030";				// 经纪公司代码
TThostFtdcInvestorIDType INVESTOR_ID = "00092";			// 投资者代码
TThostFtdcPasswordType  PASSWORD = "888888";			// 用户密码
char *ppInstrumentID[] = {"TA509"};//,"SR509"};//,"IF1506","cu1205", "cu1206"};			// 行情订阅列表
//char *ppInstrumentID[] = {"TA1605","SR1507"};//,"IF1506","cu1205", "cu1206"};			// 行情订阅列表
// char *ppInstrumentID[] = {"IF300","IH1506"};//,"IF1506","cu1205", "cu1206"};			// 行情订阅列表


int iInstrumentID = 1;									// 行情订阅数量

// 请求编号
int iRequestID = 0;

void market(){
	std::cerr << "======== 启动行情 ========" << std::endl;
	// 初始化UserApi
	pUserApi = CThostFtdcMdApi::CreateFtdcMdApi();			// 创建UserApi
	CThostFtdcMdSpi* pUserSpi = new CMdSpi();
	pUserApi->RegisterSpi(pUserSpi);						// 注册事件类
	pUserApi->RegisterFront(FRONT_ADDR);					// connect
	pUserApi->Init();

	pUserApi->Join();
//	pUserApi->Release();
    std::cerr << "========= 结束行情 ========" << std::endl;
}


/*
服务器组1：上海域名站点
交易服务器地址：                端口
ztqh-front1.china-invf.com      41205
ztqh-front2.china-invf.com      41205
行情服务器地址：
ztqh-md1.china-invf.com         41213
ztqh-md2.china-invf.com         41213
 
服务器组2：上海IP站点        
交易服务器地址：                端口
180.166.45.50                   41205
27.115.57.180                   41205
行情服务器地址：
180.166.45.50                   41213
27.115.57.180                   41213
 
服务器组3：北京IP站点
交易服务器地址：                端口
114.255.13.180                  41205
114.255.13.181                  41205
行情服务器地址：
114.255.13.180                  41213
114.255.13.181                  41213
*/
