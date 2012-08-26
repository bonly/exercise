/**
 *  @file 20100504_download.cpp
 *  @brief 悠米平台的下载代码,不能单独运行,只供分析
 *  @date 2012-3-21
 *  @author Bonly
 */
//dll下载相关
/**
 * 初始化下载所需的观察器对象
 * @param aApnId
 */
void CJggWin::InitDloaderDll(int aApnId)
{
    if (aApnId < 0)
    {
        aApnId = 0;
    }
    if (!m_pDloaderDllTbl)
    {
        AEESimulator_LoadDLoader(NULL, NULL, (void**)&m_pDloaderDllTbl);
        if (m_pDloaderDllTbl)
        {
            m_pMemBuf = (char*)SafeMalloc(10 * 1024);  ///创建内存块
            CDLoaderObserv ob;  ///观察器
            MEMSET(&ob,0,sizeof(CDLoaderObserv));
            ob.iMemBuf = m_pMemBuf; ///把创建的内存块给观察器
            ob.iMemBufLen = 10 * 1024;

            STRCPY(ob.iAppVer,KUmvchatSoftVer);

            int index = 0;
            int count = 0;
            char apprbd[32];
            MEMSET(apprbd,0,32);
            while(index < STRLEN(UmvChatBuildDate) )
            {
                if ( '0' <= UmvChatBuildDate[index] && UmvChatBuildDate[index] <= '9')
                {
                    apprbd[count] = UmvChatBuildDate[index];
                    count++;
                    if ( count >= 10)
                    {
                        break;
                    }
                }
                index++;
            }
            STRCPY(ob.iAppRbd,apprbd);  ///包头10位是程序版本号,格式为2010091401

            ob.iIapId = aApnId;  ///APN
            ob.iIsCmwap = IsCmWap(ob.iIapId);  ///设置是否为CMWAP的APN

            STRCPY(ob.iTempDir,"."); ///设置下载目录

            /**
             * 给各个回调函数设置值
             */
            ob.DloadRecvItems = DloadRecvItems;
            ob.DloadOne = DloadOne;
            ob.DloadError = DloadError;
            ob.DloadPerNotify = DloadPerNotify;

            ///下载列表中创建此下载的观察器
            m_pDloader = m_pDloaderDllTbl->Create(m_pMainGlobal->m_pMe,&ob,this);
        }
    }
}
/**
 * 释放资源
 */
void CJggWin::UninitDloaderDll()
{
    if (m_pDloaderDllTbl)
    {
        m_pDloaderDllTbl->Destroy(m_pDloader);

        FREE(m_pDloaderDllTbl);
        m_pDloaderDllTbl = NULL;
        m_pDloaderDllHandler = NULL;
        m_pDloader = NULL;

        m_pDloader = NULL;
        SafeFree(m_pMemBuf);
        m_pMemBuf = 0;
    }
}

/**
 * 接收到一项内容时回调
 * @param aUserData
 * @param aItems
 * @param aMsg
 */
void CJggWin::DloadRecvItems( void* aUserData, SRecvItem* aItems,AECHAR* aMsg )
{
#if defined(DEBUG_DLOADER)
    DBGPRINTF("CJggWin::DloadRecvItems = %p,%p\r\n",aItems,aMsg);
#endif
    CJggWin* pThis = (CJggWin*)aUserData;
    if (pThis)
    {
        //有更新,aMsg为更新提示
        if(aItems)
        {
        }
        //无更新，异步释放dloader模块
        else
        {
            pThis->PostMsg(EPostMsg_UnInitDloader,0,0,1);
        }
    }
}
/**
 * 回调函数
 * @param aUserData
 * @param aItem
 */
void CJggWin::DloadOne( void* aUserData,SRecvItem* aItem )
{
    CJggWin* pThis = (CJggWin*)aUserData;
    if (pThis)
    {
    }
    return;
}

/**
 * 下载出错时回调
 * @param aUserData
 * @param aItem
 * @param aErrorCode
 */
void CJggWin::DloadError( void* aUserData,SRecvItem* aItem,int aErrorCode)
{
    CJggWin* pThis = (CJggWin*)aUserData;
    if (pThis)
    {
        //出错，异步释放dloader模块
        pThis->PostMsg(EPostMsg_UnInitDloader,0,0,1);
    }
}
/**
 * 下载中的百分比提示
 * @param aUserData
 * @param aItem
 * @param aCurSize
 * @param aTotalSize
 * @param aCbType
 */
void CJggWin::DloadPerNotify( void* aUserData, SRecvItem* aItem, int aCurSize, int aTotalSize, ECbType aCbType)
{
    CJggWin* pThis = (CJggWin*)aUserData;
    if (pThis)
    {
        //下载百分比
        if (aCbType == ECbType_Down)
        {
            int per = aCurSize * 100 / aTotalSize;
        }
        //安装百分比
        else if (aCbType == ECbType_Unzip)
        {
            int per = aCurSize * 100 / aTotalSize;
        }
    }
}


/**
 * 检查更新,此文件的所有函数调用的入口
 * @param aApnId
 * @return
 */
BOOL CJggWin::CheckUpdate(int aApnId)
{
    int index = 0;
    int count = 0;
    char rbd[32];
    MEMSET(rbd,0,32);

    while(index < STRLEN(UmvChatBuildDate) )
    {
        if ( '0' <= UmvChatBuildDate[index] && UmvChatBuildDate[index] <= '9')
        {
            rbd[count] = UmvChatBuildDate[index];
            count++;
            if ( count >= 10)
            {
                break;
            }
        }
        index++;
    }

    char ver[32];
    MEMSET(ver,0,32);

    int len = STRLEN(KUmvchatSoftVer);
    index = 0;
    count = 0;
    STRCPY(ver,"0x");
    count += 2;
    while ( index < len )
    {
        if ( '.' != KUmvchatSoftVer[index] )
        {
            ver[count] = KUmvchatSoftVer[index];
            count++;
        }
        index++;
    }
    while (count < 10)
    {
        ver[count] = '0';
        count++;
    }

    SDloadItem item;
    ZeroMemory(&item,sizeof(item));
    STRCPY(item.name,"0x0300009a");
    STRCPY(item.path,".");
    item.ver = VcAppGlobal::strtoi(ver);
    item.rbd = VcAppGlobal::strtoi(rbd);
    InitDloaderDll(aApnId); ///构造观察器中的数据
    if(m_pDloaderDllTbl)
    {
        m_pDloaderDllTbl->AddDloadItem(m_pDloader,0,item.rbd,item.ver,item.name,item.path); ///加入需下载的项目
        m_pDloaderDllTbl->StartDload(m_pDloader); ///开始下载
    }

    return 1;
}






