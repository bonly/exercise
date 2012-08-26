#!/usr/bin/python
#-*-coding:utf-8-*-
import MySQLdb
from combinatorics.all_pairs2 import all_pairs2
global conn, cursor

def ConnectDB():
   global conn, cursor
   conn = MySQLdb.connect(host='127.0.0.1', user='root', passwd='')
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
   value = [1, "db1"];
   cursor.execute("insert into testct values(%s,%s)", value);
     
def InsertMulData_testct():
   values = []
   for i in range(10000):
      values.append((i, "Hello mysqldb, i am recorder " + str(i)))
   cursor.executemany("""insert into testct value(%s,%s)""", values);
   
def CloseAll():
   cursor.close()
   
def SelectData_testct():
   count = cursor.execute('select * from testct')  
   print '总共有 %s 条记录' % count 
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
                     cl_7 bigint(22),
                     cl_8 bigint(22),
                     cl_9 bigint(22),
                     cl_10 date,
                     cl_11 date,
                     cl_12 date,
                     index(id)) ENGINE=MyISAM""")
         
def InsertData_testmk():
   begin_id = 0
   count = 500000
   phone = []
   phonen = []
   for i in range(20):
      phone.append(13719360007+i)
   for i in range(20):
      phonen.append(13848004162+i)   
   field = [["Windows", "Linux", "iOS", "Android", "DOS"],
            ["98", "NT", "2000", "XP", "win7", "win8", "2.4", "2.6", "Debian", "Gennto",
             "RedHat", "FeDroa", "Ubuntu", "Mint", "4.3", "1.5", "1.6", "2.0", "2.1", "2.2", "2.3", "2.3.3"],
            ['mikey', 'miny', 'mike', 'jeson', 'dance', 'rose', 'yellow', 'red', 'green', 'blue'],
            ['night', 'afternoon', 'morning'],
            ['20070103', '20081211', '20100308', '20030708', '20040802'],
            phone,
            phonen,
            ['34133444211', '4039394016', '8473410002', '39703874439'],
            ['34113444211', '4019394016', '8473210002', '39103874439'],
            ['1997-05-03', '2009-01-28', '2008-08-21', '2011-02-01', '2000-09-12', '2002-03-06'],
            ['2003-11-23', '2003-11-22', '2003-12-31', '2005-09-12', '2006-07-01', '2006-08-03', '2007-07-01'],
            ['2003-10-23', '2002-11-22', '2003-12-31', '2006-09-12', '2003-07-01', '2001-08-03', '2007-03-01', '2010-11-26']
           ]
   pair = all_pairs2(field)
   for i, v in enumerate(pair):
      if i >= count:
         return 
      cursor.execute("insert into testmk value(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)",
                     [i+begin_id, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11]])
   m = i
   while m % (count + 1) != 0 :
      if m >= count:
         break
      pair = all_pairs2(field)
      for i, v in enumerate(pair):
         m = m + 1
         if m > count :
            break
         cursor.execute("insert into testmk value(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)",
                        [m+begin_id, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11]])
         

   
def CreateTab_testmgi(ind):
   conn.select_db("testdb")
   strsql = "create table if not exists testmg" + ind
   strsql = strsql + """(
                     id int,
                     cl_1 char(50),
                     cl_2 char(50),
                     cl_3 char(50),
                     cl_4 char(50),
                     cl_5 char(50),
                     cl_6 char(50),
                     cl_7 bigint(22),
                     cl_8 bigint(22),
                     cl_9 bigint(22),
                     cl_10 date,
                     cl_11 date,
                     cl_12 date,
                     index(id)) ENGINE=MyISAM"""   
   cursor.execute(strsql)
   
def CreateTab_testmgall():
   conn.select_db("testdb")
   cursor.execute("""create table if not exists testmg (
                     id int,
                     cl_1 char(50),
                     cl_2 char(50),
                     cl_3 char(50),
                     cl_4 char(50),
                     cl_5 char(50),
                     cl_6 char(50),
                     cl_7 bigint(22),
                     cl_8 bigint(22),
                     cl_9 bigint(22),
                     cl_10 date,
                     cl_11 date,
                     cl_12 date,
                     index(id)) 
                     ENGINE=MERGE UNION=(testmg1,testmg2,testmg3,testmg4,testmg5) INSERT_METHOD=LAST"""
   )

def InsertData_testmgi(num):
   count = 100000
   field = [["Windows", "Linux", "iOS", "Android", "DOS"],
            ["98", "NT", "2000", "XP", "win7", "win8", "2.4", "2.6", "Debian", "Gennto",
             "RedHat", "FeDroa", "Ubuntu", "Mint", "4.3", "1.5", "1.6", "2.0", "2.1", "2.2", "2.3", "2.3.3"],
            ['mikey', 'miny', 'mike', 'jeson', 'dance', 'rose', 'yellow', 'red', 'green', 'blue'],
            ['night', 'afternoon', 'morning'],
            ['20070103', '20081211', '20100308', '20030708', '20040802'],
            ['13719360007', '13503950007', '13855550007', '15360534220'],
            ['123456789012', '13413101118', '13413101119', '13413109887', '13498331107'],
            ['34113444213', '4019394013', '8473210003', '39103874433', '1139387588'],
            ['34113444211', '4019394016', '8473210002', '39103874439'],
            ['1997-05-03', '2009-01-28', '2008-08-21', '2011-02-01', '2000-09-12', '2002-03-06'],
            ['2003-11-23', '2003-11-22', '2003-12-31', '2005-09-12', '2006-07-01', '2006-08-03', '2007-07-01'],
            ['2003-10-23', '2002-11-22', '2003-12-31', '2006-09-12', '2003-07-01', '2001-08-03', '2007-03-01', '2010-11-26']
           ]
   pair = all_pairs2(field)
   for i, v in enumerate(pair):
      sql = "insert into testmg" + str(num) + " value(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)"
      cursor.execute(sql, [i, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11]])
   m = i
   while m % (count + 1) != 0 :
      pair = all_pairs2(field)
      for i, v in enumerate(pair):
         cursor.execute(sql, [m, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11]])
         m = m + 1
         if m > count :
            break
         
def InsertData_testmgi_v2(num):
   count = 100000
   field = [["Windows", "Linux", "iOS", "Android", "DOS"],
            ["98", "NT", "2000", "XP", "win7", "win8", "2.4", "2.6", "Debian", "Gennto",
             "RedHat", "FeDroa", "Ubuntu", "Mint", "4.3", "1.5", "1.6", "2.0", "2.1", "2.2", "2.3", "2.3.3"],
            ['mikey', 'miny', 'mike', 'jeson', 'dance', 'rose', 'yellow', 'red', 'green', 'blue'],
            ['night', 'afternoon', 'morning'],
            ['20070103', '20081211', '20100308', '20030708', '20040802'],
            ['13719360007', '13503950007', '13855550007', '15360534220'],
            ['123456789012', '13413101118', '13413101119', '13413109887', '13498331107'],
            ['34113444213', '4019394013', '8473210003', '39103874433', '1139387588'],
            ['34113444211', '4019394016', '8473210002', '39103874439'],
            ['1997-05-03', '2009-01-28', '2008-08-21', '2011-02-01', '2000-09-12', '2002-03-06'],
            ['2003-11-23', '2003-11-22', '2003-12-31', '2005-09-12', '2006-07-01', '2006-08-03', '2007-07-01'],
            ['2003-10-23', '2002-11-22', '2003-12-31', '2006-09-12', '2003-07-01', '2001-08-03', '2007-03-01', '2010-11-26']
           ]
   pair = all_pairs2(field)
   sql = "insert into testmg" + str(num) + " value(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)"
   values = []
   for i, v in enumerate(pair):
      values.append(([i, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11]]))
   m = i
   while m % (count + 1) != 0 :
      pair = all_pairs2(field)
      for i, v in enumerate(pair):
         values.append(([m, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11]]))
         m = m + 1
         if m > count :
            break
   #生成所有记录后再插入表中
   cursor.executemany(sql, values)

def InsertData_by_testmk():
   count = 100000
   sql = []
   for i in [0, 1, 2, 3, 4]:
      sql = "insert into testmg" + str(i+1) + " (select * from testmk limit 100000 offset " + str(i*count) + ")"
      cursor.execute(sql);
   
def InsertData_testmgi_v3(num):
   count = 100000
   field = [["Windows", "Linux", "iOS", "Android", "DOS"],
            ["98", "NT", "2000", "XP", "win7", "win8", "2.4", "2.6", "Debian", "Gennto",
             "RedHat", "FeDroa", "Ubuntu", "Mint", "4.3", "1.5", "1.6", "2.0", "2.1", "2.2", "2.3", "2.3.3"],
            ['mikey', 'miny', 'mike', 'jeson', 'dance', 'rose', 'yellow', 'red', 'green', 'blue'],
            ['night', 'afternoon', 'morning'],
            ['20070103', '20081211', '20100308', '20030708', '20040802'],
            ['13719360007', '13503950007', '13855550007', '15360534220'],
            ['123456789012', '13413101118', '13413101119', '13413109887', '13498331107'],
            ['34113444213', '4019394013', '8473210003', '39103874433', '1139387588'],
            ['34113444211', '4019394016', '8473210002', '39103874439'],
            ['1997-05-03', '2009-01-28', '2008-08-21', '2011-02-01', '2000-09-12', '2002-03-06'],
            ['2003-11-23', '2003-11-22', '2003-12-31', '2005-09-12', '2006-07-01', '2006-08-03', '2007-07-01'],
            ['2003-10-23', '2002-11-22', '2003-12-31', '2006-09-12', '2003-07-01', '2001-08-03', '2007-03-01', '2010-11-26']
           ]
   pair = all_pairs2(field)
   sql = []
   for i in range(5):
      sql.append("insert into testmg" + str(i + 1) + " value(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)")
     
   values = []
   for i, v in enumerate(pair):
      if i >= count:
         break
      values.append(([i, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11]]))
   
   m = i
   once = i #记录一个正交表的大小
   
   #先把数据插入到已有的表中
   for i in range(5):
      cursor.executemany(sql[i], values)
      
   #删除列表中所有数据,但不删除变量,del values会连变量一起删除
   del values[:]
   
   while m % (count + 1) != 0 :
      if m >= count:
         break
      pair = all_pairs2(field)
      for i, v in enumerate(pair):
         if m < count:
            values.append(([m, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11]]))
         m = m + 1
         
         if m % (once * 5) == 0: #每当有一个正交表大小的数据,则插入数据一次
            for q in range(5):
               cursor.executemany(sql[q], values)
            del values[:] #清空数据
            
         if m >= count :
            for o in range(5): #插入剩余的数据
               cursor.executemany(sql[o], values)
            break

def CreateTab_testone():   
   conn.select_db("testdb");
   cursor.execute("""create table if not exists testone(
                     id int,
                     cl_1 char(50),
                     cl_2 char(50),
                     cl_3 char(50),
                     cl_4 char(50),
                     cl_5 char(50),
                     cl_6 char(50),
                     cl_7 bigint(22),
                     cl_8 bigint(22),
                     cl_9 bigint(22),
                     cl_10 date,
                     cl_11 date,
                     cl_12 date,
                     index(id)) ENGINE=MyISAM""")

def InsertData_testone():
   count = 10000
   begin_id = 500001
   phone = []
   phonen = []
   for i in range(20):
      phone.append(13719360007+i)
   for i in range(20):
      phonen.append(13848004162+i)
      
   field = [["Windows", "Linux", "iOS", "Android", "DOS"],
            ["98", "NT", "2000", "XP", "win7", "win8", "2.4", "2.6", "Debian", "Gennto",
             "RedHat", "FeDroa", "Ubuntu", "Mint", "4.3", "1.5", "1.6", "2.0", "2.1", "2.2", "2.3", "2.3.3"],
            ['mikey', 'miny', 'mike', 'jeson', 'dance', 'rose', 'yellow', 'red', 'green', 'blue'],
            ['night', 'afternoon', 'morning'],
            ['20070103', '20081211', '20100308', '20030708', '20040802'],
            phone,
            phonen,
            ['34113444213', '4019394013', '8473210003', '39103874433', '1139387588'],
            ['34113444211', '4019394016', '8473210002', '39103874439'],
            ['1997-05-03', '2009-01-28', '2008-08-21', '2011-02-01', '2000-09-12', '2002-03-06'],
            ['2003-11-23', '2003-11-22', '2003-12-31', '2005-09-12', '2006-07-01', '2006-08-03', '2007-07-01'],
            ['2003-10-23', '2002-11-22', '2003-12-31', '2006-09-12', '2003-07-01', '2001-08-03', '2007-03-01', '2010-11-26']
           ]
   pair = all_pairs2(field)
   for i, v in enumerate(pair):
      #print "insert into testmk value(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)"% \
      #         (i,v[0],v[1],v[2],v[3],v[4],v[5],v[6],v[7],v[8],v[9],v[10],v[11])
      if i >=count :
         return
      cursor.execute("insert into testone value(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)",
                     [begin_id+i, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11]])
   m = i
   while m % (count + 1) != 0 :
      if m >= count :
         return
      pair = all_pairs2(field)
      for i, v in enumerate(pair):
         m = m + 1
         if m >= count :
            return
         cursor.execute("insert into testone value(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)",
                        [begin_id+m, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11]])
         

            
if __name__ == "__main__":
   ConnectDB()
   CreateDB()
   
   #单表操作
   CreateTab_testmk()
   InsertData_testmk()
   #SelectData()
   
   #多表联合操作
   #for i in(1, 2, 3, 4, 5):
   #   CreateTab_testmgi(str(i))
   
   #主mege表建表
   #CreateTab_testmgall()
   
   #多表插入数据
   #for i in (1, 2, 3, 4, 5):
      ##InsertData_testmgi(i)
      ##InsertData_testmgi_v2(i)
      #InsertData_testmgi_v3(i)
   
   #用insert插入数据
   #InsertData_by_testmk()
   
   #CreateTab_testone()
   #InsertData_testone()
   
   CloseAll()
   
