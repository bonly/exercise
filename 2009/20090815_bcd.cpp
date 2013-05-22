//============================================================================
// Name        : BCD.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#include <cstdlib>
#include <iostream>
#include <string>
using namespace std;

string BCD_to_String(char * P_BCD, int length)
{
  string returnstring = "";
  char high_char, low_char, temp;
  for (int i = 0; i < length; i++)
  {
    temp = *P_BCD++;
    high_char = (temp & 0xf0) >> 4;
    low_char = temp & 0x0f;

    if ((high_char >= 0x00) && (high_char <= 0x09))
      high_char += 0x30;

    if ((high_char >= 0x0A) && (high_char <= 0x0F))
      high_char += 0x37;

    if ((low_char >= 0x00) && (low_char <= 0x09))
      low_char += 0x30;
    if ((low_char >= 0x0A) && (low_char <= 0x0F))
      low_char += 0x37;

    returnstring += high_char;
    returnstring += low_char;
    returnstring += " ";
  }
  return returnstring;
}

string string_To_BCD(char * P_BCD, int length)
{
  string returnstring;
  char high_char, low_char, temp;

  for (int i = 0; i < length; i++)
  {
    temp = *P_BCD++;
    high_char = (temp & 0xf0) >> 4;
    low_char = temp & 0x0f;

    if ((high_char >= 0x00) && (high_char <= 0x09))
    {
      high_char += 0x30;
    }

    if ((high_char >= 0x0A) && (high_char <= 0x0F))
    {
      high_char += 0x37;
    }

    if ((low_char >= 0x00) && (low_char <= 0x09))
    {
      low_char += 0x30;
    }
    if ((low_char >= 0x0A) && (low_char <= 0x0F))
    {
      low_char += 0x37;
    }

    returnstring += high_char;
    returnstring += low_char;
  }
  return returnstring;
}

//十进制转BCD
int ConvertBCD(int dnum)
{
  int  bcdval=0;
  if( dnum>9999 || dnum < 0 )   return   -1;
  bcdval   =   (( (dnum/1000)*16+(dnum%1000)/100)*16+(dnum%100)/10   )*16 + dnum%10;
  return   bcdval;
}

//BCD轉為十進制
int BCD_to_Int(char* BCD_String)
{
  string bcd = string("0x")+BCD_String;
  return atoi(bcd.c_str());
}

int main()
{
  cout << ":bonly^_^" << endl; // prints :bonly^_^
  return 0;
}

