#!/usr/bin/python
#-*-coding:utf-8-*-

class Person:
  """Represents a person."""
  population=0;

  def __init__(self,name):
    """Initailizes the persion's data."""
    self.name=name;
    print "(Initializing %s)"%self.name;
    Person.population+=1;

  def __del__(self):
    """I am dying."""
    print "%s says bye."%self.name;
    Person.population-=1;
    if (Person.population==0):
        print "I am the lastone.";
    else:
        print "There are still %d people life."%Person.population;

  def sayHi(self):
    """Greeting by the person."""
    print "Hi, my name is %s."%self.name;

  def howMany(self):
    if (Person.population==1):
        print "I am the only person here.";
    else:
        print "We have %d persons here."%Person.population;

swaroop=Person("Swaroop");
swaroop.sayHi();
swaroop.howMany();

kalam=Person("Abdul Kalam");
kalam.sayHi();
kalam.howMany();

swaroop.sayHi();
swaroop.howMany();


