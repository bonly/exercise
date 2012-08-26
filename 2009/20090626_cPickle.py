#!/usr/bin/python
#-*-coding:utf-8-*-

import cPickle as pick;
shoplistfile="showlist.data";
shoplist=["apple","mango","carrot"];

f=file(shoplistfile,"w");
pick.dump(shoplist,f);
f.close();

del (shoplist);

f=file(shoplistfile);
storedlist = pick.load(f);
print storedlist;