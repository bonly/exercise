#!/usr/bin/python
# -*- coding: utf-8 -*-
# Time-stamp: <2013-01-30 18:16:12 Wednesday by roowe>
import MySQLdb
import sys
import random
try:
    con = None
    con = MySQLdb.connect(host="127.0.0.1", db="paladin", user="mysql", passwd="4860e49a", charset="utf8")
    #con = MySQLdb.connect(host="183.60.126.26", db="paladin_test", user="bonly", passwd="1234", charset="utf8")
#con = MySQLdb.connect(host="127.0.0.1", db="paladin", user="bonly", passwd="iroowe", charset="utf8")
    curs = con.cursor();
    curs.execute("set names utf8");
    curs.execute("set wait_timeout=2880")
    curs.execute("set interactive_timeout=2880")
except MySQLdb.Error, e:
    print "Error %d: %s" % (e.args[0], e.args[1])
try:
    curs.execute("SET @rank=0;")
    curs.execute("INSERT INTO top_board(rank, rank_type, rank_date, player_id, name, popularity) SELECT @rank:=@rank+1 AS rank, 4, date(now()), player_id, name, popularity  FROM player WHERE vip!=0 AND level>=5 ORDER BY popularity DESC, login_time DESC;")    
    curs.execute("SET @rank=0;")
    curs.execute("INSERT INTO top_board(rank, rank_type, rank_date, player_id, name, gold, silver) SELECT @rank:=@rank+1 AS rank, 2, date(now()), player_id, name, gold, silver FROM player WHERE  vip!=0 AND level>=5 ORDER BY gold DESC,silver DESC, login_time DESC;")
    curs.execute("SET @rank=0;")
    curs.execute("INSERT INTO top_board(rank, rank_type, rank_date, player_id, name, power) SELECT @rank:=@rank+1 AS rank, 1, date(now()), player_id, name, power  FROM player WHERE  vip!=0 AND level>=5 ORDER BY power DESC, login_time DESC;")
    curs.execute("SET @rank=0;")
    curs.execute("INSERT INTO top_board(rank, rank_type, rank_date, player_id, name, level) SELECT @rank:=@rank+1 AS rank, 0, date(now()), player_id, name, level  FROM player WHERE vip!=0 AND level>=5 ORDER BY level DESC, exp DESC,login_time DESC;")
    #荣誉排行
    curs.execute("SET @rank=0;")
    curs.execute("INSERT INTO top_board(rank, rank_type, rank_date, player_id, name, level) SELECT @rank:=@rank+1 AS rank, 3, date(now()), player_id, name, level  FROM player WHERE vip!=0 AND level>=5 AND honour>0 ORDER BY honour DESC, honour_point DESC,login_time DESC;")
    #每日荣誉排行
    curs.execute("SET @rank=0;")
    curs.execute("INSERT INTO top_board(rank, rank_type, rank_date, player_id, name, level) SELECT @rank:=@rank+1 AS rank, 5, date(now()), player_id, name, level  FROM player WHERE vip!=0 AND level>=5 AND honour_day>0 ORDER BY honour_day DESC, login_time DESC;")
    con.commit()
except MySQLdb.Error, e:
    print "Error %d: %s" % (e.args[0], e.args[1])

try:
    curs.execute("UPDATE player SET is_get_top_reward_gold=0, is_get_top_reward_honour=0, is_get_top_reward_power=0, is_get_top_reward_level=0, is_get_top_reward_popularity=0, cd_times=0, slave_stat=0, slave_esc_time = NULL, reward_daily_mission_times=0, replace_gold_times=0, finish_caishen_times=0")
    con.commit()
except MySQLdb.Error, e:
    print "Error %d: %s" % (e.args[0], e.args[1])
try:
    curs.execute("DELETE FROM player_daily_mission")x
    curs.execute("DELETE FROM player_slave")
    curs.execute("DELETE FROM player_wager;")
    curs.execute("DELETE FROM wager_info;")
    curs.execute("DELETE FROM wager_fight_status;")
    curs.execute("DELETE FROM wager_board;") 
    curs.execute("UPDATE player_master SET master_id = 1001, fail_times=0;")
    curs.execute("UPDATE player_copy SET is_cool = 0, sub_id = 1;") 
    #每日荣誉清0
    curs.execute("UPDATE player set honour_day=0;")
    con.commit()
except MySQLdb.Error, e:
    print "Error %d: %s" % (e.args[0], e.args[1])
try:
    curs.execute("SELECT player_id, item_id, id, rest_days, name FROM player_festival_charge, item_template WHERE item_template_id=item_id AND rest_days > 0")
    rows = curs.fetchall()
    for row in rows:
        curs.execute("call pt_insert_item(%d, %d, 1, @ret);" % (int(row[0]), int(row[1])))
        curs.execute("UPDATE player_festival_charge SET rest_days=rest_days-1 WHERE id=%d" % (int(row[2]))) 
        curs.execute("SELECT player_id, name, vip, level, title, sex FROM player WHERE player_id=%d" % (int(row[0])))
        player_info = curs.fetchall()
        if player_info:
            player_info = player_info[0]
        else:
            continue
        msg = ""
        rest_days = int(row[3])
        if rest_days - 1 > 0:
            msg = u"大侠的%s已赠送,请前去背包查看,剩余%d天" % (row[4], rest_days - 1)
        else:
            msg = u"大侠的%s已赠送,请前去背包查看,礼包已赠送完" % (row[4])
        if player_info:
            curs.execute(u"INSERT INTO `messages` (`account_id`, `player_id`, `name`, `vip`, `level`, `title`, `sex`, `recv_player_id`, `recv_name`, `recv_vip`, `recv_level`, `recv_title`, `recv_sex`, `create_at`, `content`, `item_id`, `type`) VALUES (0, -99, '系统', 0, 0, 0, 0, %d, '%s', %d, %d, %d, %d, now(), '%s', 0, 5);" % (int(player_info[0]), player_info[1], int(player_info[2]), int(player_info[3]), int(player_info[4]), int(player_info[5]), msg))
    curs.execute("DELETE FROM player_festival_charge WHERE rest_days=0") 
    con.commit()
except MySQLdb.Error, e:
    print "Error %d: %s" % (e.args[0], e.args[1])

curs.close()
con.close()
