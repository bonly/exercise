package main;
import (
  "os"
  "fmt"
);

func chg_map(in *(map[string]string)){ //it is ok
	(*in)["first"] = "in first";
	(*in)["second"] = "my second";
}

func oth_map()(out *map[string]string){  //failed
	*out = make(map[string]string);  //pass pointer, make memory in func, memory had del outside of func
	(*out)["3th"] = "ok in 3";
	(*out)["4th"] = "right in 4";
	return out;
}

func oth_map_obj()(out map[string]string){  //OK
	out = make(map[string]string);  
	(out)["3th"] = "ok in 3";
	(out)["4th"] = "right in 4";
	return out;  //copy val to return
}

func oth_map_pt(in *map[string]string)(*map[string]string){ // OK: pointer is ref other
	                                                          // failed: just def pointer not ref, will be nil
	*in = make(map[string]string);  
	(*in)["3th"] = "ok in 3";
	(*in)["4th"] = "right in 4";
	return in;  
}

func p_map(in *map[string]string){  //failed
	var omap map[string]string;
  omap = make(map[string]string);  //it is del outside of func
	in = &omap;
	(*in)["5th"] = "5th";
	(*in)["6th"] = "6th";
}

func obj_map(in map[string]string){ //ok: but not change outside value
	in = make(map[string]string);
	in["5th"] = "5th";
	in["6th"] = "6th";
}

func nil_ptr(in *int){
	if in == nil{
		fmt.Println(" nil param\n");
	}
}

func main(){
	amp := make(map[string]string);
	chg_map(&amp);
	
	for key,val := range amp {
		fmt.Fprintf( os.Stdout, "%s = %s\n", key, val);
	}	
	
	/*
	oth := oth_map(); 
	//for key,val := range *(oth_map()){ //failed for nil pointer
	for key,val := range *oth{ //var in for is dif with outside, memory had distruct out of func
		fmt.Fprintf( os.Stdout, "%s = %s\n", key, val);
	}
	*/
	
	oth := oth_map_obj(); 
	for key,val := range oth{ 
		fmt.Fprintf( os.Stdout, "%s = %s\n", key, val);
	}
		
	/*
	var obj *map[string]string;
	obj_map(*obj);  //invalid memory address or nil pointer dereference
	for key,val := range *obj{
		fmt.Fprintf(os.Stdout, "%s = %s\n", key, val);
	}
	*/

	var obj1 *map[string]string;
	rel := make(map[string]string);
	obj1 = &rel;
	obj_map(*obj1);  //cant change outside value
	for key,val := range *obj1{
		fmt.Fprintf(os.Stdout, "%s = %s\n", key, val);
	}	
	
	/*
	var inout *map[string]string;
	oth_map_pt(inout); //cant pass nil pointer?
	for key, val := range *inout{
		fmt.Fprintf(os.Stdout, "%s = %s\n", key, val);
	}	
	*/
	
	var inout map[string]string;
	oth_map_pt(&inout); 
	for key, val := range inout{
		fmt.Fprintf(os.Stdout, "oth_map_pt %s = %s\n", key, val);
	}		
	
	var aint *int;
	nil_ptr(aint);
	
	/*
	var ap_map *map[string]string;
	p_map(ap_map);
	for key, val := range *ap_map{ //had been del
		fmt.Fprintf(os.Stdout, "%s = %s\n", key, val);
	}			
	*/
}
