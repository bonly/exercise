#!/usr/bin/python
#-*-coding:utf-8-*-
'''
Created on 2011-3-5

@author: bonly
'''
import sys
import subprocess;
import os

def modify_param():
   for i in range(len(sys.argv) - 1):
      print sys.argv[i + 1]
      if (sys.argv[i + 1]).find('/') == 0 or (sys.argv[i + 1]).find('~') == 0:
         pass
      elif (sys.argv[i + 1]).find('-') == 0 or (sys.argv[i + 1]).find('--') == 0:
         pass
      else:
         sys.argv[i + 1] = os.getcwd() + "/" + sys.argv[i + 1]
         print sys.argv[i + 1]
         pass
      

      
   
if __name__ == "__main__":
    execname = "/home/bonly/peazip/peazip"
    modify_param()
    param = " ".join(sys.argv[1:])
    print param

    subprocess.call([execname])
