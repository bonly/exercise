#include <sstream>
#include <iostream>
using namespace std;

int main(){
  unsigned int x;
  std::stringstream ss;
  ss << std::hex << "fffefffe";
  ss >> x;
  cout << x << endl;
  cout << static_cast<int>(x) << endl;
  return 0;
}

