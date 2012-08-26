#!/usr/bin/python
#coding:utf-8
'''在Linux系统中，可以用/proc/stat文件来计算cpu的利用率(详细的解释可参考：http: //www.linuxhowtos.org/System/procstat.htm)。这个文件包含了所有CPU活动的信息，该文件中的所有值都是从系统启动开始累计到当前时刻。
如：
[sailorhzr@builder ~]$ cat /proc/stat
cpu 432661 13295 86656 422145968 171474 233 5346
cpu0 123075 2462 23494 105543694 16586 0 4615
cpu1 111917 4124 23858 105503820 69697 123 371
cpu2 103164 3554 21530 105521167 64032 106 334
cpu3 94504 3153 17772 105577285 21158 4 24
intr 1065711094 1057275779 92 0 6 6 0 4 0 3527 0 0 0 70 0 20 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 7376958 0 0 0 0 0 0 0 1054602 0 0 0 0 0 0 0 30 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
ctxt 19067887
btime 1139187531
processes 270014
procs_running 1
procs_blocked 0

输出解释
CPU 以及CPU0、CPU1、CPU2、CPU3每行的每个参数意思（以第一行为例）为：
参数 解释
user (432661) 从系统启动开始累计到当前时刻，用户态的CPU时间（单位：jiffies） ，不包含 nice值为负进程。1jiffies=0.01秒
nice (13295) 从系统启动开始累计到当前时刻，nice值为负的进程所占用的CPU时间（单位：jiffies）
system (86656) 从系统启动开始累计到当前时刻，核心时间（单位：jiffies）
idle (42214596 从系统启动开始累计到当前时刻，除硬盘IO等待时间以外其它等待时间（单位：jiffies）
iowait (171474) 从系统启动开始累计到当前时刻，硬盘IO等待时间（单位：jiffies） ，
irq (233) 从系统启动开始累计到当前时刻，硬中断时间（单位：jiffies）
softirq (5346) 从系统启动开始累计到当前时刻，软中断时间（单位：jiffies）

CPU时间=user+system+nice+idle+iowait+irq+softirq

“intr”这行给出中断的信息，第一个为自系统启动以来，发生的所有的中断的次数；然后每个数对应一个特定的中断自系统启动以来所发生的次数。
“ctxt”给出了自系统启动以来CPU发生的上下文交换的次数。
“btime”给出了从系统启动到现在为止的时间（in seconds since the Unix epoch），单位为秒。
“processes (total_forks) 自系统启动以来所创建的任务的个数目。
“procs_running”：当前运行队列的任务的数目。
“procs_blocked”：当前被阻塞的任务的数目。

那么CPU利用率可以使用以下两个方法(后一种比较准确)。先取两个采样点，然后计算其差值：
cpu usage=(idle2-idle1)/(cpu2-cpu1)*100
cpu usage=[(user_2 +sys_2+nice_2) - (user_1 + sys_1+nice_1)]/(total_2 - total_1)*100


输出解释
CPU 以及CPU0、CPU1、CPU2、CPU3每行的每个参数意思（以第一行为例）为：
user (432661) 从系统启动开始累计到当前时刻，用户态的CPU时间（单位：jiffies） ，不包含 nice值为负进程。1jiffies=0.01秒
nice (13295) 从系统启动开始累计到当前时刻，nice值为负的进程所占用的CPU时间（单位：jiffies）
system (86656) 从系统启动开始累计到当前时刻，核心时间（单位：jiffies）
idle (422145968) 从系统启动开始累计到当前时刻，除硬盘IO等待时间以外其它等待时间（单位：jiffies）
iowait (171474) 从系统启动开始累计到当前时刻，硬盘IO等待时间（单位：jiffies） ，
irq (233) 从系统启动开始累计到当前时刻，硬中断时间（单位：jiffies）
softirq (5346) 从系统启动开始累计到当前时刻，软中断时间（单位：jiffies）

其他为：
“intr”    这行给出中断的信息，第一个为自系统启动以来，发生的所有的中断的次数；然后每个数对应一个特定的中断自系统启动以来所发生的次数。
“ctxt”    给出了自系统启动以来CPU发生的上下文交换的次数。
“btime”    给出了从系统启动到现在为止的时间，单位为秒。
"processes" (total_forks) 自系统启动以来所创建的任务的个数目。
“procs_running”    当前运行队列的任务的数目。
“procs_blocked”    当前被阻塞的任务的数目。

(以下每个值的2和1分别是2次获取的这个值，这里的cpu利用率就是这2者时间之间的平均值)
对于一个cpu的利用率，目前常见的计算方式是
user_pass = user2 - user1
system_pass = system2 - system1
idle_pass = idle2 - idle1
cpu利用率=(user_pass + system_pass)*100%/(user_pass + system_pass + idle_pass)
大家也会发现，这和top显示的cpu利用率不同，这又是为什么呢？
top上的cpu利用率，大致算法如下
CPU总时间2=user2+system2+nice2+idle2+iowait2+irq2+softirq2
CPU总时间1=user1+system1+nice1+idle1+iowait1+irq1+softirq1
用户cpu利用率 = user_pass * 100% / (CPU总时间2 - CPU总时间1)
内核cpu利用率 = system_pass * 100% / (CPU总时间2 - CPU总时间1)
总的cpu利用率= 用户cpu利用率 + 内核cpu利用率
这2者，谁优谁劣？
对于第一种，其计算出来的 = 100% - 系统空闲的百分率。
对于第二钟，这应该是真正意义上的 cpu利用率。
'''

import sys
import re
import time
#from scribe import scribe
from thrift.transport import TTransport, TSocket
from thrift.protocol import TBinaryProtocol

def read_cpu_usage():
    """Read the current system cpu usage from /proc/stat."""
    lines = open("/proc/stat").readlines()
    for line in lines:
        #print "l = %s" % line
        l = line.split()
        if len(l) < 5:
            continue
        if l[0].startswith('cpu'):
            return l;
    return {}

def sendlog(host,port,messa):
    #"""send log to scribe
    socket = TSocket.TSocket(host=host, port=port)
    transport = TTransport.TFramedTransport(socket)
    protocol = TBinaryProtocol.TBinaryProtocol(trans=transport, strictRead=False, strictWrite=False)
    #client = scribe.Client(iprot=protocol, oprot=protocol)
    transport.open()
    #log_entry = scribe.LogEntry(dict(category='SYSD', message=messa))
    result = client.Log(messages=[log_entry])
    transport.close()
    return result

if len(sys.argv) >= 2:
  host_port = sys.argv[1].split(':')
  host = host_port[0]
  if len(host_port) > 1:
    port = int(host_port[1])
  else:
    port = 1463
else:
  sys.exit('usage : py.test  host[:port]] ')

cpustr=read_cpu_usage()
down=True
#cpu usage=[(user_2 +sys_2+nice_2) - (user_1 + sys_1+nice_1)]/(total_2 - total_1)*100
usni1=long(cpustr[1])+long(cpustr[2])+long(cpustr[3])+long(cpustr[5])+long(cpustr[6])+long(cpustr[7])+long(cpustr[4])
usn1=long(cpustr[1])+long(cpustr[2])+long(cpustr[3])
#usni1=long(cpustr[1])+long(cpustr[2])+long(cpustr[3])+long(cpustr[4])
while(down):
       time.sleep(2)
       cpustr=read_cpu_usage()
       usni2=long(cpustr[1])+long(cpustr[2])+float(cpustr[3])+long(cpustr[5])+long(cpustr[6])+long(cpustr[7])+long(cpustr[4])
       usn2=long(cpustr[1])+long(cpustr[2])+long(cpustr[3])
       #usni2=long(cpustr[1])+long(cpustr[2])+float(cpustr[3])+long(cpustr[4])
       print usn2
       print usni2
       cpuper=(usn2-usn1)/(usni2-usni1)
       s="CPUTotal used percent =%.4f \r\n" % cpuper
       print s
       #sendlog(host,port,s)
       usn1=usn2
       usni1=usni2
