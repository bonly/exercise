#!/usr/bin/python
#-*-coding:utf-8-*-
'''
Created on 2011-3-11

@author: bonly
'''
from combinatorics.all_pairs2 import all_pairs2
def create_data():
   fil = open("./testim.txt","w")
   count = 100
   begin_id = 500001
   phone = []
   phonen = []
   for i in range(20):
      phone.append(13719360007+i)
   for i in range(20):
      phonen.append(13848004162+i)
      
   field = [["Windows", "Linux", "iOS", "Android", "DOS"],
            ["98", "NT", "2000", "XP", "win7", "win8", "2.4", "2.6", "Debian", "Gennto",
             "RedHat", "FeDroa", "Ubuntu", "Mint", "4.3", "1.5", "1.6", "2.0", "2.1", "2.2", "2.3", "2.3.3"],
            ['mikey', 'miny', 'mike', 'jeson', 'dance', 'rose', 'yellow', 'red', 'green', 'blue'],
            ['night', 'afternoon', 'morning'],
            ['20070103', '20081211', '20100308', '20030708', '20040802'],
            phone,
            phonen,
            ['34113444213', '4019394013', '8473210003', '39103874433', '1139387588'],
            ['34113444211', '4019394016', '8473210002', '39103874439'],
            ['1997-05-03', '2009-01-28', '2008-08-21', '2011-02-01', '2000-09-12', '2002-03-06'],
            ['2003-11-23', '2003-11-22', '2003-12-31', '2005-09-12', '2006-07-01', '2006-08-03', '2007-07-01'],
            ['2003-10-23', '2002-11-22', '2003-12-31', '2006-09-12', '2003-07-01', '2001-08-03', '2007-03-01', '2010-11-26']
           ]
   pair = all_pairs2(field)
   for i, v in enumerate(pair):
      #print "insert into testmk value(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)"% \
      #         (i,v[0],v[1],v[2],v[3],v[4],v[5],v[6],v[7],v[8],v[9],v[10],v[11])
      if i >=count :
         fil.close()
         return
      fil.write("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n"%(str(i+begin_id),v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11]))
   m = i
   while m % (count + 1) != 0 :
      if m >= count :
         fil.close()
         return
      pair = all_pairs2(field)
      for i, v in enumerate(pair):
         m = m + 1
         if m >= count :
            fil.close()
            return
         fil.write("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n" %(str(begin_id+m), v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9], v[10], v[11]))
   fil.close()
         
if __name__ == '__main__':
   create_data()

