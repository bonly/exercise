#!/usr/bin/env monkeyrunner
#-*- coding=utf-8 -*-
from com.android.monkeyrunner import MonkeyRunner, MonkeyDevice
from com.android.monkeyrunner import MonkeyImage
from com.android.monkeyrunner.easy import EasyMonkeyDevice
from com.android.monkeyrunner.easy import By
import uuid
import random
from time import gmtime, strftime

Filename='./20110822_randomname.txt'
File=open(Filename,'r').readlines()

device = MonkeyRunner.waitForConnection(0, "12f7348f")
easy_device = EasyMonkeyDevice(device)
 
def reg():
  print 'begin'
  device.startActivity(component='com.hytc.sg/.LwsgActivity')
  
  MonkeyRunner.sleep(1)
  print 'click'
  device.drag((244,465),(245,466),0.5,10);
  
  MonkeyRunner.sleep(1)
  print 'click 注册'
  device.drag((242,556),(245,558),0.5,10);
  MonkeyRunner.sleep(1)
  
  MonkeyRunner.sleep(1)
  print 'click 名字'
  device.drag((258,104),(259,105),0.5,10);
  
  name=random.choice(File)[:-1]
  #name=name+'_'
  ulen = 8
  if (15-len(name)>=8) :
    ulen = 8
  else:
    ulen = 15-len(name)
  name=name+str(uuid.uuid4())[:ulen]
    
  print '输入名字 ' + name
  MonkeyRunner.sleep(1)
  device.type(name);
  MonkeyRunner.sleep(1)
  #device.press('KEYCODE_ENTER', MonkeyDevice.DOWN_AND_UP);
  device.drag((120,520),(122,522),0.5,10);
    
  passwd = str(uuid.uuid4())[:8]
  print 'click 密码:' + passwd
  MonkeyRunner.sleep(1)
  device.drag((270,130),(273,130),0.5,10);
  MonkeyRunner.sleep(1)
  device.type(passwd);
  #device.press('KEYCODE_ENTER', MonkeyDevice.DOWN_AND_UP);
  device.drag((120,520),(122,522),0.5,10);
  
  print 'click 提交'
  MonkeyRunner.sleep(2)
  device.drag((245,507),(240,507),0.5,10);
  MonkeyRunner.sleep(2)
  
  print '选区'
  MonkeyRunner.sleep(2)
  device.drag((233,633),(236,635),0.5,10);
  
  print '选人'
  MonkeyRunner.sleep(2)
  device.drag((235,631),(238,635),0.5,10);
  
  MonkeyRunner.sleep(2);
  print '村外'
  device.drag((239,503),(243,503),0.5,10); 
  MonkeyRunner.sleep(2);
  print '营救'
  device.drag((239,503),(243,503),0.5,10);  #营救
  MonkeyRunner.sleep(3);
  print '自动战斗'
  device.drag((222,247),(222,252),0.5,10);  #自动战斗
  MonkeyRunner.sleep(15);
  print '自动战斗'
  device.drag((222,247),(222,252),0.5,10);  #自动战斗
  MonkeyRunner.sleep(6);
  print '继续战斗'
  device.drag((376,660),(380,660),0.5,10);  #继续战斗
  MonkeyRunner.sleep(2);
  print '姑娘没事吧'
  device.drag((239,503),(243,503),0.5,10);  #姑娘没事吧
  MonkeyRunner.sleep(2);
  print '返回'
  device.drag((239,503),(243,503),0.5,10);  #返回
  MonkeyRunner.sleep(2);
  print '确定'
  device.drag((230,389),(233,389),0.5,10);  #确定  
  MonkeyRunner.sleep(2);
  print '确定'
  device.drag((214,471),(220,471),0.5,10);  #确定
  MonkeyRunner.sleep(2);
  print '对话'
  device.drag((368,91),(370,91),0.5,10);  #对话
  MonkeyRunner.sleep(2);
  print '前往营寨'
  device.drag((239,503),(243,503),0.5,10);  #前往营寨  
  MonkeyRunner.sleep(2.5);
  print '确定'
  device.drag((230,400),(233,400),0.5,10);  #确定  
  MonkeyRunner.sleep(2);  
  print '山贼营寨'
  device.drag((313,80),(322,80),0.5,10);  #第一行地点  
  MonkeyRunner.sleep(2);  
  print '山贼群'
  device.drag((433,110),(440,110),0.5,10);  #第二行地点  
  MonkeyRunner.sleep(2);    
  print '自动战斗'
  device.drag((222,247),(222,252),0.5,10);  #自动战斗
  MonkeyRunner.sleep(4);  
  print '自动战斗'
  device.drag((222,247),(222,252),0.5,10);  #自动战斗
  MonkeyRunner.sleep(4);    
  print '继续战斗'
  device.drag((376,660),(380,660),0.5,10);  #继续战斗
  MonkeyRunner.sleep(2);  

  MonkeyRunner.sleep(2);    
  print '自动战斗'
  device.drag((222,247),(222,252),0.5,10);  #自动战斗
  MonkeyRunner.sleep(4);  
  print '自动战斗'
  device.drag((222,247),(222,252),0.5,10);  #自动战斗
  MonkeyRunner.sleep(4);    
  print '继续战斗'
  device.drag((376,660),(380,660),0.5,10);  #继续战斗
  MonkeyRunner.sleep(2);    
  
  device.drag((216,620),(220,620),0.5,10);  #中间键返回
  MonkeyRunner.sleep(2);   

  print '确定'
  device.drag((214,471),(220,471),0.5,10);  #确定
  MonkeyRunner.sleep(2);  
  
  device.shell('am force-stop com.hytc.sg')

while 1:
  print 'begin time: '+  strftime("%Y-%m-%d %H:%M:%S", gmtime())
  reg();
  print 'end time: ' +  strftime("%Y-%m-%d %H:%M:%S", gmtime())

#adb shell 'am force-stop com.hytc.sg'
#http://gm.hanfenggame.com/hytc/qd_login.jsp
#2732  123123
