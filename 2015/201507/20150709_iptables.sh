#!/bin/bash
#清空所有规则链
iptables -F 
#删除特定手工设置的链
iptables -X 
#清空计数器
iptables -Z 
#默认INPUT规则 丢弃
iptables -P INPUT DROP 
#默认OUTPUT规则 丢弃
iptables -P OUTPUT DROP 
#默认FORWARD规则 丢弃
iptables -P FORWARD DROP 
#让已经建立的连接能通过先
iptables -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT
iptables -A OUTPUT -m state --state ESTABLISHED,RELATED -j ACCEPT

#开SSH服务进站端口
#iptables -A INPUT -d 104.168.15.189 -p tcp -m tcp --dport 22 -j ACCEPT 
iptables -A INPUT -p tcp -m tcp --dport 1024 -j ACCEPT 
#对SSH的服务进入的数据包开启出站端口
#iptables -A OUTPUT -s 104.168.15.189 -p tcp -m tcp --sport 22 -m state --state ESTABLISHED -j ACCEPT 
iptables -A OUTPUT -p tcp -m tcp --sport 1024 -m state --state ESTABLISHED -j ACCEPT 
#去往远程DNS服务器53端口的数据包出站通过
iptables -A OUTPUT -p udp -m udp --dport 53 -j ACCEPT  
#来自远程DNS服务器53端口的数据包进站通过
iptables -A INPUT -p udp -m udp --sport 53 -j ACCEPT  
#对WEB的服务进入的数据包开启出站端口
iptables -A OUTPUT -p tcp --dport 80  -j ACCEPT 
iptables -A OUTPUT -p tcp --dport 443  -j ACCEPT 
iptables -A OUTPUT -p udp --dport 443  -j ACCEPT 
#接受v2ray
# iptables -A INPUT -p tcp --dport 10000:10010 -j ACCEPT 
# iptables -A INPUT -p udp --dport 10000:10010 -j ACCEPT 
# iptables -A OUTPUT -p udp --dport 4040  -j ACCEPT 
#接受ss
# iptables -A INPUT -p tcp --dport 8387 -j ACCEPT  
# iptables -A INPUT -p udp --dport 8387 -j ACCEPT 
# iptables -A INPUT -p tcp --dport 8388 -j ACCEPT  
# iptables -A INPUT -p udp --dport 8388 -j ACCEPT 
# iptables -A INPUT -p tcp --dport 8389 -j ACCEPT  
# iptables -A INPUT -p udp --dport 8389 -j ACCEPT 
iptables -A INPUT -p udp --dport 8390 -j ACCEPT 
iptables -A INPUT -p tcp --dport 8390 -j ACCEPT 
iptables -A INPUT -p tcp --dport 8388 -j ACCEPT  
iptables -A INPUT -p udp --dport 8388 -j ACCEPT 
