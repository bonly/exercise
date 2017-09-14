/*
auth: bonly
create: 2016.9.20
desc: 公共全局配置
*/
package config

import(
"flag"
"os"
)

var Box_srv = flag.String("box_srv", "0.0.0.0:5020", "对盒子的服务地址及端口");
var Oms_srv = flag.String("oms_srv", "0.0.0.0:5010", "对OMS的服务地址及端口");
var Memprf = flag.String("memprf", "", "内存分析");
var Cpuprf = flag.String("cpuprf", "", "CPU分析");

var Run = true;
var Run_chn chan bool;

var MemFile *os.File;
var CpuFile *os.File;