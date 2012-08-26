#!/usr/bin/python
#-*-coding:utf-8-*-
import subprocess
import re
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

def statist():
   count = []
   times = [1,10,100,1000,10000,100000,200000]
   for i in range(len(times)):
      count.append(get_second_by_re(2,times[i]))
   print "select from merge: "
   print times
   print count
   del count[:]

   for i in range(len(times)):
      count.append(get_second_by_re(3,times[i]))
   print "select from bigtable: "
   print times
   print count

if __name__ == "__main__":
   statist()