#!/usr/bin/python
#-*-coding:utf-8-*-

shoplist=["apple",'mango',"Cannot",'banana'];
print "I have",len(shoplist),"items to purchase.";
print("These items are:");
for item in shoplist:
    print (item);
print("\nI also have to buy rice.");
shoplist.append("rice");
print "My shopping list is now:",shoplist;
print "I will sort my list now";
shoplist.sort();
print "sorted shopping list is:",shoplist;

print "the first item I will buy is,",shoplist[0];
olditem=shoplist[0];
del shoplist[0];
print "I bought the",olditem;
print "my shopping list is now:",shoplist;