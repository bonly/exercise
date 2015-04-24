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

device = MonkeyRunner.waitForConnection(0,"HT049PL08237")
easy_device = EasyMonkeyDevice(device)
 
def reg():
  print 'begin'
  device.startActivity(component='com.hytc.xyol.android/.XYOL_Activity')
  
  MonkeyRunner.sleep(1)
  print 'click'
  device.drag((244,465),(245,466),0.5,50);
  
  MonkeyRunner.sleep(1)
  print 'click 注册'
  device.drag((244,465),(245,469),0.5,50);
  MonkeyRunner.sleep(1)
  device.drag((244,465),(245,469),0.5,50);
  
  MonkeyRunner.sleep(1)
  print 'click 名字'
  device.drag((308,506),(309,507),0.5,50);
  #device.touch(244,465,"DOWN_AND_UP");
  
  name=random.choice(File)[:-1]
  #name=name+' '
  ulen = 8
  if (15-len(name)>=8) :
    ulen = 8
  else:
    ulen = 15-len(name)
  name=name+str(uuid.uuid4())[:ulen]
    
  print '输入名字 ' + name
  device.type(name);
  #device.press('KEYCODE_ENTER', MonkeyDevice.DOWN_AND_UP);
  device.drag((177,468),(178,469),0.5,50);
  MonkeyRunner.sleep(1)
  
  passwd = str(uuid.uuid4())[:8]
  print 'click 密码:' + passwd
  device.drag((286,550),(290,555),1,50);
  MonkeyRunner.sleep(1)
  device.drag((286,550),(290,555),1,50);
  MonkeyRunner.sleep(1)
  device.type(passwd);
  #device.press('KEYCODE_ENTER', MonkeyDevice.DOWN_AND_UP);
  device.drag((177,468),(178,469),0.5,50);
  
  print 'click 提交'
  MonkeyRunner.sleep(1)
  device.drag((70,760),(75,765),1,10);
  MonkeyRunner.sleep(1)
  device.drag((70,760),(75,765),1,10);
  MonkeyRunner.sleep(1)
  
  print '选区'
  MonkeyRunner.sleep(3)
  device.drag((220,144),(222,147),1,10);
  MonkeyRunner.sleep(1)
  device.drag((220,144),(222,147),1,10);
    
  print '选人'
  MonkeyRunner.sleep(3)
  device.drag((62,770),(65,774),1,10);
  
  MonkeyRunner.sleep(7);
  device.shell('am force-stop com.hytc.xyol.android')

while 1:
  print 'begin time: '+  strftime("%Y-%m-%d %H:%M:%S", gmtime())
  reg();
  print 'end time: ' +  strftime("%Y-%m-%d %H:%M:%S", gmtime())

#http://gm.hanfenggame.com/hytc/qd_login.jsp
#70382  123123
