#!/usr/bin/env monkeyrunner
#-*- coding=utf-8 -*-
from com.android.monkeyrunner import MonkeyRunner, MonkeyDevice
from com.android.monkeyrunner import MonkeyImage
from com.android.monkeyrunner.easy import EasyMonkeyDevice
from com.android.monkeyrunner.easy import By

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
  device.drag((244,465),(245,466),0.5,50);
  
  MonkeyRunner.sleep(1)
  print 'click 注册'
  device.drag((242,556),(245,558),0.5,50);
  MonkeyRunner.sleep(1)
  
  MonkeyRunner.sleep(1)
  print 'click 名字'
  device.drag((258,104),(259,105),0.5,50);
  
  name=random.choice(File)[:-1]
  #name=name+'_'
  #name=name+random.choice(File)[:-1]
    
  print '输入名字 ' + name
  MonkeyRunner.sleep(1)
  device.type(name);
  MonkeyRunner.sleep(1)
  #device.press('KEYCODE_ENTER', MonkeyDevice.DOWN_AND_UP);
  device.drag((120,520),(122,522),0.5,50);
    
  print 'click 密码'
  MonkeyRunner.sleep(1)
  device.drag((270,130),(273,130),1,50);
  MonkeyRunner.sleep(1)
  device.type('mypass');
  #device.press('KEYCODE_ENTER', MonkeyDevice.DOWN_AND_UP);
  device.drag((120,520),(122,522),0.5,50);
  
  print 'click 提交'
  MonkeyRunner.sleep(1)
  device.drag((245,507),(240,507),0.5,10);
  MonkeyRunner.sleep(1)
  
  print '选区'
  MonkeyRunner.sleep(2)
  device.drag((233,633),(236,635),0.5,10);
  
  print '选人'
  MonkeyRunner.sleep(2)
  device.drag((235,631),(238,635),0.5,10);
  
  MonkeyRunner.sleep(5);
  device.shell('am force-stop com.hytc.sg')

while 1:
  print 'begin time: '+  strftime("%Y-%m-%d %H:%M:%S", gmtime())
  reg();
  print 'end time: ' +  strftime("%Y-%m-%d %H:%M:%S", gmtime())

#adb shell 'am force-stop com.hytc.sg'
#http://gm.hanfenggame.com/hytc/qd_login.jsp
#2732  123123
