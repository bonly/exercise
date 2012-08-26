#!/usr/bin/python

number=23;
running=True;
while (running):
    guess=int(raw_input("Enteran integer"));
    if guess==number:
        print ("Congratulations you guested it");
        running=False;
    elif guess<number:
        print ("No,it is little higher than that");
    else:
        print ("No,it is a little bower than that");

print ("Done");