/**
 *  @file 201004015_array.cpp
 *
 *  @date 2012-2-29
 *  @Author: Bonly
 */
#include <iostream>
using namespace std;

class Ar
{
  public:
    int** p;

    Ar()
    {
      ///[3][4]
      p = new int*[3];
      int k = 0;
      for (int i=0; i<3; ++i)
      {
         p[i] = new int;
         for (int j=0; j<4; ++j)
         {
           p[i][j] = k;
           ++k;
         }
      }
    }
    ~Ar()
    {
      delete []p[0];
      delete []p[1];
      delete []p[2];
      delete []p;
    }
};

int main()
{
  Ar ar;
  clog << ar.p[0][0] << endl;
  return 0;
}



