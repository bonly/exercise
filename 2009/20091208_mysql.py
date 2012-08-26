#!/usr/bin/python
#-*-coding:utf-8-*-
import MySQLdb

global conn,cursor

def ConnectDB():
   global conn,cursor
   conn = MySQLdb.connect(host='127.0.0.1',user='root',passwd='')
   cursor = conn.cursor()
   
def CreateDB():
   cursor.execute("""create database if not exists testdb""")
   
def CreateTab():   
   conn.select_db("testdb");
   cursor.execute("""create table if not exists testct(
                     id int,
                     b  char(50),
                     index(id)) ENGINE=MyISAM""")
 
def InsertData():
   value = [1,"db1"];
   cursor.execute("insert into testct values(%s,%s)",value);
     
def InsertMulData():
   values = []
   for i in range(10000):
      values.append((i,"Hello mysqldb, i am recorder "+str(i)))
   cursor.executemany("""insert into testct value(%s,%s)""",values);
   
def CloseAll():
   cursor.close()
   
def SelectData():
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
   
if __name__=="__main__":
   ConnectDB()
   CreateDB()
   CreateTab()
   InsertMulData()
   #SelectData()
   CloseAll()
   