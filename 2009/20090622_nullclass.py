#!/usr/bin/python
#-*-coding:utf-8-*-

class Person:
    def sayHi(self):
        print "Hello, my name is",self.name;
    def __init__(self,name):
        self.name=name;

p=Person("bonly");
p.sayHi();