#!/usr/bin/python
#-*-coding:utf-8-*-

def printMax(x,y):
    '''Prints the maxnum of two number'''
    x=int(x);
    y=int(y);
    if (x>y):
        print(x,"is maxnum");
    else:
        print(y,"is maxnum");

printMax(3,5);
print printMax.__doc__;  #doc是属性，不是方法，所以不能加()