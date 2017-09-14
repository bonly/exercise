############## HEADER HERE
# Create new chain
iptables -t nat -N SHADOWSOCKS

# Replace 1080 to you own proxy port
{{range .}}iptables -t nat -A SHADOWSOCKS -d {{.}} -p tcp -j REDIRECT --to-ports 1080
{{end}}

# Apply the rules
iptables -t nat -A PREROUTING -p tcp -j SHADOWSOCKS