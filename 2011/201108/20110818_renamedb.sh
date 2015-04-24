#!/bin/bash

mysqlconn="mysql -u root -pmobi2us -S /var/lib/mysql/mysql.sock -h localhost"
olddb="td"
newdb="account"

#$mysqlconn -e “CREATE DATABASE $newdb”
params=$($mysqlconn -N -e "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE table_schema='$olddb'")

for name in $params; do
$mysqlconn -e "RENAME TABLE $olddb.$name to $newdb.$name";
done;

#$mysqlconn -e "DROP DATABASE $olddb"

