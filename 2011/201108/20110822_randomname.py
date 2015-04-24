#!/usr/bin/env python

#  http://www.youtube.com/user/theregrunner

import random

### change this path to your path ####
Filename='./20110822_randomname.txt'
File=open(Filename,'r').readlines()
name=random.choice(File)[:-1]
name=name+' '
name=name+random.choice(File)[:-1]
print name.lower()


