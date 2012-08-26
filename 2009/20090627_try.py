#!/usr/bin/python
#-*-coding:utf-8-*-
import sys;
try:
    s=raw_input("enter some thing->");
except E0FError:
    print "Why did you do an EOF on me?";
    sys.exit();
except:
    print "Some error exception occurred.";
print "Done";