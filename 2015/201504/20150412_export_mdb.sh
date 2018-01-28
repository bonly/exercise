#!/bin/bash
#for i in $( mdb-tables $DB ); do echo $i ; mdb-export -D "%Y-%m-%d %H:%M:%S" -H -I mysql $DB $i > sql/$i.sql; done

DB=LookinBody30
mkdir -p $DB
for i in $( mdb-tables $DB.mdb ); do 
echo process $i ; 
mdb-schema --drop-table -T $i LookinBody30.mdb mysql | sqlite3 ib.db
mdb-export -D "%Y-%m-%d %H:%M:%S" -I mysql -H $DB.mdb $i > $DB/$i.sql; 
cat $DB/$i.sql | sqlite3 ib.db ;
done


DB=LBM
mkdir -p $DB
for i in $( mdb-tables $DB.mdb ); do 
echo process $i ; 
mdb-schema --drop-table -T $i LookinBody30.mdb mysql | sqlite3 lbm.db
mdb-export -D "%Y-%m-%d %H:%M:%S" -I mysql -H $DB.mdb $i > $DB/$i.sql; 
cat $DB/$i.sql | sqlite3 lbm.db ;
done