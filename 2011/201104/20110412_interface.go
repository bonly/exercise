package main
import "fmt"

type USB interface{
	Name() string;
	Connect();
};

type PhoneConnect struct{
	name string;
};

func (p PhoneConnect) Name() string{
	return p.name;
};

func (p PhoneConnect) Connect() {
	fmt.Println("Connect: ", p.Name());
};

func main(){
   a := PhoneConnect{"iPhone"};
   a.Connect();
   
   var b USB;
   b = PhoneConnect{"Android"};
   b.Connect();
   
   Disconnect (b);
   AllDisconnect(a);
   
   //c := PhoneConnect{"symban"};
   //c.Connect();
}

func Disconnect (usb USB){
	if pc, ok := usb.(PhoneConnect); ok{
		fmt.Println("Disconnect: ", pc.name);
		return;
	}
	fmt.Println("unknown device.");
}

func AllDisconnect (usb interface{}){
	switch v:=usb.(type){
		case PhoneConnect:
				fmt.Println("Disconnect: ", v.name);
		    break;
		default:
		    fmt.Println("unknown device.");
	}
}