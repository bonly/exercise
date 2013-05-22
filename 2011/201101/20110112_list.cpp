#include <list>
#include <iostream>
#include <algorithm>

using namespace std;
struct AA{
  int b;
  int c;
};

int main(){
  list<int> lst;
  lst.push_back(4);
  lst.push_back(10);
  lst.push_back(20);

  for(list<int>::iterator ip=lst.begin(); ip!=lst.end(); ++ip)
  {
    std::clog << *ip << std::endl;
  }

  list<int>::iterator ret = find(lst.begin(), lst.end(), 10);
  clog << "search: " << *ret << endl;


  list<AA> lb;
  lb.push_back(AA{3,4});
  lb.push_back(AA{5,6});

  for(list<AA>::iterator pp = lb.begin(); pp != lb.end(); ++pp)
  {
     clog << "lst: " << pp->b << endl;
  }
  int key = 3;
  list<AA>::iterator oret = find_if(lb.begin(), lb.end(), [=](AA k)->bool{ return key == k.b;});
  clog << "find_if: " << oret->b << endl;
  return 0;
}

