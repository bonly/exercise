# OMS端接口

POST请求：

http://120.25.106.243:5010/cmd

## 增加用户

<pre><code>
{
"Name":"Add_OpenUser_REQ",
"MidId":"00000001",
"LockId":"1",
"CardType":"3",
"CardData":"132132",
"BeginTime":"2016-09-21 00:00:00",
"EndTime":"2016-09-29 23:59:43"
}
</code></pre>

## 删除用户

<pre><code>
{
"Name":"Delete_OpenUser_REQ",
"MidId":"00000001",
"LockId":"1",
"CardType":"3",
"CardData":"132132"
}
</code></pre>

## 时间校正

<pre><code>
{
"Name":"Time_REQ",
"MidId":"00000001",
"LockId":"1"
}
</code></pre>

## 远程开门

<pre><code>
{
"Name":"OpenLock_REQ",
"MidId":"00000001",
"LockId":"1"
}
</code></pre>

## 门锁状态查询

<pre><code>
{
"Name":"Lock_stat_REQ",
"MidId":"0000001",
"LockId":"1"
}
</code></pre>

## 清空所有用户

<pre><code>
{
"Name":"User_Clean_REQ",
"MidId":"00000001",
"LockId":"1"
}
</code></pre>

# 错误代码

<pre><code>
	OK = iota;                  // 0
	Fail;                       // 1
	ERR_UNKNOW;                 // 2:未知错误
	ERR_PROTO;					// 3:协议错误
	ERR_BOX_NOT_ONLINE;         // 4:盒子不在线
	ERR_BOX_BUSY;               // 5:门锁没处理完前一任务
	ERR_BOX_TIMEOUT;			// 6:接收结果超时
	ERR_BOX_NOT_RES;			// 7:门锁无应答
	ERR_BOX_RES_CHAN_CLOSED;    // 8:结果通道收到结束标识
</code></pre>