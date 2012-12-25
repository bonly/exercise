#!/usr/bin/python
#-*-coding:utf-8-*-
import MySQLdb
import subprocess

global conn, cursor
newdb = "paladin";
olddb = "paladin_test";

def ConnectDB():
   global conn, cursor
   conn = MySQLdb.connect(host='183.60.126.26', user='bonly', passwd='')
   cursor = conn.cursor();
    
def CloseAll():
   cursor.close()

def UseDb():
   conn.select_db(newdb);

def CreateDB():
   cursor.execute("""create database if not exists %s""" % newdb)
   
def CopyTable(table_name, nodata):
   if nodata == True:
      cursor.execute("""create table %s select * 
        from %s.%s where 1=2 """ % (table_name, olddb, table_name));
   else :
        cursor.execute("""create table %s select * 
        from %s.%s  """ % (table_name, olddb, table_name));    

def DelData():
   subprocess.call("sed -i '/INSERT INTO `player`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `player_bag`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `charge_record`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `enemy_record`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `equipment`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `gambling_message`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `item`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `mall_log`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `messages`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `paladin`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `player_friend`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `player_friend_request`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `player`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `player_daily_mission`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `player_finished_activity`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `pvp_state`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `top_board`/d' paladin_test.sql", shell=True);
   subprocess.call("sed -i '/INSERT INTO `wager_info`/d' paladin_test.sql", shell=True);

if __name__ == "__main__":
   subprocess.call("""mysqldump -h 183.60.126.26 -ubonly \
                       --default-character-set=utf8  \
                       --opt --extended-insert=false  \
                       --triggers -R --hex-blob  \
                       -x paladin_test > paladin_test.sql""", shell=True);
   #--no-data 

   ConnectDB();
   CreateDB();
   UseDb();
   
   DelData();

   subprocess.call("""mysql -h 183.60.126.26 -ubonly \
                       paladin < paladin_test.sql """, shell=True);



   """CopyTable("activity_list", False);
   CopyTable("activity_template", False);
   CopyTable("charge_record", True);
   CopyTable("copy", False);
   CopyTable("copy_chapter",False);
   CopyTable("daily_mission_reward",False);
   CopyTable("daily_mission_template",False);
   CopyTable("drop_group",False);
   CopyTable("drop_probability",False);
   CopyTable("enemy_record",True);
   CopyTable("equipment",True);"""

   """CloseAll();"""
   
