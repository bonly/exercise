. /etc/functions.sh
NAME=ipv6
COMMAND=/usr/sbin/ip
[ "$ACTION" = "ifup" -a "$INTERFACE" = "wan" ] && {
        [ -x $COMMAND ] && {
                IFNAME=$(nvram get ${INTERFACE}_ifname)
                IPV4=$(ip addr show $IFNAME | grep inet | cut -f6 -d' ')
                IPV6PREFIX=$(echo $IPV4 | awk -F. '{ printf "2002:%02x%02x:%02x%02x", $1, $2, $3, $4 }')
                ip tunnel add tun6to4 mode sit ttl 64 remote any local $IPV4
                ip link set dev tun6to4 up
                ip -6 addr add ${IPV6PREFIX}::1/16 dev tun6to4
                ip -6 route add 2000::/3 via ::192.88.99.1 dev tun6to4 metric 1
                ip -6 addr add ${IPV6PREFIX}:5678::1/64 dev br0
        } &
}
[ "$ACTION" = "ifdown" -a "$INTERFACE" = "wan" ] && {
        [ -x $COMMAND ] && {
                IFNAME=$(nvram get ${INTERFACE}_ifname)
                IPV4=$(ip addr show $IFNAME | grep inet | cut -f6 -d' ')
                IPV6PREFIX=$(echo $IPV4 | awk -F. '{ printf "2002:%02x%02x:%02x%02x", $1, $2, $3, $4 }')
                ip -6 addr del ${IPV6PREFIX}:5678::1/64 dev br0
                ip -6 route flush dev tun6to4
                ip link set dev tun6to4 down
                ip tunnel del tun6to4
        } &
}