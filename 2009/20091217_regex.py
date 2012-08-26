#!/usr/bin/python
#-*-coding:utf-8-*-
import subprocess
import re
import urllib2

def get_second_by_strip():
   #res = subprocess.Popen(["time /home/bonly/worksp/mysql/Debug/mysql 0"],stderr=subprocess.PIPE,stdout=subprocess.PIPE,shell=True)
   res = subprocess.Popen(["time ls -l"], stderr=subprocess.PIPE, stdout=subprocess.PIPE, shell=True)
   t0 = res.stderr.read().strip()
   ind = t0.split('\t')
   ind = ind[1].split('\n')
   print ind[0]

def get_second_by_re():
   #res = subprocess.Popen(["time /home/bonly/worksp/mysql/Debug/mysql 0"],stderr=subprocess.PIPE,stdout=subprocess.PIPE,shell=True)
   res = subprocess.Popen(["time ls -l"], stderr=subprocess.PIPE, stdout=subprocess.PIPE, shell=True)
   t0 = res.stderr.read().strip()
   re_obj = re.compile(r'.m.*s')
   result = re_obj.findall(t0)
   print result  #返回的列表,可通过result[0]访问第一个

def get_iter_by_re():
   #res = subprocess.Popen(["time /home/bonly/worksp/mysql/Debug/mysql 0"],stderr=subprocess.PIPE,stdout=subprocess.PIPE,shell=True)
   res = subprocess.Popen(["time ls -l"], stderr=subprocess.PIPE, stdout=subprocess.PIPE, shell=True)
   t0 = res.stderr.read().strip()
   #print t0
   re_obj = re.compile(r'(?P<num>\b.m.*s\b)', re.VERBOSE | re.IGNORECASE )
   iter = re_obj.finditer(t0)
   print type(iter)
   #print map(str,iter)
   #print "%s: %s" % (iter.start(),iter.end(0))
   for match in iter:
      print "%s: %s" % (match.start(), match.group("num"))

def get_url_re():
   html = urllib2.urlopen('http://www.google.com.hk/search?hl=zh-CN&newwindow=1&safe=strict&q=python+finditer+compile&btnG=Google+%E6%90%9C%E7%B4%A2&aq=f&aqi=&aql=&oq=').read()
   pattern = r'\b(the\s+\w+)\s+'
   regex = re.compile(pattern, re.IGNORECASE)
   for match in regex.finditer(html):
      print "%s: %s" % (match.start(), match.group(1))

if __name__ == "__main__":
   #get_second_by_strip()
   #get_second_by_re()
   get_iter_by_re()
   #get_url_re()
