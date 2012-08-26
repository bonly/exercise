#!/usr/bin/python
#-*-coding:utf-8-*-
import subprocess
import re
from GChartWrapper import *
#import urllib2

def get_second_by_re(param,count):
   cmd = "time ./mysql " + str(param)
   if param in [2,3]:
      cmd = cmd + " " + str(count)
   res = subprocess.Popen([cmd],stderr=subprocess.PIPE,stdout=subprocess.PIPE,shell=True)
   #res = subprocess.Popen(["time ls -l"], stderr=subprocess.PIPE, stdout=subprocess.PIPE, shell=True)
   t0 = res.stderr.read().strip()
   re_obj = re.compile(r'.m.*s')
   result = re_obj.findall(t0)
   return result[0]

def get_realtime_by_re(param):
   re_se = re.compile(r'[0-9]*')
   re_mi = re.compile(r'[0-9]*\.[0-9]*')
   result = (int((re_se.findall(param))[0])) * 60  * 1000 + float((re_mi.findall(param))[0])*1000
   return result

def statist():
   count = []
   times = [10]
   realtime1 =[]
   for i in range(len(times)):
      tm = get_second_by_re(2,times[i])
      count.append(tm)
      realtime1.append( get_realtime_by_re(tm))
   print "select from merge: "
   print times
   #print count
   print realtime1
   del count[:]

   realtime2 = []
   for i in range(len(times)):
      tm = get_second_by_re(3,times[i])
      count.append(tm)
      realtime2.append( get_realtime_by_re(tm))
   print "select from bigtable: "
   print times
   #print count
   print realtime2
   
if __name__ == "__main__":
   statist()