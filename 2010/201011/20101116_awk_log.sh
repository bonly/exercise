awk '$5 ~ /\"port\"/ {print $5 >> $1}' gsrv.sxh-009.localdomain.wlmz.log.INFO.2013020*
