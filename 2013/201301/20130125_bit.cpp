#include <iostream>

#pragma pack(push)
#pragma pack(4)  //按4byte 对齐

struct ac{
	  char chr;  //1 byte
    short int  sin;  //2 byte
          int  uin;  //16位机中是2byte  32/64位机中是4byte
 unsigned int  ain:3; // 3 bit
};

#pragma pack(pop)
//按min(pack(4), ac中最大成员长度len(int))=4  用len(ac)%4来对齐

struct bc{
	int  a;
	char b;
}__attribute__((packed)); 

int main(){
   using namespace std;

   int a,b,c,d;
   a = 42;        //10进制
   b = 0x52;      //8进制
   c = 0x2a;      //16进制
   d = 0B101010;  //二进制 0b101010
   std::clog << "10: " << dec << a << std::endl;
   std::clog << "8:  " << oct << b << std::endl;
   std::clog << "16: " << hex << c << std::endl;
   std::clog << "2:  " <<  d << std::endl;
   return 0;

}

