/**
 * Protocal.cpp
 * @file
 * @brief 传协议的封装及解释实现
 * @author bonly
 *
 * @date 2011-5-4 bonly created
 * @date 2011-8-16 bonly 增加路由组状态查询及路由状态变更接口协议
 */

#include "Protocal.h"
#include <arpa/inet.h>
#include <errno.h>
#include <limits.h>

namespace bas
{
    /*
    int str2l(const string &buf,int &res)
    {
        errno = 0;
        res = strtol(buf.c_str(),0,10);\
        if((errno==ERANGE && (res == LONG_MAX||res == LONG_MIN))
           ||(errno!=0 && res==0))
        {
            perror("数据不正确");
            return 1;
        }
        return 0;
    }
    */

   int BufferEncoder::Encode(const char* s)
   {
      if (first_ == false)
      {
         strncat(MsgBody_, " ", 1);// 补一个空格
         MsgLen_++;
      }
      else
         first_ = false;

      int len = strlen(s); /// 取得字符串的长度,不包括'\0'
      if (MsgLen_ + len >= (MAXLEN - HEADLEN - 1)) /// 数据包长度不能大于 MAXLEN,否则返回1
         return 1;

      strncat(MsgBody_, s, len);
      MsgLen_ += len;
      return 0;
   }

   void BufferEncoder::EndOfBuffer(int i)
   {
      Command_ = i;
      MsgBody_[MsgLen_] = '\n';
      MsgLen_ = MsgLen_ + 2 + HEADLEN; /// 最后包长是要加上包头长度
      int len = htonl(MsgLen_); /// 包长值转换为网络字节
      memcpy(buffer_, &len, 4);
      int cmd = htonl(i); /// 命令码值转换为网络字节
      memcpy(buffer_ + 4, &cmd, 4);
   }

   int BufferEncoder::GetMsg(char*& s)
   {
      s = buffer_;
      return MsgLen_;
   }

   int BufferEncoder::Decode(const char* s)
   {
      /// 解释包头
      if (0 != Decode_Head(s))
         return -1;

      int num = 0;
      /// 检查数据包是否正确
      MsgBody_ = (char*) (s + HEADLEN);
      num = strlen(MsgBody_) + HEADLEN + 1;
      return (num == MsgLen_) ? 0 : 1;
   }

   int BufferEncoder::Decode_Head(const char* s)
   {
      /// 把网络字节序的长度值转为本地字节
      int num = 0;
      memcpy(&num, s, 4);
      MsgLen_ = ntohl(num);

      /// 把网络字节序的命令码值转为本地字节
      memcpy(&num, s + 4, 4);
      Command_ = ntohl(num);

      return 0;
   }

   int Protocal::DecodeQry(const char* s)
   {
      if (buffer_.Decode(s) != 0 || s == NULL)
         return -1;

      /// 获取命令码及整个包长
      data_.qcommand = buffer_.GetCommand();
      data_.msgLen = buffer_.GetMsgLen();

      istringstream ss;
      ss.str(buffer_.GetMsgBody());
      switch(data_.qcommand)
      {
          case 1:  /// 多结果查询与
          case 3:  /// 单结果查询处理相同
          {
              /// 获取电话号码phoneNum
              ss >> data_.phoneNum;
              break;
          }
          case 5: /// 路由组状态查询
          {
              /// 获取路由组ID
              string str;
              ss >> str;
              data_.groupID = atoi(str.c_str());
              break;
          }
          case 7: /// 修改路由状态
          {
              /// 获取要修改的路由节点个数
              string str;
              ss >> str;
              data_.nodeMsgNum = atoi(str.c_str());
              /// 取出所有要修改的节点信息
              for (int i=0; i < data_.nodeMsgNum; ++i)
              {
                  /// 获取数据库连接字符串
                  ss >> str;
                  strncpy (data_.nodes[i].connectStr, str.c_str(),
                              sizeof(data_.nodes[i].connectStr));

                  /// 获取表状态
                  ss >> str;
                  data_.nodes[i].status = atoi(str.c_str());
              }
              break;
          }
      }

      return 0;
   }

   int Protocal::Encode(char*& s)
   {
      if (!this->good() && !this->eof())
      {
         return -1; /// 如果包状态是坏的并且没有结束,则返回-1
      }
      buffer_.GetMsg(s);
      return buffer_.GetMsgLen();
   }

   int Protocal::DecodeAns(const char* s)
   {
      if (buffer_.Decode(s) != 0 || s == NULL)
         return -1;

      try
      {
          /// 获取命令码及整个包长
          data_.acommand = buffer_.GetCommand();
          data_.msgLen = buffer_.GetMsgLen();

          istringstream ss;
          ss.str(buffer_.GetMsgBody());

          string str;

          /// 取ResultCode
          ss >> str;
          data_.resultcode = atoi(str.c_str());

          switch(data_.acommand)
          {
              case 0: /// 心跳直接返回0
                  return 0;
              case 2: /// 多结果查询的应答
              case 4: /// 单结果查询的应答
              {

                  if (data_.resultcode != 0) ///如果结果为失败,无需解释后面的内容
                      return 0;

                  /// 取SubNo
                  ss >> str;
                  strncpy(data_.phoneNum, str.c_str(), sizeof(data_.phoneNum));
                  /// 取SubNoEnd
                  ss >> str;
                  strncpy(data_.phoneNumEnd, str.c_str(), sizeof(data_.phoneNumEnd));
                  /// 取节点所属的路由组ID
                  ss >> str;
                  data_.groupID = atoi(str.c_str());
                  /// 取NodeMsgNum
                  ss >> str;
                  data_.nodeMsgNum = atoi(str.c_str());
                  for (int i = 0; i < data_.nodeMsgNum && i < MAXNODE; ++i)
                  {
                     /// 取instanceName
                     ss >> str;
                     strncpy(data_.nodes[i].instanceName, str.c_str(),
                              sizeof(data_.nodes[i].instanceName));
                      /// 取connectStr
                     ss >> str;
                     strncpy(data_.nodes[i].connectStr, str.c_str(),
                              sizeof(data_.nodes[i].connectStr));

                     /// 取status
                     ss >> str;
                     data_.nodes[i].status = atoi(str.c_str());

                     /// 取chreshold
                     ss >> str;
                     data_.nodes[i].chreshold = atoi(str.c_str());
                  }
                  break;
              }
              case 6: /// 路由组状态查询的应答
              {
                  /// 取GroupID
                  ss >> str;
                  data_.groupID = atoi(str.c_str());
                  //if(0 != str2l(str.c_str(),data_.groupID))
                  //    return -1;


                  /// 取NodeMsgNum
                  ss >> str;
                  data_.nodeMsgNum = atoi(str.c_str());

                  /// 取各个节点的资料
                  for (int i = 0; i < data_.nodeMsgNum && i < MAXNODE; ++i)
                  {
                      /// 取数据库连接字符串
                      ss >> str;
                      strncpy(data_.nodes[i].connectStr, str.c_str(),
                                  sizeof(data_.nodes[i].connectStr));

                      /// 取表状态
                      ss >> str;
                      data_.nodes[i].status = atoi(str.c_str());
                  }
                  break;

              }
              case 8: /// 路由状态修改的应答
              {
                  /// 取NodeMsgNum
                  ss >> str;
                  data_.nodeMsgNum = atoi(str.c_str());

                  /// 取各个节点的资料
                  for (int i = 0; i < data_.nodeMsgNum && i < MAXNODE; ++i)
                  {
                      /// 取数据库连接字符串
                      ss >> str;
                      strncpy(data_.nodes[i].connectStr, str.c_str(),
                                  sizeof(data_.nodes[i].connectStr));

                      /// 取表状态
                      ss >> str;
                      data_.nodes[i].status = atoi(str.c_str());
                  }
                  break;
              }
              default: /// 非已定义的应答返回-1
              {
                  return -1;
              }
          }

      }
      catch(...)
      {
          return -1;
      }

      return 0;
   }

}

/*//test main
using namespace bas;
using namespace std;
int main()
{
 { /// 请求
     Protocal p;
     p << "13719360007" << 1; // 以整型命令码值表示赋值结束,如果结束后再加字串字段会不让编译通过
     cout << "encoded: " << p << endl; // 打印包体内容(不包括命令码)

     char *buf = 0;
     int ret = p.Encode(buf); // ret 返回整个数据的包长度,-1表示失败
     ret = p.DecodeQry(buf); // ret 0:成功 其它:失败

     p.Reset();  // 如果想要重用对象来压包或者出错后重用同一对象压包,必须先重置一次对象,以便清理状态;解包无此要求
     p << "13719360008" << 1; // 以整型命令码值表示赋值结束,在别的地方再加入字段内容,将舍弃新加的内容
     ret = p.Encode(buf);
     cout << "encoded: " << p << endl;
 }

 { ///查询结果
     Protocal p;
     p << "0" << "1371936" << "1371937" << "2" // 结果,号码段起始,号码段结束,节点数
     << "inst1" << "abc/ok@testdb1" << "0" << "100000" //第1个节点
     << "inst2" << "abc/no@testdb2" << "0" << "100000" //第2个节点
     << 2; //命令码
     cout << "encoded: " << p << endl;

     char *buf = 0;
     int ret = p.Encode(buf); // ret 返回整个数据的包长度,-1表示失败
     ret = p.DecodeAns(buf); // buf是收到的数据包指针, ret 0:成功 其它:失败
     cout << "共有节点数: " << p.data()->nodeMsgNum << endl;
     cout << "号码段: " << p.data()->phoneNum << endl;
 }

 { /// 路由组状态查询
     Protocal p;
     p << "3" << 5; // 路由组,命令码
     cout << "encoded: " << p << endl;

     p.Reset();
     // 应答码,组ID,节点数,连接串,状态,命令码
     p << "0" << "3" << "1" << "abc/ok@testdb3" << "2" << 6;
     char *buf = 0;
     int ret = p.Encode(buf);
     ret = p.DecodeAns(buf);
     cout << "组ID: " << p.data()->groupID << endl;
     cout << "共有节点数: " << p.data()->nodeMsgNum << endl;
 }

 { /// 路由状态修改
     Protocal p;
     p << "1" << "abc/ok@pair1" << "0" << 7;
     cout << "endcode: " << p << endl;

     p.Reset();
     // 应答码,失败节点数,连接串,状态,命令码
     p << "0" << "2" << "abc/ok@test1" << "3" << "abd/ok@test2" << "4" << 8;
     char *buf = 0;
     int ret = p.Encode(buf);
     ret = p.DecodeAns(buf);
     cout << "失败节点数: " << p.data()->nodeMsgNum << endl;
 }

 return 0;
}
// */

