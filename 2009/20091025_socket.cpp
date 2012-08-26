#include <unistd.h>
#include <signal.h>
#include <errno.h>
#include <iostream>
#include "net/Socket.h"

using namespace std;
using namespace Network;

namespace NetSignalHandler
{
	bool set_handler = false;
	void signal_handler(int signal_code)
	{
#ifdef _DEBUG
		cout << "ignore signal SIGPIPE" << endl;
#endif
		::signal(SIGPIPE, signal_handler);
	}
};

Socket::Socket()
{
	if (!NetSignalHandler::set_handler)
	{
		::signal(SIGPIPE, NetSignalHandler::signal_handler);
		NetSignalHandler::set_handler = true;
	}
	m_hSocket = INVALID_SOCKET;
}
Socket::~Socket()
{
	close();
}

bool Socket::create(u_int nSocketPort, int nSocketType, const char* lpszSocketAddress)
{
	if (socket(nSocketType))
	{
		if (nSocketPort != 0)
		{
			int reuse_addr = 1;
			setSockOpt(SO_REUSEADDR, &reuse_addr, sizeof(int));
		}
		if (bind(nSocketPort, lpszSocketAddress))
			return true;
		int nResult = errno;
		close();
		errno = nResult;
	}
	return false;
}

bool Socket::getPeerName(string& rPeerAddress, u_int& rPeerPort)
{
	SOCKADDR_IN sockAddr;
	memset(&sockAddr, 0, sizeof(sockAddr));

	int nSockAddrLen = sizeof(sockAddr);
	bool bResult = getPeerName((SOCKADDR*)&sockAddr, &nSockAddrLen);
	if (bResult)
	{
		rPeerPort = ntohs(sockAddr.sin_port);
		rPeerAddress = inet_ntoa(sockAddr.sin_addr);
	}
	return bResult;
}
bool Socket::getSockName(string& rPeerAddress, u_int& rPeerPort)
{
	SOCKADDR_IN sockAddr;
	memset(&sockAddr, 0, sizeof(sockAddr));

	int nSockAddrLen = sizeof(sockAddr);
	bool bResult = getSockName((SOCKADDR*)&sockAddr, &nSockAddrLen);
	if (bResult)
	{
		rPeerPort = ntohs(sockAddr.sin_port);
		rPeerAddress = inet_ntoa(sockAddr.sin_addr);
	}
	return bResult;
}

bool Socket::socket(int nSocketType, int nProtocolType, int nAddressFormat)
{
	if (m_hSocket != INVALID_SOCKET)
		return false;
	
	m_hSocket = ::socket(nAddressFormat, nSocketType, nProtocolType);
	if (m_hSocket == INVALID_SOCKET)
		return false;

	return true;
}
bool Socket::bind(u_int nSocketPort, const char* lpszSocketAddress)
{
	SOCKADDR_IN sockAddr;
	memset(&sockAddr, 0, sizeof(sockAddr));

	sockAddr.sin_family = AF_INET;

	if (lpszSocketAddress == NULL)
		sockAddr.sin_addr.s_addr = htonl(INADDR_ANY);
	else
	{
		u_int lResult = inet_addr(lpszSocketAddress);
		if (lResult == INADDR_NONE)
		{
			errno = EINVAL;
			return false;
		}
		sockAddr.sin_addr.s_addr = lResult;
	}

	sockAddr.sin_port = htons((u_short)nSocketPort);

	bool bResult = bind((SOCKADDR*)&sockAddr, sizeof(sockAddr));
	return bResult;
}

bool Socket::connect(const string& lpszHostAddress, u_int nHostPort)
{
	if (lpszHostAddress.empty())
		return false;

	SOCKADDR_IN sockAddr;
	memset(&sockAddr, 0, sizeof(sockAddr));

	sockAddr.sin_family = AF_INET;
	sockAddr.sin_addr.s_addr = inet_addr(lpszHostAddress.c_str());

	if (sockAddr.sin_addr.s_addr == INADDR_NONE)
	{
		PHOSTENT lphost = gethostbyname(lpszHostAddress.c_str());
		if (lphost != NULL)
			sockAddr.sin_addr.s_addr = ((PIN_ADDR)lphost->h_addr)->s_addr;
		else
		{
			errno = EINVAL;
			return false;
		}
	}

	sockAddr.sin_port = htons((u_short)nHostPort);

	bool bResult = connect((SOCKADDR*)&sockAddr, sizeof(sockAddr));
	return bResult;
}
bool Socket::accept(Socket& rConnectedSocket, SOCKADDR* lpSockAddr, int* lpSockAddrLen)
{
	if (rConnectedSocket.m_hSocket != INVALID_SOCKET)
		return false;

	SOCKET hTemp = ::accept(m_hSocket, lpSockAddr, (socklen_t*)lpSockAddrLen);

	if (hTemp == INVALID_SOCKET)
		rConnectedSocket.m_hSocket = INVALID_SOCKET;
	else
	{
		rConnectedSocket.m_hSocket = hTemp;
//		LINGER accept_linger;
//		accept_linger.l_onoff = 1;
//		accept_linger.l_linger = 0;
//		rConnectedSocket.setSockOpt(SO_LINGER, &accept_linger, sizeof(accept_linger));
	}

	return (hTemp != INVALID_SOCKET);
}
void Socket::close()
{
	if (m_hSocket != INVALID_SOCKET)
	{
		::close(m_hSocket);
		m_hSocket = INVALID_SOCKET;
	}
}

int Socket::receive(void* lpBuf, int nBufLen, int nFlags)
{
	return ( ::recv(m_hSocket, (char *)lpBuf, nBufLen, nFlags) );
}
int Socket::receiveFrom(void* lpBuf, int nBufLen, string& rSocketAddress, u_int& rSocketPort, int nFlags)
{
	SOCKADDR_IN sockAddr;

	memset(&sockAddr, 0, sizeof(sockAddr));

	int nSockAddrLen = sizeof(sockAddr);
	int nResult = receiveFrom(lpBuf, nBufLen, (SOCKADDR*)&sockAddr, &nSockAddrLen, nFlags);
	if(nResult != SOCKET_ERROR)
	{
		rSocketPort = ntohs(sockAddr.sin_port);
		rSocketAddress = inet_ntoa(sockAddr.sin_addr);
	}
	return nResult;
}
int Socket::send(const void* lpBuf, int nBufLen, int nFlags)
{
	int send_len = ::send(m_hSocket, (char *)lpBuf, nBufLen, nFlags);
#ifndef PEER_SET_LINGER
	if (send_len > 0)
	{
		if (::send(m_hSocket, NULL, 0, 0) == SOCKET_ERROR)
			return SOCKET_ERROR;
	}
#endif
	return send_len;
}
int Socket::sendTo(const void* lpBuf, int nBufLen, u_int nHostPort, const string& lpszHostAddress, int nFlags)
{
	SOCKADDR_IN sockAddr;

	memset(&sockAddr, 0, sizeof(sockAddr));

	sockAddr.sin_family = AF_INET;

	if (lpszHostAddress.empty())
		sockAddr.sin_addr.s_addr = htonl(INADDR_BROADCAST);
	else
	{
		sockAddr.sin_addr.s_addr = inet_addr(lpszHostAddress.c_str());
		if (sockAddr.sin_addr.s_addr == INADDR_NONE)
		{
			PHOSTENT lphost;
			lphost = gethostbyname(lpszHostAddress.c_str());
			if (lphost != NULL)
				sockAddr.sin_addr.s_addr = ((PIN_ADDR)lphost->h_addr)->s_addr;
			else
			{
				errno = EINVAL;
				return SOCKET_ERROR;
			}
		}
	}

	sockAddr.sin_port = htons((u_short)nHostPort);

	int iResult = sendTo(lpBuf, nBufLen, (SOCKADDR*)&sockAddr, sizeof(sockAddr), nFlags);
	return iResult;
}

// �жϽ��ճ�ʱʱ��
int Socket::canReceive(int nTimeoutSeconds)
{
	fd_set readfds;
	FD_ZERO(&readfds);
	FD_SET(m_hSocket, &readfds);
	TIMEVAL tv = { nTimeoutSeconds, 0 };
	
	int ret = ::select(m_hSocket+1, &readfds, NULL, NULL, &tv);
	if ( ret > 0 )
		return FD_ISSET(m_hSocket, &readfds);

	return ret;
}