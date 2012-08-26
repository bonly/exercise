#!/usr/bin/python
#-*-coding:utf-8-*-

poem = """\
鹅，鹅，鹅，
曲颈向天歌，
白毛浮绿水，
红掌拨清波。"""
f = file("poem.txt","w");
f.write(poem);
f.close();

f=file("poem.txt");
while (True):
    line = f.readline();
    if (len(line)==0):
        break;
    print line;
f.close();