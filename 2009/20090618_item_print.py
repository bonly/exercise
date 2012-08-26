#!/usr/bin/python
#-*-coding:utf-8-*-

ab = {"Swaroop":"swaroopch@byteofpython.info",
      "Larry":"larry@wall.rog",
      "Matsumoto":"matz@ruby-lang.org",
      "Spammer":"spammer@hotmail.com"
};
print "Swaroop's address is %s"%ab['Swaroop'];
ab["Guido"]="gudo@python.org";
del(ab["Spammer"]);
print "There are %d contacts in the address-book\n"%len(ab);
for (name,address) in ab.items():
    print "Contact %s at %s"%(name,address);
if ("Guido") in ab:
    print "Guido's address is %s"%ab["Guido"];