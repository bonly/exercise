/**
 * @file 20100317_vector.cpp
 * @brief
 *
 * @author bonly
 * @date 2011-11-22 bonly created
 */


#include <malloc.h>
#include <vector>
#include <iostream>
using namespace std;

void testmap()
{
  cout << "begin stat:\n";
  malloc_stats();
  cout << "begin to create data:\n";

  vector<int> testmap; //方法1
  //map<int, float, less<int>, std::allocator<pair<int, float> > > testmap; //方法2
  //map<int, float, less<int>, __gnu_cxx::new_allocator<pair<int, float> > > testmap;  //方法3
  //方法1/2/3同效果

  //新版中找不到此方法了?
  //map<int, float, less<int>, __gnu_cxx::__pool_alloc<pair<int, float> > >  testmap; //方法4，使用了基于cache的allocator
  testmap.resize(1000000);
  for (int i = 0; i < 1000000; i++) {
    testmap[i] = (int)i;
  }
  malloc_stats();
  int tmp; cout << "use ps to see my momory now, and enter int to continue:"; cin >> tmp;
  testmap.clear();
  vector<int> btmp;
  testmap.swap(btmp);
  malloc_stats();
  cout << "use ps to see my momory now, and enter int to continue:"; cin >> tmp;
  testmap.resize(0);
  malloc_stats();
  cout << "use ps to see my momory now, and enter int to continue:"; cin >> tmp;
}

int main()
{
    //malloc_info();  ///把内存资料写到文件中
    testmap();
    return 0;
}




