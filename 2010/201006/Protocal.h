/**
 * Protocal.h
 * @file
 * @brief 传协议的封装及解释声明
 *
 * @author bonly
 * @date 2011-5-4 bonly created
 * @date 2011-5-5 bonly Protocal类增加GetMsg()取得数据包,增加data()取得指向数据结构的指针
 * @date 2011-5-5 bonly 如果封包操作已结束,舍弃再新增到包中的字串内容,改名Protocal的GetMsg()为Encode()
 * @date 2011-5-5 bonly 重用Protocal实例做多次压包操作,需先Reset()
 * @date 2011-8-18 bonly 增加路由组查询及修改路由状态
 */

#ifndef PROTOCAL_H_
#define PROTOCAL_H_
#include <string.h>
#include <string>
#include <fstream>
#include <iostream>
#include <sstream>
#include <cstdlib>

/// @brief 定义bas命名空间,以免污染调用者的空间
namespace bas
{
   using namespace std;
   /**
    * @brief 定义各种用到的常量值
    */
   enum
   {
      MAXLEN = 1024, ///< 缓冲区的最大长度
      HEADLEN = 8, ///< 包头长度
      MAXFIELD = 20, ///< 包体最大字段数
      MAXNODE = 8, ///< 结果包中返回最大的节点数
      END = 0
   };

   /**
    * @brief 各个节点的数据描述
    */
   struct nodeData
   {
         char instanceName[8 + 1]; ///< 实例名
         char connectStr[64 + 1]; ///< 连接串
         //char tableName[32 + 1]; ///< 表名
         int chreshold; ///< 阀值
         int status; ///< 状态
   };

   /**
    * @brief 数据包的数据结构
    */
   struct CR
   {
         char qcommand; ///< 请求命令码
         char acommand; ///< 应答命令码
         long msgLen; ///< 整个数据包的长度
         int resultcode; ///< 结果码
         char phoneNum[32 + 1]; ///< 号码或号码段起始
         char phoneNumEnd[32 + 1]; ///< 号码段结束
         int nodeMsgNum; ///< 节点数量
         int groupID; ///< 路由组
         nodeData nodes[MAXNODE]; ///< 节点数组
   };

   /**
    * @brief 数据包的基础处理
    */
   class BufferEncoder: public streambuf
   {
      private:
         char buffer_[MAXLEN]; ///< 缓冲区
         int MsgLen_; ///< 整个消息包的长度
         int Command_; ///< 命令码
         char *MsgBody_; ///< 数据体
         bool first_; ///< 用于标识是否第一个字段
      public:
         BufferEncoder() :
            streambuf(), MsgLen_(0), Command_(-1),
                     first_(true)
         {
            memset(buffer_, 0, MAXLEN);
            MsgBody_ = buffer_ + HEADLEN;
         }
         ~BufferEncoder()
         {
         }

         /** @brief 获取当前消息长度
          *
          */
         inline int GetMsgLen()
         {
            return MsgLen_;
         }

         /** @brief 获取命令码
          *
          */
         inline int GetCommand()
         {
            return Command_;
         }

         /** @brief 清理状态位及缓冲中的数据
          *
          */
         inline void Clear()
         {
            memset(buffer_, 0, MAXLEN);
            MsgBody_ = buffer_ + HEADLEN;
            MsgLen_ = 0;
            Command_ = -1;
            first_ = true;
         }

         /** @brief 压一个数据到包中
          *  @return 0:成功 其它:失败
          */
         int Encode(const char* s);

         /** @brief 缓冲区结束
          *  @param i 命令码
          */
         void EndOfBuffer(int i);

         /** @brief 取消息体部分内容
          *
          */
         inline char* GetMsgBody()
         {
            return MsgBody_;
         }

         /** @brief 取得完整的消息包指针
          *  @param *s 指向消息包的指针
          *  @return 完整的包长
          */
         inline int GetMsg(char*& s);

         /** @brief 解包
          *  @return 0:成功 1:失败
          */
         int Decode(const char* s);

         /** @brief 解包头
          *  @return 0:成功 -1:失败
          */
         int Decode_Head(const char* s);

   };

   /** @brief 协议压包及解释类
    *  @par 封请求包使用范例
    @code
    Protocal p;
    p << "13719360007" << 1 ;  // 以整型命令码值表示赋值结束,如果结束后再加字串字段会不让编译通过
    cout << "encoded: " << p << endl;  // 打印包体内容(不包括命令码)

    char *buf = 0;
    int ret = p.Encode(buf);  // ret 返回整个数据的包长度,-1表示失败
    ret = p.DecodeQry(buf);  // ret 0:成功 其它:失败

    p.Reset();  // 如果想要重用对象来压包或者出错后重用同一对象压包,必须先重置一次对象,以便清理状态;解包无此要求
    p << "13719360008" << 1; // 以整型命令码值表示赋值结束,在别的地方再加入字段内容,将舍弃新加的内容
    ret = p.Encode(buf);
    cout << "encoded: " << p << endl;
    @endcode
    *  @par 封结果包使用范例
    @code
    Protocal p;
    p << "0" << "1371936" << "2"   // 结果,号码段,节点数
      << "inst1" << "abc/ok@testdb1" << "0" << "100000"  //第1个节点
      << "inst2" << "abc/no@testdb2" << "0" << "100000"  //第2个节点
      << 2;  //命令码
    cout << "encoded: " << p << endl;

    char *buf = 0;
    int ret = p.Encode(buf); // ret 返回整个数据的包长度,-1表示失败
    @endcode
    *  @par 解包使用范例
    @code
    Protocal p;
    ret = p.DecodeAns(buf);  // buf是收到的数据包指针, ret 0:成功 其它:失败
    cout << "共有节点数: " << p.data()->nodeMsgNum << endl;
    cout << "号码段: " << p.data()->phoneNum << endl;
    @endcode
    *  @par 路由组状态查询
    @code
     Protocal p;
     p << "3" << 5; // 路由组,命令码
     cout << "encoded: " << p << endl;

     p.Reset();
     //应答码,组ID,节点数,连接串,状态,命令码
     p << "0" << "3" << "1" << "abc/ok@testdb3" << "2" << 6;
     char *buf = 0;
     int ret = p.Encode(buf);
     ret = p.DecodeAns(buf);
     cout << "组ID: " << p.data()->groupID << endl;
     cout << "共有节点数: " << p.data()->nodeMsgNum << endl;
    @endcode
    *   @par 路由状态修改
    @code
     Protocal p;
     p << "1" << "abc/ok@pair1" << "0" << 7;
     cout << "endcode: " << p << endl;

     p.Reset();
     //应答码,失败节点数,连接串,状态,命令码
     p << "0" << "2" << "abc/ok@test1" << "3" << "abd/ok@test2" << "4" << 8;
     char *buf = 0;
     int ret = p.Encode(buf);
     ret = p.DecodeAns(buf);
     cout << "失败节点数: " << p.data()->nodeMsgNum << endl;
    @endcode
    */
   class Protocal: public ostream
   {
      private:
         BufferEncoder buffer_;
         CR data_;

      public:
         Protocal(filebuf* fb) :
            ios(0), ostream(fb)
         {
         }
         Protocal() :
            ios(0), ostream(&buffer_)
         {
         }

         /** @brief 提供指向数据结构的指针给外部使用
          *  @return CR*
          */
         inline CR* data()
         {
            return &data_;
         }

         /** @brief 结束数据包
          *  @param i 任何数值类数据表示结束
          */
         inline void operator<<(int i)
         {
            buffer_.EndOfBuffer(i);
            this->setstate(eofbit);
         }

         /** @brief 加入数据到包中
          *  @param *s 要加入的字符串
          */
         inline Protocal& operator<<(const char* s)
         {
            if (this->good())
            {
               if (0 != buffer_.Encode(s))
                  this->setstate(badbit);
            }
            else
               cerr << "封包错误,操作已经结束,舍弃新增内容[" << s << "]!\n";
            return *this;
         }

         /** @brief 输出到流中
          *
          */
         friend ostream& operator<<(ostream &s, Protocal& c)
         {
            s << c.buffer_.GetMsgBody();
            return s;
         }

         /** @brief 取得缓冲区中的数据包指针
          *
          */
         inline BufferEncoder* Buffer()
         {
            return &buffer_;
         }

         /** @brief 只解包头得到长度和命令码,用以socket取包体
          *  @return 0:成功 其它:失败
          */
         inline int DecodeHead(const char* s)
         {
            if (0 != buffer_.Decode_Head(s))
               return -1;
            data_.qcommand = buffer_.GetCommand();
            data_.acommand = buffer_.GetCommand();
            data_.msgLen = buffer_.GetMsgLen();
            return 0;
         }

         /** @brief 解请求包
          *  @param *s 指向请求包数据缓冲区的指针
          *  @return 0:成功 其它:失败
          */
         int DecodeQry(const char* s);

         /** @brief 解应答包
          *  @param *s 指向应答包数据缓冲区的指针
          *  @return 0:成功 其它:失败
          */
         int DecodeAns(const char* s);

         /** @brief 取得指向数据包的指针
          *  @param *&s 用于存放指向数据包的指针
          *  @return 返回包长 值为-1时表示失败
          */
         int Encode(char*& s);

         /** @brief 清空状态位,以便压包时重用对象
          *
          */
         void Reset()
         {
            buffer_.Clear();
            clear();
         }

         virtual ~Protocal()
         {
         }
   };
}

#endif /* PROTOCAL_H_ */
