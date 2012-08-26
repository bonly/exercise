#!/bin/sh /etc/rc.common
# 2008 weedy2887@gmail.com

START=42

prep() {
        lanif="$(uci -P /var/state get network.lan.ifname)"
        wanif="$(uci -P /var/state get network.wan.ifname)"
        # retrieve the public IPv4 address
        ipv4=$(ifconfig $wanif | grep 'inet addr' | awk '{print $2}' | cut -d':' -f 2)
        # get the IPv6 prefix from the IPv4 address
        ipv6prefix=$(echo $ipv4 | awk -F. '{ printf "2002:%02x%02x:%02x%02x", $1, $2, $3, $4 }')
        # the local subnet (any 4 digit hex number)
        ipv6subnet=0666

        # The 6to4 relay: here are a few, use the anycast address when possible
        # For others see http://www.kfu.com/~nsayer/6to4/#list or google
        #relay6to4=144.232.8.254 #jakllsch@freenode told me
        #relay6to4=66.117.34.140 # Old Cox?
        # anycast:
        relay6to4=192.88.99.1
        # uni-leipzig.de:
        #relay6to4=139.18.25.33
        # 6to4.ipv6.bt.com
        #relay6to4=194.73.82.244
        # microsoft
        #relay6to4=131.107.33.60
        # japan kddilab.6to4.jp
        #relay6to4=192.26.91.178
}

start() {
        prep

        echo "Creating tunnel interface..."
        ip tunnel add tun6to4 mode sit ttl 64 remote any local $ipv4
        echo "Setting tunnel interface up..."
        ip link set dev tun6to4 up
        echo "Assigning ${ipv6prefix}::1/16 address to tunnel interface..."
        ip -6 addr add ${ipv6prefix}::1/16 dev tun6to4
        echo "Adding route to IPv6 internet on tunnel interface via relay..."
        ip -6 route add 2000::/3 via ::${relay6to4} dev tun6to4 metric 1
        echo "Assigning ${ipv6prefix}:${ipv6subnet}::1/64 address to $lanif (local lan interface)..."
        ip -6 addr add ${ipv6prefix}:${ipv6subnet}::1/64 dev $lanif
        echo "Done."
}

stop() {
        prep

        echo "Removing $lanif (internal lan) interface IPv6 address..."
        ip -6 addr del ${ipv6prefix}:${ipv6subnet}::1/64 dev $lanif
        echo "Removing routes to 6to4 tunnel interface..."
        ip -6 route flush dev tun6to4
        echo "Setting tunnel interface down..."
        ip link set dev tun6to4 down
        echo "Removing tunnel interface..."
        ip tunnel del tun6to4
        echo "Done."
}