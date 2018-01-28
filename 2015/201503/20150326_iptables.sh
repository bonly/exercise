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
#开SSH服务进站端口
iptables -A INPUT -d 104.168.15.189 -p tcp -m tcp --dport 1024 -j ACCEPT 
#开WEB服务进站端口
iptables -A INPUT -d 104.168.15.189 -p tcp -m tcp --dport 80 -j ACCEPT 
#允许本地环回数据
iptables -A INPUT -s 127.0.0.1 -d 127.0.0.1 -j ACCEPT 
#来自远程DNS服务器53端口的数据包进站通过
iptables -A INPUT -p udp -m udp --sport 53 -j ACCEPT  
#进入本地服务器53端口的数据包进站通过
iptables -A INPUT -p udp -m udp --dport 53 -j ACCEPT  
#ICPM数据包可进入本地服务器
iptables -A INPUT -d 104.168.15.189 -p icmp -j ACCEPT 
#远端20到本地(ftp)
iptables -A INPUT -d 104.168.15.189 -p tcp -m tcp --dport 20 -j ACCEPT 
iptables -A INPUT -d 104.168.15.189 -p tcp -m tcp --dport 21 -j ACCEPT 

#对SSH的服务进入的数据包开启出站端口
iptables -A OUTPUT -s 104.168.15.189 -p tcp -m tcp --sport 1024 -m state --state ESTABLISHED -j ACCEPT 
#对WEB的服务进入的数据包开启出站端口
iptables -A OUTPUT -s 104.168.15.189 -p tcp -m tcp --sport 80 -m state --state ESTABLISHED -j ACCEPT 
#允许本地环回数据
iptables -A OUTPUT -s 127.0.0.1 -d 127.0.0.1 -j ACCEPT 
#从本地53端口出站的数据包出站通过
iptables -A OUTPUT -p udp -m udp --sport 53 -j ACCEPT  
#去往远程DNS服务器53端口的数据包出站通过
iptables -A OUTPUT -p udp -m udp --dport 53 -j ACCEPT  
#对对方ICMP数据包回应(ping命令回应数据包)
iptables -A OUTPUT -s 104.168.15.189 -p icmp -j ACCEPT 
#本地到远端20的请求(ftp)
iptables -A OUTPUT -s 104.168.15.189 -p tcp --dport 20 -j ACCEPT 
iptables -A OUTPUT -s 104.168.15.189 -p tcp --dport 21 -j ACCEPT 

#保存配置信息
#service iptables save 
#开启防火墙服务
#service iptables start 
