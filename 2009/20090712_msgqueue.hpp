#ifndef __MSG_QUEUE_H__
#define __MSG_QUEUE_H__

struct msgbuff
{
    long mtype;
    char mtext[512];
};

class CMsgQueue
{
public:
    CMsgQueue();
    ~CMsgQueue();
    
    int Open(const char *szPathName);
    int Close();
    
    int PutMsg(const char *szBuff, int nMsgLen, long lMsgType = 0);
    int GetMsg(char *szBuff, int nSize, long lMsgType = 0);

    int GetQueueDepth() const;
    
    int IsOpen() const;
    
    static CMsgQueue* qry();
    static CMsgQueue* ret();

private:
    int            m_nMsgId;
    struct msgbuff m_msg;    // 记录当时操作的消息
    static CMsgQueue* qry_;
    static CMsgQueue* ret_;
};

#endif // __MSG_QUEUE_H__