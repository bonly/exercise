#https://docs.google.com/document/d/1mmMiMYbviMxJ-DhTyIGdK7OOg581LSD1CZV4XY1OMG8/mobilebasic?pli=1#h.5ey2khz7nphu
from netfilterqueue import NetfilterQueue
import subprocess
import signal
import dpkt
import traceback
import socket
import sys
 
DNS_IP = '8.8.8.8'
 
# source http://zh.wikipedia.org/wiki/%E5%9F%9F%E5%90%8D%E6%9C%8D%E5%8A%A1%E5%99%A8%E7%BC%93%E5%AD%98%E6%B1%A1%E6%9F%93
WRONG_ANSWERS = {
   '4.36.66.178',
   '8.7.198.45',
   '37.61.54.158',
   '46.82.174.68',
   '59.24.3.173',
   '64.33.88.161',
   '64.33.99.47',
   '64.66.163.251',
   '65.104.202.252',
   '65.160.219.113',
   '66.45.252.237',
   '72.14.205.99',
   '72.14.205.104',
   '78.16.49.15',
   '93.46.8.89',
   '128.121.126.139',
   '159.106.121.75',
   '169.132.13.103',
   '192.67.198.6',
   '202.106.1.2',
   '202.181.7.85',
   '203.161.230.171',
   '207.12.88.98',
   '208.56.31.43',
   '209.36.73.33',
   '209.145.54.50',
   '209.220.30.174',
   '211.94.66.147',
   '213.169.251.35',
   '216.221.188.182',
   '216.234.179.13'
}
 
current_ttl = 1
 
def locate_dns_hijacking(nfqueue_element):
   global current_ttl
   try:
       ip_packet = dpkt.ip.IP(nfqueue_element.get_payload())
       if dpkt.ip.IP_PROTO_ICMP == ip_packet['p']:
           print(socket.inet_ntoa(ip_packet.src))
       elif dpkt.ip.IP_PROTO_UDP == ip_packet['p']:
           if DNS_IP == socket.inet_ntoa(ip_packet.dst):
               ip_packet.ttl = current_ttl
               current_ttl += 1
               ip_packet.sum = 0
               nfqueue_element.set_payload(str(ip_packet))
           else:
               if contains_wrong_answer(dpkt.dns.DNS(ip_packet.udp.data)):
                   sys.stdout.write('* ')
                   sys.stdout.flush()
                   nfqueue_element.drop()
                   return
               else:
                   print('END')
       nfqueue_element.accept()
   except:
       traceback.print_exc()
       nfqueue_element.accept()
 
 
def contains_wrong_answer(dns_packet):
   for answer in dns_packet.an:
       if socket.inet_ntoa(answer['rdata']) in WRONG_ANSWERS:
           return True
   return False
 
nfqueue = NetfilterQueue()
nfqueue.bind(0, locate_dns_hijacking)
 
def clean_up(*args):
   subprocess.call('iptables -D OUTPUT -p udp --dst %s -j QUEUE' % DNS_IP, shell=True)
   subprocess.call('iptables -D INPUT -p udp --src %s -j QUEUE' % DNS_IP, shell=True)
   subprocess.call('iptables -D INPUT -p icmp -m icmp --icmp-type 11 -j QUEUE', shell=True)
 
signal.signal(signal.SIGINT, clean_up)
 
try:
   subprocess.call('iptables -I INPUT -p icmp -m icmp --icmp-type 11 -j QUEUE', shell=True)
   subprocess.call('iptables -I INPUT -p udp --src %s -j QUEUE' % DNS_IP, shell=True)
   subprocess.call('iptables -I OUTPUT -p udp --dst %s -j QUEUE' % DNS_IP, shell=True)
   print('running..')
   nfqueue.run()
except KeyboardInterrupt:
   print('bye')