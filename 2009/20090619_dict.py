#!/usr/bin/python
#-*-coding:utf-8-*-

print "Simple Assignm ent";
shoplist=["apple","mango","cannot","banana"];
mylist=shoplist; #引用
del (shoplist[0]);
print "shop list is",shoplist;
print "mylist is ",mylist;

print "copy by making a full slice";
mylist=shoplist[:]; #copy所有的内容
del (mylist[0]);

print "shoplist is",shoplist;
print "mylist is",mylist;