//============================================================================
// Name        : bindump_hex.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#include <cstdio>
#include <string.h>
#include <ctype.h>

struct FormatChar
{
      FormatChar(char c, char* s)
      {
         sprintf((char*) s, "%02X", c);
         s[2] = ' ';
      }
};

template<typename PRINT = FormatChar, int LINE_SIZE = 16>
class FormatLine
{
   private:
      enum
      {
         REAL_SIZE = LINE_SIZE * 4 + 4
      };
      /// 定义一行buffer
      char line[REAL_SIZE];

   public:
      friend struct FormatChar;

      char* operator()(char* p, size_t& nBegin, size_t& nLeft)
      {
         if (nLeft <= 0)
            return NULL;

         /// 指向字符开始位置
         char* pt = p + nBegin;

         /// 每次取n个byte出来处理
         for (size_t i = 0; i < LINE_SIZE; ++i)
         {
            if (nLeft > 0)
            {
               /// 把字符打印到左部
               char c = pt[i];
               PRINT(c, line + i * 3);

               /// 转换后打印到右部,不可显字符显示为'.'
               line[LINE_SIZE * 3 + 3 + i] = isprint(c) ? c : '.';

               /// 剩余未打印字符数量减１
               --nLeft;
            }
            else ///　全部字符处理完成,但未完成一行所需字符时补空格
            {
               /// 左部结束后补空格
               line[i * 3] = ' ';
               line[i * 3 + 1] = ' ';

               ///　右部结束后加补空格
               line[LINE_SIZE * 3 + 3 + i] = ' ';
            }
         }
         /// 左部末尾显示空格
         line[LINE_SIZE * 3] = ' ';
         line[LINE_SIZE * 3 + 1] = ' ';
         line[LINE_SIZE * 3 + 2] = ' ';

         /// 右部末尾显示空格
         line[REAL_SIZE - 1] = '\0';
         /// 偏移记录开始的位置
         nBegin += LINE_SIZE;
         return line;
      }
};

template<typename LINE = FormatLine<FormatChar,16> >
void hexdump(char* p, size_t len)
{
   size_t i = 0;
   size_t j = len;
   char* s = 0;
   LINE ft;
   while (0 != (s = ft(p, i, j)))
      printf("%s\n", s);
}

int main()
{
   char test[] = "this is a sample\n, and how you feel about these ?\n";
   /*
   size_t i = 0;
   size_t j = strlen((const char*) test);
   char *buf = 0;
   FormatLine<FormatChar, 6> ft;
   while (0 != (buf = ft(test, i, j)))
   {
      printf("%s\n", buf);
   }
   */

   //hexdump<FormatLine<FormatChar,16> >((char*)test, strlen(test));
   hexdump((char*)test, strlen(test));

   return 0;
}

/*
 g++ -std=c++0x
 */
