#!/usr/bin/env monkeyrunner
#-*- coding=utf-8 -*-
from com.android.monkeyrunner import MonkeyRunner, MonkeyDevice
from com.android.monkeyrunner import MonkeyImage
from com.android.monkeyrunner.easy import EasyMonkeyDevice
from com.android.monkeyrunner.easy import By
import uuid
import random
from time import gmtime, strftime
import time

global Filename;
global File;

def loop_job():
  while 1:  #重启后的循环
    global device ;
    global easy_device;
    device = MonkeyRunner.waitForConnection(0,"SH13TPL09945");
    easy_device = EasyMonkeyDevice(device);  
    time_beg = time.time();
    while 1:  #循环一次操作
      onetime();
      time_end = time.time();
      time_div = time_end - time_beg;
      print 'time div = ' + str(time_div);  
      if (time_div >= 7200):  #大于一定时间后,重启机器
        print 'need to reboot';
        #device.reboot();
        #time.sleep(60);
        break;
        
def onetime(): 
  print 'begin time: '+  strftime("%Y-%m-%d %H:%M:%S", time.localtime());
  reg();
  print 'end time: ' +  strftime("%Y-%m-%d %H:%M:%S", time.localtime());

 
def reg():
  print 'begin'
  device.startActivity(component='com.hytc.xyol.android/.XYOL_Activity')
  
  MonkeyRunner.sleep(1)
  print 'click'
  device.drag((244,465),(245,466),0.5,50);
  
  MonkeyRunner.sleep(1)
  print 'click 注册'
  device.drag((244,455),(245,469),0.5,50);
  MonkeyRunner.sleep(1)
  device.drag((244,455),(245,469),0.5,50);
  
  MonkeyRunner.sleep(1)
  print 'click 名字'
  device.drag((300,490),(300,507),0.5,50);
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
  device.drag((286,540),(290,555),1,50);
  MonkeyRunner.sleep(1)
  device.drag((286,540),(290,555),1,50);
  MonkeyRunner.sleep(1)
  device.type(passwd);
  #device.press('KEYCODE_ENTER', MonkeyDevice.DOWN_AND_UP);
  device.drag((177,468),(178,469),0.5,50);
  
  print 'click 提交'
  MonkeyRunner.sleep(1)
  device.drag((70,760),(75,765),1,10);
  MonkeyRunner.sleep(1)
  #device.drag((70,760),(75,765),1,10);
  #MonkeyRunner.sleep(1)
  
  print '选区'
  MonkeyRunner.sleep(1)
  device.drag((220,144),(222,147),1,10);
  
  print '选人'
  MonkeyRunner.sleep(3)
  device.drag((62,770),(65,774),1,10);
  MonkeyRunner.sleep(1)
  device.drag((62,770),(65,774),1,10);
   
  print '卷帘大将:原来是威灵显赫上仙驾到'
  MonkeyRunner.sleep(7);
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5);
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5);
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(2);
  device.drag((202,373),(212,373),1,10);   
  
  print '王母:今日的蟠桃...'
  MonkeyRunner.sleep(6);  
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5);
  device.drag((412,412),(412,420),1,10);
  MonkeyRunner.sleep(0.5); 
  device.drag((202,373),(212,373),1,10);  #现在就去
  
  name=random.choice(File)[:-1]
  #name=name+' '
  ulen = 8
  if (10-len(name)>=8) :
    ulen = 8
  else:
    ulen = 10-len(name)
  name=name+str(uuid.uuid4())[:ulen]
    
  print '输入名字 ' + name
  MonkeyRunner.sleep(3); 
  device.drag((67,150),(69,153),1,10); 
  MonkeyRunner.sleep(1); 
  device.type(name[:1]);
  MonkeyRunner.sleep(0.5); 
  device.type(str(uuid.uuid4())[:4]);
  MonkeyRunner.sleep(1); 
  device.drag((177,460),(178,469),0.5,50); #确定
  MonkeyRunner.sleep(1); 
  device.drag((61,345),(61,350),1,10);  #提交
  MonkeyRunner.sleep(3); 
  device.drag((412,412),(412,420),1,10);
  MonkeyRunner.sleep(0.5); 
  device.drag((412,412),(412,420),1,10);
  MonkeyRunner.sleep(0.5); 
  device.drag((412,412),(412,420),1,10);
  MonkeyRunner.sleep(0.5); 
  device.drag((412,412),(412,420),1,10);
  MonkeyRunner.sleep(0.5); 
  device.drag((412,412),(412,420),1,10);
  MonkeyRunner.sleep(0.5); 
  device.drag((412,412),(412,420),1,10);       
  MonkeyRunner.sleep(0.5); 
  device.drag((412,412),(412,420),1,10);          
  MonkeyRunner.sleep(0.5); 
  device.drag((202,373),(212,373),1,10);  #现在就去    
  
  print '嫦娥'
  MonkeyRunner.sleep(4); 
  device.drag((412,412),(412,420),1,10);  
  MonkeyRunner.sleep(0.5); 
  device.drag((412,412),(412,420),1,10);  
  MonkeyRunner.sleep(0.5); 
  device.drag((412,412),(412,420),1,10);  
  MonkeyRunner.sleep(0.5);     
  device.drag((202,373),(212,373),1,10);  #现在就去     
  
  print '天蓬元帅'
  MonkeyRunner.sleep(4); 
  device.drag((412,412),(412,420),1,10);  
  MonkeyRunner.sleep(0.5); 
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5); 
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5); 
  device.drag((412,412),(412,420),1,10);      
  MonkeyRunner.sleep(0.5);     
  device.drag((202,373),(212,373),1,10);  #战斗    
    
  print '战斗'
  MonkeyRunner.sleep(3);  
  device.drag((186,214),(186,220),1,10); #攻击
  MonkeyRunner.sleep(0.5);  
  device.drag((228,300),(228,310),1,10); #对象
  MonkeyRunner.sleep(1);  
  device.drag((186,214),(186,220),1,10); #攻击
  MonkeyRunner.sleep(0.5);  
  device.drag((228,300),(228,310),1,10); #对象
  MonkeyRunner.sleep(8);  
  print '回合二'
  device.drag((186,214),(186,220),1,10); #攻击
  MonkeyRunner.sleep(0.5);  
  device.drag((228,300),(228,310),1,10); #对象
  MonkeyRunner.sleep(1);  
  device.drag((186,214),(186,220),1,10); #攻击
  MonkeyRunner.sleep(0.5);  
  device.drag((228,300),(228,310),1,10); #对象
  MonkeyRunner.sleep(8);    
  device.drag((412,412),(412,420),1,10); #按任意
  
  print '可醒酒酒了'
  MonkeyRunner.sleep(3);    
  device.drag((412,412),(412,420),1,10);
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10);
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10);
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10);
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10);   #奖励100点经验
  MonkeyRunner.sleep(0.5);     
  print '去见王母'
  device.drag((202,373),(212,373),1,10);  #去见
  
  print '1-2升级'
  MonkeyRunner.sleep(5); 
  device.drag((227,412),(230,412),1,10); #升级确定
  
  print '王母:丢尽我的脸'
  MonkeyRunner.sleep(3);  
  device.drag((412,412),(412,420),1,10);   
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10);  
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10);   #奖励金疮药10个   
  print '寻找琉璃盏'
  MonkeyRunner.sleep(1);  
  device.drag((202,373),(212,373),1,10);  #去
  MonkeyRunner.sleep(7);  
  device.drag((227,412),(230,412),1,10); #找到确定
  MonkeyRunner.sleep(4);  
  
  print '回到王母,得有专人看管才是'
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10);  
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10);  
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10);  
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10);  #经验250
  MonkeyRunner.sleep(1);  
  device.drag((202,373),(212,373),1,10);  #去
  
  print '卷帘大将,神奇法力'
  MonkeyRunner.sleep(3);  
  device.drag((412,412),(412,420),1,10);          
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5);  
  device.drag((412,412),(412,420),1,10); 
  MonkeyRunner.sleep(0.5);  
  device.drag((202,373),(212,373),1,10);  #去        
  
  print '去蟠桃园'
  MonkeyRunner.sleep(4);  
  print '2-3升级'
  device.drag((227,412),(230,412),1,10); #升级确定
  MonkeyRunner.sleep(1);
  
  device.shell('am force-stop com.hytc.xyol.android');

#http://gm.hanfenggame.com/hytc/qd_login.jsp
#70382  123123
 
if __name__=='__main__':
  Filename='./20110822_randomname.txt';
  File=open(Filename,'r').readlines();
  loop_job();