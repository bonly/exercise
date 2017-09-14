/*
auth: bonly
create: 2016.9.20
desc: 主程序
*/
package main 

import(
"flag"
"manage"
"oms"
)


func main(){
	flag.Parse();

	go func(){
		var pms oms.SRV;
		pms.Srv();
	}();

	// go func(){
		var box manage.SRV;
		box.Srv();
	// }
}
