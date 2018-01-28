#!/bin/bash
#for i in $( mdb-tables $DB ); do echo $i ; mdb-export -D "%Y-%m-%d %H:%M:%S" -H -I mysql $DB $i > sql/$i.sql; done

DB=LookinBody30.mdb
for i in $( mdb-tables $DB ); do 
echo $i ; 
mdb-export -D "%Y-%m-%d %H:%M:%S" -H -I sqlite $DB $i > LookinBody30/$i.sql; 
echo -e ".separator ","\n.import LookinBody30/$i.sql $i" | sqlite3 testdatabase.db
done


DB=LBM.mdb
for i in $( mdb-tables $DB ); do 
echo $i ; 
mdb-export -D "%Y-%m-%d %H:%M:%S" -H -I sqlite $DB $i > LBM/$i.sql; 
done