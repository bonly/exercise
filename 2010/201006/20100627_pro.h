/**
 * Protocal.h
 * @file
 * @brief 传协议的封装及解释声明
 *
 * @author bonly
 * @date 2012-7-20 bonly created
 */

#ifndef PROTOCAL_H_
#define PROTOCAL_H_
#include <cstddef>
#include <arpa/inet.h>
#include <string.h>

#ifndef __MTK__
#define hton(x) htonl(x)
#define ntoh(x) ntohl(x)
#define MEMCPY memcpy
#endif

/**
 * @brief 数据包的数据结构
 */
struct Pack
{
        unsigned int len; ///< 包体总长
        unsigned int cmd; ///< 命令码
};

/**
 * @brief 定义各种用到的常量值
 */
enum
{
    MAXLEN = 1024, ///< 缓冲区的最大长度
    HEADLEN = sizeof(Pack), ///< 包头长度
};

extern unsigned long (*cmd_size)(void*, int);
extern int (*cmd_can_read)(void *, int, int);

template<typename T>
struct Cmd: public T
{
        virtual int write(void *_buffer, int offset) = 0;
        virtual int read(void *_buffer, int offset) = 0;
        virtual ~Cmd(){};
};

class Protocol
{
    public:
        enum
        {
            bad, good
        } stat;

        Protocol() :
                    stat(bad),_buf(0),_head(0),_data(0),_cmd(0),_data_len(0)
        {
        }
        Protocol(char *buf, const unsigned int len = HEADLEN) :
                    stat(good), _buf(buf),_data_len(0)
        {
            init(buf, len);
        }
        int init(char *buf, const unsigned int len = HEADLEN)
        {
            if (len < HEADLEN || buf == 0)
                stat = bad;
            _head = (Pack*) buf;
            _data = buf + HEADLEN;
            stat = good;
            return stat;
        }
        /**
         * @note 须在压入数据后再调用
         */
        int Encode(bool check=true)
        {
            if (stat == bad)
                return -1;
            if (true == check)
            {
                if (0 == cmd_can_read(_data, 0, _data_len))
                {
                    stat = bad;
                    return -1;
                }
            }
            return _data_len + HEADLEN;
        }
        Protocol& operator<<(const unsigned int cmd)
        {
            if (stat == bad)
                return *this;
            _cmd = cmd;
            _head->cmd = hton(cmd);
            return *this;
        }
        template<typename T>
        Protocol& operator<<(T *t)
        {
            if (stat == bad)
                return *this;
            _data_len = t->write(_data, 0);
            _head->len = hton(_data_len);  //cmd_size(_data, 0);
            return *this;
        }
        template<typename T>
        Protocol& operator<<(T &t)
        {
            if (stat == bad)
                return *this;
            _data_len = t.write(_data, 0);
            _head->len = hton(_data_len); //cmd_size(_data, 0);
            return *this;
        }

        /**
         * 解包
         * @param check true:详解并验证 false:只取包头长度
         * @return -1:失败 其它：成功的包长
         */
        int Decode(bool check = false)
        {
            if (stat == bad)
                return -1;
            if (true == check)
            {
                if (0 == cmd_can_read(_data, 0, _head->len))
                {
                    stat = bad;
                    return -1;
                }
            }
            _data_len = ntoh(_head->len);
            _cmd = ntoh(_head->cmd);
            return _data_len + HEADLEN;
        }
        /**
         * @note 须先调用Decode
         */
        Protocol& operator>>(unsigned int &cmd)
        {
            if (stat == bad)
                return *this;
            cmd = _cmd;
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
        /**
         * @note 可在Decode前调用
         */
        unsigned int getLen()
        {
            if (stat == bad)
                return -1;
            return ntoh(_head->len);
        }
        /**
         * @note 可在Decode前调用
         */
        unsigned int getCmd()
        {
            if (stat == bad)
                return -1;
            return ntoh(_head->cmd);
        }
        Pack* getHead()
        {
            return _head;
        }
        void* getData()
        {
            return _data;
        }
        char* getBuf()
        {
            return _buf;
        }

    private:
        char *_buf; ///< 缓冲区
        Pack *_head; ///< 包头
        void *_data; ///< 数据
        unsigned int _cmd; ///< host格式的命令码
        unsigned int _data_len; ///< host格式的长度
};

#define CMD(N) \
    struct Cmd##N: public Cmd<N> \
    { \
        virtual int write(void *buf, int off) \
        { \
            return N##_write_delimited_to((N*)this, buf, off); \
        } \
        virtual int read(void *buf, int off) \
        { \
            return N##_read_delimited_from(buf, (N*)this, off); \
        } \
        virtual ~Cmd##N() {} \
    }
#endif /* PROTOCAL_H_ */
