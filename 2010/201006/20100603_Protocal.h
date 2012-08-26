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

#include "20100603_chat.h"
#include "20100605_chat.h"
#include <new>
/// @brief 定义bas命名空间,以免污染调用者的空间
namespace bas {
using namespace std;
/**
 * @brief 定义各种用到的常量值
 */
enum
{
    MAXLEN = 1024, ///< 缓冲区的最大长度
    HEADLEN = 24, ///< 包头长度
};

/**
 * @brief 数据包的数据结构
 */
struct Pack
{
        int len; ///< 包体总长
        int cmd; ///< 命令码
};

enum
{
    headlen = sizeof(Pack)
};

class Protocol
{
    public:
        enum
        {
            bad, good
        } stat;

        Protocol(char *buf, const size_t len) :
                    stat(good), _buf(buf)
        {
            init(buf, len);
        }
        int init(char *buf, const size_t len)
        {
            if (len < headlen || buf == 0)
                stat = bad;
            _head = (Pack*) buf;
            _data = buf + headlen;
            return stat;
        }
        int Encode()
        {
            if (stat == bad)
                return -1;
            return _head->len;
        }
        Protocol& operator<<(const int cmd)
        {
            if (stat == bad)
                return *this;
            _head->cmd = cmd;
            return *this;
        }
        template<typename T>
        Protocol& operator<<(T *t)
        {
            if (stat == bad)
                return *this;
            _head->len = t->write(_data, 0);
            return *this;
        }
        template<typename T>
        Protocol& operator<<(T &t)
        {
            if (stat == bad)
                return *this;
            _head->len = t.write(_data, 0);
            return *this;
        }
        int Decode()
        {
            if (stat == bad)
                return -1;
            return _head->len;
        }
        Protocol& operator>>(int &cmd)
        {
            if (stat == bad)
                return *this;
            cmd = _head->cmd;
            return *this;
        }
        template<typename T>
        Protocol& operator>>(T *t)
        {
            if (stat == bad)
                return *this;
            t->read(_data, 0);
            return *this;
        }
        template<typename T>
        Protocol& operator>>(T &t)
        {
            if (stat == bad)
                return *this;
            t.read(_data, 0);
            return *this;
        }

    private:
        char *_buf; ///< 缓冲区
        Pack *_head; ///< 包头
        void *_data; ///< 数据
};
}

template<typename T>
struct CMD: public T
{
        virtual int write(void *_buffer, int offset) = 0;
        virtual int read(void *_buffer, int offset) = 0;
};

struct CmdTalk: public CMD<Talk>
{
        virtual int write(void *buf, int off)
        {
            return Talk_write_delimited_to(this, buf, off);
        }
        virtual int read(void *_buffer, int offset)
        {
            return Talk_read_delimited_from(_buffer, this, offset);
        }
};

struct CmdAct: public CMD<ACTION>
{
        virtual int write(void *buf, int off)
        {
            return ACTION_write_delimited_to(this, buf, off);
        }
        virtual int read(void *_buffer, int offset)
        {
            return ACTION_read_delimited_from(_buffer, this, offset);
        }
};
#endif /* PROTOCAL_H_ */
