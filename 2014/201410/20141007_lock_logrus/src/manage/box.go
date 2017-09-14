/*
auth: bonly
create: 2016.9.19
desc: 门锁管理盒子
*/
package manage

import (
// "fmt"
"net"
"proto"
"sync"
log "logrus"
)

type Box struct{
	ID uint32;  //盒子ID
	Conn net.Conn; //盒子socket
	Task chan func()([]byte,int,[]byte); //任务通道
	Result chan interface{};//结果通道
	Cur_task interface{}; //当前任务
};

type Work_map map[uint32]*Box;
type Conn_map map[net.Conn]*Box;

type Worker struct{
	sync.RWMutex;
	Work_map;
	Conn_map;
}

var work *Worker;
func Works()(*Worker){
	if work == nil{
		work = &Worker{};
		work.Work_map = make(map[uint32]*Box);
		work.Conn_map = make(map[net.Conn]*Box);
	}
	return work;
}


func (this *Worker)Health(ID uint32, Conn net.Conn){
	defer func(){
		if err := recover(); err != nil{
			log.Warnf("检查连接健康失败, %v\n", err);
		}
		return;
	}();
	if val, ok := this.Work_map[ID]; ok{
		if val.Conn == Conn{
			log.Infof("旧连接存在而且正常[%v]\n", val.Conn);
			return;
		}else{
			log.Warnf("新旧连接不一致，关闭旧连接，登记新连接\n");
			this.Lost(Conn);
		}
	}
	this.Push(ID, Conn);
	// log.Infof("%+v\n", *this);
}

func (this *Worker)Push(ID uint32, Conn net.Conn){
	log.Infof("登记新连接\n");
	var box Box;
	box.ID = ID;
	box.Conn = Conn;
	box.Task = make(chan func()([]byte, int, []byte)); //创建任务通道
	this.Lock();
	this.Work_map[ID] = &box;	//修改记录表
	this.Conn_map[Conn] = &box; //增加连接表
	this.Unlock();
	box.Run();
	log.Infof("登记新连接完成:\n%v\n", *this);
}

func (this *Worker)Lost(Conn net.Conn){
	defer func(){
		if err := recover(); err != nil{
			log.Warnf("注销失败, %v\n", err);
		}
		this.Unlock();
		return;
	}();
	log.Infof("开始注销连接\n");
	this.Lock();
	if _, ok := this.Conn_map[Conn]; ok{	
		// var wg sync.WaitGroup;
		// wg.Add(3);
		// go func(){
		{
			defer func(){
				if err := recover(); err != nil{
					log.Warnf("关闭任务通道失败, %v\n", err);
				}	
				// wg.Done();
				// return;
			}();
			close(this.Conn_map[Conn].Task); //关闭任务通道
			log.Infof("关闭任务通道\n");
		}
		// go func(){
		{
			defer func(){
				if err := recover(); err != nil{
					log.Warnf("关闭结果通道失败, %v\n", err);
				}	
				// wg.Done();
				// return;
			}();
			close(this.Conn_map[Conn].Result);//关闭结果通道
			log.Infof("关闭结果通道\n");
		}
		// go func(){
		{
			// defer wg.Done();
			Conn.Close(); //关闭连接
			log.Infof("关闭连接\n");
		}
		// wg.Wait();
		
		delete(this.Work_map, this.Conn_map[Conn].ID);//删除记录表
		delete(this.Conn_map, Conn);//删除连接
	}
	log.Infof("注销结束，管理器状态:\n %+v\n", *this);
}


func (this *Box)Run(){
	go func(){
		for {
			log.Infof("盒子[%X]就绪...\n", this.ID);
			select {
				case task, ok := <- this.Task:{
					if !ok {
						return;
					}
					log.Infof("Box[%X]recv task\n", this.ID);
					
					pack, size, cmd_name := task();
					//发送数据包
					this.Conn.Write(pack);
					log.Infof("发送数据包: \n");
					proto.Hex_Dump(pack, size);

					//修改当前任务命令
					this.Cur_task = cmd_name;
					log.Infof("当前的box: %+v\n", this);
				}
			}
		}
		log.Infof("结束接收任务\n");
	}();
}