#!/bin/sh

POSIONEDDOMAIN="www.twitter.com twitter.com www.facebook.com facebook.com www.youtube.com youtube.com encrypted.google.com plus.google.com www.appspot.com appspot.com www.openvpn.net openvpn.net forums.openvpn.net svn.openvpn.net shell.cjb.net"
LOOPTIMES=1
RULEFILENAME=/var/g.firewall.user

# wait tail has internet connection
while ! ping -W 1 -c 1 8.8.8.8 >&/dev/null; do sleep 30; done

badip=""

querydomain=""
matchregex="^${POSIONEDDOMAIN//\ /|^}"
for i in $(seq $LOOPTIMES) ; do
        querydomain="$querydomain $POSIONEDDOMAIN"
done

for DOMAIN in $POSIONEDDOMAIN ; do
        for IP in $(dig +time=1 +tries=1 +retry=0 @$DOMAIN $querydomain | grep -E "$matchregex" | grep -o -E "([0-9]+\.){3}[0-9]+") ; do
                if [ -z "$(echo $badip | grep $IP)" ] ; then
                        badip="$badip   $IP"
                fi
        done
done

for IP in $badip ; do
        hexip=$(printf '%02X ' ${IP//./ }; echo)
        echo "iptables -I INPUT -p udp --sport 53 -m string --algo bm --hex-string \"|$hexip|\" --from 60 --to 180  -j DROP" >> $RULEFILENAME.tmp 
        echo "iptables -I FORWARD -p udp --sport 53 -m string --algo bm --hex-string \"|$hexip|\" --from 60 --to 180 -j DROP" >> $RULEFILENAME.tmp
done

echo "iptables -I INPUT -p udp --sport 53 -m u32 --u32 \"4 & 0x1FFF = 0 && 0 >> 22 & 0x3C @ 8 & 0x8000 = 0x8000 && 0 >> 22 & 0x3C @ 14 = 0\" -j DROP" >> $RULEFILENAME.tmp
echo "iptables -I FORWARD -p udp --sport 53 -m u32 --u32 \"4 & 0x1FFF = 0 && 0 >> 22 & 0x3C @ 8 & 0x8000 = 0x8000 && 0 >> 22 & 0x3C @ 14 = 0\" -j DROP" >> $RULEFILENAME.tmp

if [[ -s $RULEFILENAME ]] ; then
        grep -Fvf $RULEFILENAME $RULEFILENAME.tmp > $RULEFILENAME.action
        cat $RULEFILENAME.action >> $RULEFILENAME
else
        cp $RULEFILENAME.tmp $RULEFILENAME
        cp $RULEFILENAME.tmp $RULEFILENAME.action
fi

. $RULEFILENAME.action
rm $RULEFILENAME.tmp
rm $RULEFILENAME.action