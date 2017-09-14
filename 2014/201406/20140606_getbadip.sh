#!/bin/sh

POSIONEDDOMAIN="www.twitter.com twitter.com www.facebook.com facebook.com www.youtube.com youtube.com encrypted.google.com plus.google.com www.appspot.com appspot.com www.openvpn.net openvpn.net forums.openvpn.net svn.openvpn.net shell.cjb.net"
LOOPTIMES=1
#RULEFILENAME=/var/g.firewall.user
RULEFILENAME=/tmp/g.firewall.user
DUTY_DNS="210.21.4.130 221.5.88.88"

# wait tail has internet connection
#while ! ping -W 1 -c 1 8.8.8.8 >&/dev/null; do sleep 30; done

badip=""

#复制出第二个POSIONEDDOMAIN
querydomain=""
matchregex="^${POSIONEDDOMAIN//\ /|^}"
for i in $(seq $LOOPTIMES) ; do
        querydomain="$querydomain $POSIONEDDOMAIN"
done

#建表
echo "iptables -N ill_ip" > $RULEFILENAME.tmp
echo "iptables -I INPUT -p udp --sport 53 -j ill_ip" >> $RULEFILENAME.tmp 
echo "iptables -I FORWARD -p udp --sport 53 -j ill_ip" >> $RULEFILENAME.tmp

for DOMAIN in $POSIONEDDOMAIN ; do
        for bad_dns in $DUTY_DNS ; do
                #for IP in $(dig +time=1 +tries=1 +retry=0 @$DOMAIN $querydomain | grep -E "$matchregex" | grep -o -E "([0-9]+\.){3}[0-9]+") ; do
                for IP in $(dig +time=1 +tries=1 +retry=0 @$bad_dns $querydomain | grep -E "$matchregex" | grep -o -E "([0-9]+\.){3}[0-9]+") ; do
                        if [ -z "$(echo $badip | grep $IP)" ] ; then
                                badip="$badip   $IP"
                        fi
                done
        done
done

for IP in $badip ; do
        hexip=$(printf '%02X ' ${IP//./ }; echo)
        echo "#block ${IP}" >> $RULEFILENAME.tmp 
        echo "iptables -I ill_ip -p udp --sport 53 -m string --algo bm --hex-string \"|$hexip|\" --from 60 --to 180  -j DROP" >> $RULEFILENAME.tmp 
done

#丢掉不包含任何查询结果的包
echo "iptables -I ill_ip -p udp --sport 53 -m u32 --u32 \"4 & 0x1FFF = 0 && 0 >> 22 & 0x3C @ 8 & 0x8000 = 0x8000 && 0 >> 22 & 0x3C @ 14 = 0\" -j DROP" >> $RULEFILENAME.tmp
#丢掉Answer、Authority和Additional均为0的应答
echo "iptables -I ill_ip -p udp --sport 53 -m string --algo bm --hex-string \"|81 80 00 01 00 00 00 00 00 00|\" --from 30 --to 40 -j DROP" >> $RULEFILENAME.tmp

if [[ -s $RULEFILENAME ]] ; then
        grep -Fvf $RULEFILENAME $RULEFILENAME.tmp > $RULEFILENAME.action
        cat $RULEFILENAME.action >> $RULEFILENAME
else
        cp $RULEFILENAME.tmp $RULEFILENAME
        cp $RULEFILENAME.tmp $RULEFILENAME.action
fi

#. $RULEFILENAME.action
rm $RULEFILENAME.tmp
#rm $RULEFILENAME.action