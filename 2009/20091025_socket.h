#ifndef _NET_SOCKET_H_
#define _NET_SOCKET_H_

#include <netdb.h>
#include <sys/time.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <sys/ioctl.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <string>

namespace Network
{
	typedef struct sockaddr		SOCKADDR, *PSOCKADDR;
	typedef struct sockaddr_in	SOCKADDR_IN, *PSOCKADDR_IN;
	typedef struct linger		LINGER, *PLINGER;
	typedef struct in_addr		IN_ADDR, *PIN_ADDR;
	typedef struct hostent		HOSTENT, *PHOSTENT;
	typedef struct servent		SERVENT, *PSERVENT;
	typedef struct protoent		PROTOENT, *PPROTOENT;
	typedef struct timeval		TIMEVAL, *PTIMEVAL;
	
	// Socket handler
	typedef int  SOCKET;
	
	const SOCKET INVALID_SOCKET = (SOCKET)~0;
	const SOCKET SOCKET_ERROR   = -1;
	
	class Socket
	{
	public:
	// Construction
		Socket();
		virtual ~Socket();
		
		bool create(u_int nSocketPort =0, int nSocketType =SOCK_STREAM, const char* lpszSocketAddress =NULL);
		operator SOCKET() const
			{ return m_hSocket; }
		Socket& operator =(const Socket& rSocket)
		{
			if (this != &rSocket)
				m_hSocket = rSocket.m_hSocket;
			return *this;
		}
	
	public:
	// Attributes
		static Socket fromHandle(SOCKET hSocket)
			{ return Socket(hSocket); }
		
		bool getPeerName(std::string& rPeerAddress, u_int& rPeerPort);
		bool getPeerName(SOCKADDR* lpSockAddr, int* lpSockAddrLen)
			{ return (SOCKET_ERROR != getpeername(m_hSocket, lpSockAddr, (socklen_t*)lpSockAddrLen)); }
		bool getSockName(std::string& rPeerAddress, u_int& rPeerPort);
		bool getSockName(SOCKADDR* lpSockAddr, int* lpSockAddrLen)
			{ return (SOCKET_ERROR != getsockname(m_hSocket, lpSockAddr, (socklen_t*)lpSockAddrLen)); }
		bool getSockOpt(int nOptionName, void* lpOptionValue, int* lpOptionLen, int nLevel =SOL_SOCKET)
			{ return (SOCKET_ERROR != getsockopt(m_hSocket, nLevel, nOptionName, (char *)lpOptionValue, (socklen_t*)lpOptionLen)); }
		bool setSockOpt(int nOptionName, const void* lpOptionValue, int nOptionLen, int nLevel =SOL_SOCKET)
			{ return (SOCKET_ERROR != setsockopt(m_hSocket, nLevel, nOptionName, (const char*)lpOptionValue, nOptionLen)); }

//		enum // SOCKET error codes
//		{
//		};
//		static int getLastError()
//			{ return m_nLastSocketError; }
		
	public:
	// Operations
		bool socket(int nSocketType =SOCK_STREAM, 
			int nProtocolType =0, int nAddressFormat =PF_INET);
		bool bind(u_int nSocketPort, const char* lpszSocketAddress =NULL);
		bool bind(const SOCKADDR* lpSockAddr, int nSockAddrLen)
			{ return (SOCKET_ERROR != ::bind(m_hSocket, lpSockAddr, nSockAddrLen)); }
		
		bool connect(const std::string& lpszHostAddress, u_int nHostPort);
		bool connect(const SOCKADDR* lpSockAddr, int nSockAddrLen)
			{ return (SOCKET_ERROR != ::connect(m_hSocket, lpSockAddr, nSockAddrLen)); }
		bool listen(int nConnectionBacklog =5)
			{ return (SOCKET_ERROR != ::listen(m_hSocket, nConnectionBacklog)); }
		virtual bool accept(Socket& rConnectedSocket, SOCKADDR* lpSockAddr =NULL, int* lpSockAddrLen =NULL);
		virtual void close();

		virtual int receive(void* lpBuf, int nBufLen, int nFlags =0);
		int receiveFrom(void* lpBuf, int nBufLen, 
				std::string& rSocketAddress, u_int& rSocketPort, int nFlags =0);
		int receiveFrom(void* lpBuf, int nBufLen, 
				SOCKADDR* lpSockAddr, int* lpSockAddrLen, int nFlags =0)
			{ return (SOCKET_ERROR != ::recvfrom(m_hSocket, (char *)lpBuf, nBufLen, nFlags, lpSockAddr, (socklen_t*)lpSockAddrLen)); }
		virtual int send(const void* lpBuf, int nBufLen, int nFlags =0);
		int sendTo(const void* lpBuf, int nBufLen, 
				u_int nHostPort, const std::string& lpszHostAddress =NULL, int nFlags =0);
		int sendTo(const void* lpBuf, int nBufLen, 
				const SOCKADDR* lpSockAddr, int nSockAddrLen, int nFlags =0)
			{ return (SOCKET_ERROR != ::sendto(m_hSocket, (char *)lpBuf, nBufLen, nFlags, lpSockAddr, nSockAddrLen)); }
		
		enum { receives =0, sends =1, both =2 };
		bool ioctl(long lCommand, u_int* lpArgument)
			{ return (SOCKET_ERROR != ::ioctl(m_hSocket, lCommand, lpArgument)); }
		bool shutdown(int nHow =sends)
			{ return (SOCKET_ERROR != ::shutdown(m_hSocket,nHow)); }
		// 判断接收超时时间
		int canReceive(int nTimeoutSeconds =60);
	
	protected:
		Socket(SOCKET hSocket)
			{ m_hSocket = hSocket; }
		
	// Data Members
		SOCKET m_hSocket;
//		static int m_nLastSocketError;
	};
};

#endif