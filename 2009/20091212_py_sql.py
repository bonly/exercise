#!/usr/bin/python
#-*-coding:utf-8-*-
import MySQLdb
from combinatorics.all_pairs2 import all_pairs2
global conn,cursor

def ConnectDB():
   global conn,cursor
   conn = MySQLdb.connect(host='127.0.0.1',user='root',passwd='')
   cursor = conn.cursor()
   
def CreateDB():
   cursor.execute("""create database if not exists testdb""")
   
def CreateTab_testct():   
   conn.select_db("testdb");
   cursor.execute("""create table if not exists testct(
                     id int,
                     b  char(50),
                     index(id)) ENGINE=MyISAM""")
 
def InsertData_testct():
   value = [1,"db1"];
   cursor.execute("insert into testct values(%s,%s)",value);
     
def InsertMulData_testct():
   values = []
   for i in range(10000):
      values.append((i,"Hello mysqldb, i am recorder "+str(i)))
   cursor.executemany("""insert into testct value(%s,%s)""",values);
   
def CloseAll():
   cursor.close()
   
def SelectData_testct():
   count = cursor.execute('select * from testct')  
   print '总共有 %s 条记录'%count 
   #result = cursor.fetchone();
   #print result
   #result =  cursor.fetchmany(5)
   #for r in result:
   #   print r
   # 重置洲标位置, 0 ,为偏移量 mode=absolute|relative,默认为 relative
   #cursor.scroll(0,mode='absolute')
   results = cursor.fetchall()
   for r in results:
      print r
      
def CreateTab_testmk():   
   conn.select_db("testdb");
   cursor.execute("""create table if not exists testmk(
                     id int,
                     cl_1 char(50),
                     cl_2 char(50),
                     cl_3 char(50),
                     cl_4 char(50),
                     cl_5 char(50),
                     cl_6 char(50),
                     index(id)) ENGINE=MyISAM""")
         
def InsertData_testmk():
   field = [["Windows","Linux","iOS","Android","DOS"],
            ["98","NT","2000","XP","win7","win8","2.4","2.6","Debian","Gennto",
             "RedHat","FeDroa","Ubuntu","Mint","4.3","1.5","1.6","2.0","2.1","2.2","2.3","2.3.3"],
            ['mikey','miny','mike','jeson','dance','rose','yellow','red','green','blue'],
            ['night','afternoon','morning'],
            [20070103,20081211,20100308,20030708,20040802],
            [13719360007,13503950007,13855550007,15360534220]
           ]
   pair = all_pairs2(field)
   for i,v in enumerate(pair):
      cursor.execute("insert into testmk value(%s,%s,%s,%s,%s,%s,%s)",[i,v[0],v[1],v[2],v[3],v[4],v[5]])
   
if __name__=="__main__":
   ConnectDB()
   CreateDB()
   CreateTab_testmk()
   InsertData_testmk()
   #SelectData()
   CloseAll()
   