#!/usr/bin/python
#-*-coding:utf-8-*-
#Password generater that uses type and length.
#There are 4 types to use: alphanum, alpha, alphacap, all
#d3hydr8[at]gmail[dot]com

'''Usage: ./passgen.py  [类型] [密码长度] [密码数量]
        [选项] -w/-write  : 写入密码文件

四个选项类型: alphanum, alpha, alphacap, all
例如：生成10000个，包含数字，字母大小写、特殊符号的密码文件，密码的长度为8位，并保存到password.txt中

passgen.py all 8 10000 -w password.txt
可根据经验调整生成密码的个数和长度。也可以修改alphanum, alpha, alphacap, all这些数组的定义，生成更精确的密码。

我觉得用它生成拼音+数字的密码文件也是可以的。只需要把拼音输入脚本中定义的其中一个数组就可以了。如果你知道，不妨分享一下。

Python脚本全文：'''

import random,sys

def title():
   print "\n\t   d3hydr8[at]gmail[dot]com Password Gen v1.1"
   print "\t-----------------------------------------------\n"

def passgen(choice, length):

   passwd = ""

   alphanum = ('0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ')
   alpha = ('abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ')
   alphacap = ('ABCDEFGHIJKLMNOPQRSTUVWXYZ')
   all = ('abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&amp;*()-_+=~[]{}|\:;"\'&lt;&gt;,.?/')

   if str(choice).lower() == "alphanum":
      choice = alphanum

   elif str(choice).lower() == "alpha":
      choice = alpha

   elif str(choice).lower() == "alphacap":
      choice = alphacap

   elif str(choice).lower() == "all":
      choice = all

   else:
      print "Type doesn't match\n"
      sys.exit(1)

   return passwd.join(random.sample(choice, int(length)))

title()
if len(sys.argv) == 3 or len(sys.argv) == 5:
   print "\nUsage: ./passgen.py   "
   print "\t[options]"
   print "\t   -w/-write  : Writes passwords to file\n"
   print "There are 4 types to use: alphanum, alpha, alphacap, all\n"
   sys.exit(1)

for arg in sys.argv[1:]:
   if arg.lower() == "-w" or arg.lower() == "-write":
      txt = sys.argv[int(sys.argv[1:].index(arg))+2]

if sys.argv[3].isdigit() == False:
   print sys.argv[3],"must be a number\n"
   sys.exit(1)
if sys.argv[2].isdigit() == False:
   print sys.argv[2],"must be a number\n"
   sys.exit(1)
try:
   if txt:
      print "[+] Writing Data:",txt
      output = open(txt, "a")
except(NameError):
   txt = None
   pass

for x in xrange(int(sys.argv[3])):
   if txt != None:
      output.writelines(passgen(sys.argv[1],sys.argv[2])+"\n")
   else:
      print "Password:",passgen(sys.argv[1],sys.argv[2])
print "\n[-] Done\n"