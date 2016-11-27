package main

import(
"github.com/golang/glog"
"flag"
)

func main(){
	flag.Parse();
	flag.Set("alsologtostderr", "true");
	glog.Info("only info");
	glog.Error("only Error");
	glog.Waring("only Waring");
	glog.V(0).Info("info 0");
	glog.V(1).Info("info 1");
	glog.V(2).Info("info 2");
	glog.V(3).Info("info 3");
	glog.V(4).Info("info 4");
	glog.V(5).Info("info 5");	
	glog.InfoDepth(0,"depth");
	glog.InfoDepth(1,"depth 1");
	glog.InfoDepth(2,"depth 2");	
	glog.InfoDepth(3,"depth 3");	
	glog.InfoDepth(4,"depth 4");	
	// glog.V(0).InfoDepth("d 0");
	// glog.V(1).InfoDepth("d 1");
	// glog.V(2).InfoDepth("d 2");
	// glog.V(3).InfoDepth("d 3");		
}
