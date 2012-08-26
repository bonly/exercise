'''
Created on 2011-3-4

@author: bonly
'''
#-*-coding:utf8-*-

import subprocess

def getSysInfo():
   res = subprocess.Popen(['uname', '-sv'], stdout=subprocess.PIPE)
   uname = res.stdout.read().strip()
   return uname

def ind_find(un):
   print un
   if 'Linux' in un :
      print "'Linux' is in uname str"
   if 'Darwin' not in un:
      print "'Darwin' is not in uname"
   inx = un.index('Linux')
   print "index of Linux is %d" % inx
   fnd = un.find('Linux')
   print "find return Linuex is %d" % fnd
   #inx = un.index("Darwin")  
   print "index of not in string throw exception"  
   fnd = un.find("Darwin")
   print "find of 'Darwin' in string return %d" % fnd

def before_or_after(un):
   print '\nun is:\n %s' % un
   smp_index = un.index('SMP')
   print 'index of SMP is %d' % smp_index
   print 'un[smp_index:]: %s' % un[smp_index:]
   print 'un[:smp_index]: %s' % un[:smp_index]
   
if __name__ == '__main__':
   un = getSysInfo()
   ind_find(un)
   before_or_after(un)
   
   pass
