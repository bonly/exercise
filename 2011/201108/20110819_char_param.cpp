#include <string.h>
#include <iostream>
int fun(char *&player_name){
  strcpy(player_name, "hello");
  return 0;
}

int main(){
  char aaa[25]="";
  char *p = aaa;
  fun (p);
  std::clog << aaa << std::endl;
  return 0;
}

