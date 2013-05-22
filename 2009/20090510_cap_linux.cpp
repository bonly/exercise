#include <stdio.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netinet/ip.h>
#include <string.h>
#include <netdb.h>
#include <netinet/tcp.h>
#include <netinet/udp.h>
#include <stdlib.h>
#include <unistd.h>
#include <signal.h>
#include <net/if.h>
#include <sys/ioctl.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <linux/if_ether.h>
#include <net/ethernet.h>


void die(char *why, int n)
{
  perror(why);
  exit(n);
}

int do_promisc(char *nif, int sock )
{
struct ifreq ifr;

strncpy(ifr.ifr_name, nif,strlen(nif)+1);
   if((ioctl(sock, SIOCGIFFLAGS, &ifr) == -1)) //���flag

   {
     die("ioctl", 2);
   }

   ifr.ifr_flags |= IFF_PROMISC; //����flag��־


   if(ioctl(sock, SIOCSIFFLAGS, &ifr) == -1 ) //�ı�ģʽ

   {
     die("ioctl", 3);
   }
}
//�޸�������PROMISC(����)ģʽ

char buf[40960];

main()
{
struct sockaddr_in addr;
struct ether_header *peth;
struct iphdr *pip;
struct tcphdr *ptcp;
struct udphdr *pudp;

char mac[16];
int i,sock, r, len;
char *data;
char *ptemp;
char ss[32],dd[32];

if((sock = socket(AF_PACKET, SOCK_RAW, htons(ETH_P_ALL))) == -1) //����socket

//man socket���Կ������漸�������˼

{
        die("socket", 1);
}

do_promisc("eth0", sock); //eth0Ϊ��������



system("ifconfig");

for(;;)
{
     len = sizeof(addr);

     r = recvfrom(sock,(char *)buf,sizeof(buf), 0, (struct sockaddr *)&addr,&len);
     //���Ե�ʱ���������һ�����r������ж��Ƿ�ץ����

     buf[r] = 0;
     ptemp = buf;
     peth = (struct ether_header *)ptemp;

     ptemp += sizeof(struct ether_header); //ָ�����ethͷ�ĳ���

     pip = (struct ip *)ptemp; //pipָ��ip��İ�ͷ


     ptemp += sizeof(struct ip);//ָ�����ipͷ�ĳ���


     switch(pip->protocol) //���ݲ�ͬЭ���ж�ָ������

     {
         case IPPROTO_TCP:
         ptcp = (struct tcphdr *)ptemp; //ptcpָ��tcpͷ��

         printf("TCP pkt :FORM:[%s]:[%d]\n",inet_ntoa(*(struct in_addr*)&(pip->saddr)),ntohs(ptcp->source));
         printf("TCP pkt :TO:[%s]:[%d]\n",inet_ntoa(*(struct in_addr*)&(pip->daddr)),ntohs(ptcp->dest));

         break;

         case IPPROTO_UDP:
         pudp = (struct udphdr *)ptemp; //ptcpָ��udpͷ��

              printf("UDP pkt:\n len:%d payload len:%d from %s:%d to %s:%d\n",
             r,
             ntohs(pudp->len),
             inet_ntoa(*(struct in_addr*)&(pip->saddr)),
             ntohs(pudp->source),
             inet_ntoa(*(struct in_addr*)&(pip->daddr)),
             ntohs(pudp->dest)
         );
         break;

         case IPPROTO_ICMP:
         printf("ICMP pkt:%s\n",inet_ntoa(*(struct in_addr*)&(pip->saddr)));
         break;

         case IPPROTO_IGMP:
         printf("IGMP pkt:\n");
         break;

         default:
         printf("Unkown pkt, protocl:%d\n", pip->protocol);
         break;
    } //end switch


perror("dump");
 }

}

/*
[playmud@fc3 test]$ gcc -v
Reading specs from /usr/lib/gcc/i386-redhat-linux/3.4.2/specs
Configured with: ../configure --prefix=/usr --mandir=/usr/share/man --infodir=/usr/share/info --enable-shared --enable-threads=posix --disable-checking --with-system-zlib --enable-__cxa_atexit --disable-libunwind-exceptions --enable-java-awt=gtk --host=i386-redhat-linux
Thread model: posix
gcc version 3.4.2 20041017 (Red Hat 3.4.2-6.fc3)

************************eth�Ľṹ**************************************
struct ether_header
{
  u_int8_t ether_dhost[ETH_ALEN]; // destination eth addr
  u_int8_t ether_shost[ETH_ALEN]; // source ether addr
  u_int16_t ether_type; // packet type ID field
} __attribute__ ((__packed__));

***********************IP�Ľṹ***********************************
struct iphdr
  {
#if __BYTE_ORDER == __LITTLE_ENDIAN
    unsigned int ihl:4;
    unsigned int version:4;
#elif __BYTE_ORDER == __BIG_ENDIAN
    unsigned int version:4;
    unsigned int ihl:4;
#else
# error "Please fix <bits/endian.h>"
#endif
    u_int8_t tos;
    u_int16_t tot_len;
    u_int16_t id;
    u_int16_t frag_off;
    u_int8_t ttl;
    u_int8_t protocol;
    u_int16_t check;
    u_int32_t saddr;
    u_int32_t daddr;
  };

***********************TCP�Ľṹ****************************
struct tcphdr
  {
    u_int16_t source;
    u_int16_t dest;
    u_int32_t seq;
    u_int32_t ack_seq;
# if __BYTE_ORDER == __LITTLE_ENDIAN
    u_int16_t res1:4;
    u_int16_t doff:4;
    u_int16_t fin:1;
    u_int16_t syn:1;
    u_int16_t rst:1;
    u_int16_t psh:1;
    u_int16_t ack:1;
    u_int16_t urg:1;
    u_int16_t res2:2;
# elif __BYTE_ORDER == __BIG_ENDIAN
    u_int16_t doff:4;
    u_int16_t res1:4;
    u_int16_t res2:2;
    u_int16_t urg:1;
    u_int16_t ack:1;
    u_int16_t psh:1;
    u_int16_t rst:1;
    u_int16_t syn:1;
    u_int16_t fin:1;
# else
# error "Adjust your <bits/endian.h> defines"
# endif
    u_int16_t window;
    u_int16_t check;
    u_int16_t urg_ptr;
};
***********************UDP�Ľṹ*****************************
struct udphdr
{
  u_int16_t source;
  u_int16_t dest;
  u_int16_t len;
  u_int16_t check;
};

*************************************************************
*/

