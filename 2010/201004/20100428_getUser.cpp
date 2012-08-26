/**
 *  @file 2100428_getUser.cpp
 *
 *  @date 2012-3-17
 *  @Author: Bonly
 */
/*
 * 需要你自已在stdafx.h头文件中定义.编译器根据此宏来确定windows的版本,如果你需要使用高版本的WIN32函数,只有你定义了此宏后才能使用;
  Windows   XP                                       _WIN32_WINNT>=0x0501
  Windows   2000                                   _WIN32_WINNT>=0x0500
  Windows   NT   4.0                               _WIN32_WINNT>=0x0400
  Windows   Me                                       _WIN32_WINDOWS=0x0490
  Windows   98                                       _WIN32_WINDOWS>=0x0410
  Internet   Explorer   6.0                 _WIN32_IE>=0x0600
  Internet   Explorer   5.01,   5.5     _WIN32_IE>=0x0501
  Internet   Explorer   5.0,   5.0a,   5.0b   _WIN32_IE>=0x0500
  Internet   Explorer   4.01   _WIN32_IE>=0x0401
  Internet   Explorer   4.0   _WIN32_IE>=0x0400
  Internet   Explorer   3.0,   3.01,   3.02   _WIN32_IE>=0x0300
  */

#define _WIN32_WINNT 0x0501
#include <windows.h>
#include <stdio.h>
#include <process.h>
#include <Tlhelp32.h>
#include <tchar.h>
#include <psapi.h>
#include <stdio.h>
#include <STDLIB.H>
#include <tlhelp32.h>
#include <wtsapi32.h>
//#pragma comment(lib, "WtsApi32.lib")
//#pragma warning( disable:4018 )
//#pragma warning( disable:4996 )
extern HANDLE WTSOpenServer(LPTSTR);
extern void WTSCloseServer(HANDLE);


typedef unsigned int uint;
enum{NAMELEN =64};
typedef struct _my_wts_process_info {
    DWORD       _session_id;
    DWORD       _process_id;
    TCHAR       _process_name[NAMELEN];
    TCHAR       _domain_name[NAMELEN];
    TCHAR       _user_name[NAMELEN];
    SID         _sid;
} MY_WTS_PROCESS_INFO, *PMY_WTS_PROCESS_INFO;

typedef struct _WTS_PROCESS_INFO {
  DWORD  SessionId;
  DWORD  ProcessId;
  LPTSTR pProcessName;
  PSID   pUserSid;
} WTS_PROCESS_INFO, *PWTS_PROCESS_INFO;

extern BOOL WTSEnumerateProcesses(HANDLE hServer, DWORD Reserved,
            DWORD Version, PWTS_PROCESS_INFO *ppProcessInfo, DWORD *pCount);
extern BOOL WTSLogoffSession(
            HANDLE hServer,
            DWORD SessionId,
            BOOL bWait
          );
class CSessionManage
{
public:
    CSessionManage          ( LPTSTR lpszServerName = NULL );
    virtual ~CSessionManage ();

    int     GetSessions     ( PWTS_SESSION_INFO pSessions, uint count );
    int     GetProcesses    ( DWORD dwSessionId, PMY_WTS_PROCESS_INFO pProcesses, uint count );
    int     GetProcesses    ( PMY_WTS_PROCESS_INFO pProcesses, uint count );
    BOOL    GetSessionUser  ( DWORD dwSessionId, LPTSTR lpszUserName );

    BOOL    DisconnectSession   ( DWORD dwSessionId, BOOL bWait );
    BOOL    DisconnectSession   ( BOOL bWait );
    BOOL    LogoffSession       ( DWORD dwSessionId, BOOL bWait );
    BOOL    LogoffSession       ( BOOL bWait );
    BOOL    LogoffUser          ( LPCTSTR lpszUserName, BOOL bWait );

protected:
    CSessionManage  ( const CSessionManage& );
    CSessionManage& operator = ( const CSessionManage& );

    BOOL    OpenServer  ( LPTSTR lpszServerName );
    void    CloseServer ();

private:
    LPTSTR      _server_name;
    HANDLE      _server_handle;
};

//////////////////////////////////////////////////////////////////////////

CSessionManage::CSessionManage( LPTSTR lpszServerName /* = NULL */ )
    : _server_name( NULL )
    , _server_handle( NULL )
{
    OpenServer( lpszServerName );
}

CSessionManage::~CSessionManage()
{
    CloseServer();
}

//////////////////////////////////////////////////////////////////////////

BOOL CSessionManage::OpenServer( LPTSTR lpszServerName )
{
    if( NULL == lpszServerName ) {
        _server_handle = WTS_CURRENT_SERVER_HANDLE;
        return TRUE;
    }

    _server_name = new TCHAR[_tcslen(lpszServerName) + 1];
    if( NULL == _server_name )
        return FALSE;

    _tcscpy( _server_name, lpszServerName );

    _server_handle = WTSOpenServer( lpszServerName );
    if( NULL == _server_handle )
        return FALSE;

    return TRUE;
}

void CSessionManage::CloseServer()
{
    if( NULL != _server_name ) {
        delete []_server_name;
        _server_name = NULL;
    }
    if( WTS_CURRENT_SERVER_HANDLE != _server_handle )
        WTSCloseServer( _server_handle );
    _server_handle = NULL;
}

//////////////////////////////////////////////////////////////////////////

int CSessionManage::GetSessions( PWTS_SESSION_INFO pSessions, uint count )
{
    if( NULL == pSessions )
        return 0;

    PWTS_SESSION_INFO   pSessionInfo    = NULL;
    DWORD               dwCount         = 0;

    if( WTSEnumerateSessions(_server_handle, 0, 1, &pSessionInfo, &dwCount) ) {
        dwCount = (dwCount <= count) ? dwCount : count;
        for( int i = 0; i < dwCount; i++ )
            pSessions[i] = pSessionInfo[i];
    }

    WTSFreeMemory( pSessionInfo );
    return dwCount;
}

int CSessionManage::GetProcesses( PMY_WTS_PROCESS_INFO pProcesses, uint count )
{
    if( NULL == pProcesses )
        return 0;

    PWTS_PROCESS_INFO   pProcessInfo    = NULL;
    DWORD               dwCount         = 0;

    if( WTSEnumerateProcesses(_server_handle, 0, 1, &pProcessInfo, &dwCount) ) {
        dwCount = (dwCount <= count) ? dwCount : count;
        for( int i = 0; i < dwCount; i++ ) {
            // session id
            pProcesses[i]._session_id = pProcessInfo[i].SessionId;

            // process id
            pProcesses[i]._process_id = pProcessInfo[i].ProcessId;

            // SID
            if( NULL != pProcessInfo[i].pUserSid )
                pProcesses[i]._sid = *(SID*)(pProcessInfo[i].pUserSid);
            else
                memset( &(pProcesses[i]._sid), 0, sizeof(SID) );

            // process name
            if( NULL != pProcessInfo[i].pProcessName )
                _tcscpy( pProcesses[i]._process_name, pProcessInfo[i].pProcessName );
            else
                memset( pProcesses[i]._process_name, 0, sizeof(TCHAR) * NAMELEN );

            // domain name and user name
            DWORD           dwNameLen   = NAMELEN;
            SID_NAME_USE    nameuse     = SidTypeUser;
            LookupAccountSid( _server_name, pProcessInfo[i].pUserSid, pProcesses[i]._user_name,
                &dwNameLen, pProcesses[i]._domain_name, &dwNameLen, &nameuse );
        }
    }

    WTSFreeMemory( pProcessInfo );
    return dwCount;
}

int CSessionManage::GetProcesses( DWORD dwSessionId, PMY_WTS_PROCESS_INFO pProcesses, uint count )
{
    if( NULL == pProcesses )
        return 0;

    MY_WTS_PROCESS_INFO pi[512] = { 0 };
    DWORD               dwCount = GetProcesses( pi, 512 );
    int                 rst     = 0;

    for( int i = 0; i < dwCount; i++ ) {
        if( dwSessionId == pi[i]._session_id ) {
            pProcesses[rst++] = pi[i];
            if( rst >= count )
                break;
        }
    }

    return rst;
}

BOOL CSessionManage::GetSessionUser( DWORD dwSessionId, LPTSTR lpszUserName )
{
    if( NULL == lpszUserName )
        return FALSE;

    LPTSTR  lpszName    = NULL;
    DWORD   dwCount     = 0;
    BOOL    bRet        = FALSE;

    if( WTSQuerySessionInformation(_server_handle, dwSessionId, WTSUserName, &lpszName, &dwCount) ) {
        _tcscpy( lpszUserName, lpszName );
        bRet = TRUE;
    }

    WTSFreeMemory( lpszName );
    return bRet;
}

BOOL CSessionManage::DisconnectSession( DWORD dwSessionId, BOOL bWait )
{
    return WTSDisconnectSession( _server_handle, dwSessionId, bWait );
}

BOOL CSessionManage::DisconnectSession( BOOL bWait )
{
    DWORD dwSessionId = 0;
    if( !ProcessIdToSessionId(GetCurrentProcessId(), &dwSessionId) )
        return FALSE;
    return DisconnectSession( dwSessionId, bWait );
}

BOOL CSessionManage::LogoffSession( DWORD dwSessionId, BOOL bWait )
{
    return WTSLogoffSession( _server_handle, dwSessionId, bWait );
}

BOOL CSessionManage::LogoffSession( BOOL bWait )
{
    DWORD dwSessionId = 0;
    if( !ProcessIdToSessionId(GetCurrentProcessId(), &dwSessionId) )
        return FALSE;
    return LogoffSession( dwSessionId, bWait );
}

BOOL CSessionManage::LogoffUser( LPCTSTR lpszUserName, BOOL bWait )
{
    if( NULL == lpszUserName )
        return FALSE;

    PWTS_SESSION_INFO   pSessionInfo    = NULL;
    DWORD               dwCount         = NULL;
    BOOL                bReturn         = FALSE;

    if( WTSEnumerateSessions(_server_handle, 0, 1, &pSessionInfo, &dwCount) ) {
        bReturn = TRUE;
        for( int i = 0; i < dwCount; i++ ) {
            LPTSTR  lpszName    = NULL;
            DWORD   dwRet       = 0;
            if( WTSQuerySessionInformation(_server_handle, pSessionInfo[i].SessionId, WTSUserName,
                &lpszName, &dwRet) ) {
                    if( 0 == _tcscmp(lpszUserName, lpszName) )
                        WTSLogoffSession( _server_handle, pSessionInfo[i].SessionId, bWait );
            }
            WTSFreeMemory( lpszName );
        }
    }

    WTSFreeMemory( pSessionInfo );
    return bReturn;
}

int main()
{
    CSessionManage sm;

    WTS_SESSION_INFO si[16] = { 0 };
    int session_count = sm.GetSessions( si, 16 );
    for( int i = 0; i < session_count; i++ )
        _tprintf( _T("session%02d: %d\n"), i+1, si[i].SessionId );

    printf( "\r\n" );

    MY_WTS_PROCESS_INFO pi[128] = { 0 };
    int process_count = sm.GetProcesses( pi, 128 );
    for( int i = 0; i < process_count; i++ )
        _tprintf( _T("process%02d: %04d %s\n"), i+1, pi[i]._process_id, pi[i]._process_name );

    return 0;
}

