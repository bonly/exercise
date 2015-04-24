#!/usr/bin/env monkeyrunner
#-*- coding=utf-8 -*-
import uuid
import random
from time import gmtime, strftime

Filename='./20110822_randomname.txt'
File=open(Filename,'r').readlines()


def reg():
  name=random.choice(File)[:-1]
  #name=name+'_'
  name=name+str(uuid.uuid4())[:8]
    
  print '输入名字 ' + name

while 1:
  print 'begin time: '+  strftime("%Y-%m-%d %H:%M:%S", gmtime())
  reg();
  print 'end time: ' +  strftime("%Y-%m-%d %H:%M:%S", gmtime())

