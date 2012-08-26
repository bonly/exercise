/*
 * 20100330_addtion.cpp
 *
 *  Created on: 2012-2-17
 *      Author: Bonly
 */
#include <mem.h>
struct Point
{
    int x,y,z,w;
    Point():x(0),y(0),z(0),w(1){}
    Point(int ix, int iy, int iz, int iw):x(ix),y(iy),z(iz),w(iw){}
};

template<typename C, typename T>
struct OBJ : public C
{
  T  attr;
  void operator=(const C& b)
  {
    memcpy (this, (const void*)&b, sizeof(b));
    return;
  }
  void operator=(const T& b)
  {
    memcpy (this+sizeof(C), (const void*)&b, sizeof(b));
    return;
  }
};

template<typename C, typename T>
struct ATTR : public C, T
{
  void operator=(const C& b)
  {
    memcpy (this, (const void*)&b, sizeof(b));
    return;
  }
  void operator=(const T& b)
  {
    memcpy (this+sizeof(C), (const void*)&b, sizeof(b));
    return;
  }
};

class myClass
{
  public:
    int mx;
};

struct othrFun
{
    int ox;
};

#include <iostream>
using namespace std;
int main()
{
   OBJ<myClass,Point> test;
   test.mx = 10;
   test.attr.w = 11;

   clog << "w=" << test.attr.w << endl;

   myClass a;
   a.mx = 23;
   OBJ<myClass,Point> test2;
   test2.attr.x = 30;
   test2 = a;
   clog << "mx= " << test2.mx << endl
       << "x= " << test2.attr.x << endl;

   myClass *b = new myClass();
   b->mx = 25;
   OBJ<myClass,Point> test3;
   test3.attr.x = 31;
   test3 = *b;
   clog << "mx= " << test3.mx << endl
       << "x= " << test3.attr.x << endl;

   OBJ<othrFun,OBJ<myClass,Point> > test4;
   test4.attr.mx = 34;
   test4.ox = 13;
   test4.attr.attr.z = 32;
   clog << "mx= " << test4.attr.mx << endl
       << "z= " << test4.attr.attr.z << endl;

   OBJ<OBJ<myClass,Point>,othrFun> test5;
   test5.mx = 44;
   test5.attr.ox = 14;


   ATTR<myClass,Point> test6;
   test6.mx = 10;
   Point k(13,34,0,2);
   test6 = k;
   clog << "point x= " << test6.x << endl
       << "mx= " << test6.mx << endl;


   return 0;
}
