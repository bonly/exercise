#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/msg.h>

#include <unistd.h>  // for sleep()


#include "msgqueue.h"

CMsgQueue* CMsgQueue::ret_ = 0;
CMsgQueue* CMsgQueue::qry_ = 0;

CMsgQueue*
CMsgQueue::ret()
{
  if (CMsgQueue::ret == 0)
  {
    CMsgQueue::ret = new CMsgQueue;
  }
  return ret_;
}

CMsgQueue*
CMsgQueue::qry()
{
  if (CMsgQueue::qry == 0)
  {
    CMsgQueue::qry = new CMsgQueue;
  }
  return qry_;
}

CMsgQueue::CMsgQueue()
{
    m_nMsgId = -1;
}

CMsgQueue::~CMsgQueue()
{
}

/*
 * 创建消息队列
 */
int CMsgQueue::Open(const char *szPathName)
{
    key_t key = ftok(szPathName, 1);
    if (key < 0)
    {
        perror("ftok error\n");
        
        return -1;
    }
    
    printf("key::%ld\n", key);
    
    m_nMsgId = msgget(key, IPC_CREAT|0660);
    if (m_nMsgId < 0)
    {
        printf("msgget error\n");
        return -1;
    }
    
    return 0;
}


int CMsgQueue::IsOpen() const
{
    return (m_nMsgId == -1)? 0:1;
}


int CMsgQueue::PutMsg(const char *szBuff, int nMsgLen, long lMsgType)
{
    if (nMsgLen < 0)
        return -1;
    
    memset(&m_msg, 0, sizeof(m_msg));
    
    m_msg.mtype = lMsgType;
    memcpy(m_msg.mtext, szBuff, nMsgLen);
    
    /* 无阻塞
     */
    return msgsnd(m_nMsgId, &m_msg, nMsgLen, IPC_NOWAIT);
}


int CMsgQueue::GetMsg(char *szBuff, int nSize, long lMsgType)
{
    if (szBuff == NULL)
        return -1;
    
    /* 指定消息类型，无阻塞读取
     * 如果消息超过最大缓存大小，则消息被截断。
     */
    int nMsgLen = msgrcv(m_nMsgId, &m_msg, nSize, lMsgType, IPC_NOWAIT|MSG_NOERROR);
    if (nMsgLen > 0)
    {
        memcpy(szBuff, m_msg.mtext, nMsgLen);
    }
    
    return nMsgLen;
}


/*
 * 读取消息队列中的消息个数
 */
int CMsgQueue::GetQueueDepth() const
{
    struct msqid_ds buf;
    
    if (msgctl(m_nMsgId, IPC_STAT, &buf) < 0)
    {
        return -1;
    }
    
    /* 消息个数
     */
    return buf.msg_qnum;
}



/*
 * 关闭消息队列
 */
int CMsgQueue::Close()
{
    struct msqid_ds buf;
    
    return msgctl(m_nMsgId, IPC_RMID, &buf);
}


/* TEST FUNCTION
 */
int main()
{
    CMsgQueue mq;
    
    if (mq.Open("/tmp/.emppsimulator") < 0)
    {
        return -1;
    }
    
    char szBuff[512];
    long lMsgType = 1000;
    
    int nMsgLen = 0;
    
    for (int i=0; i<10; i++)
    {
        nMsgLen = sprintf(szBuff, "hello, msg::%d", i);
        lMsgType ++;
           
        mq.PutMsg(szBuff, nMsgLen, lMsgType);
    }
    
    printf("Queue Depth::%d\n", mq.GetQueueDepth());
    
    lMsgType = 1000;
    for (int i=0; i<15; i++)
    {
        lMsgType ++;
        
        nMsgLen = mq.GetMsg(szBuff, sizeof(szBuff), lMsgType);
        szBuff[nMsgLen] = '\0';
        
        printf("Msg::%s\n", szBuff);
    }
    
    sleep(30);
    
    mq.Close();
    
    return 0;
}

