// +build linux
// Input device event monitor.
package main

import (
	"errors"
	"fmt"
	"github.com/gvalkov/golang-evdev"
	"os"
	"strings"
	"uinput"
	"time"
)

const (
	usage       = "usage: evtest <device> [<type> <value>]"
	device_glob = "/dev/input/event*"
	tm = 8;
)

var run = true;

func select_device() (*evdev.InputDevice, error) {
	devices, _ := evdev.ListInputDevices(device_glob)

	lines := make([]string, 0)
	max := 0
	if len(devices) > 0 {
		for i := range devices {
			dev := devices[i]
			str := fmt.Sprintf("%-3d %-20s %-35s %s", i, dev.Fn, dev.Name, dev.Phys)
			if len(str) > max {
				max = len(str)
			}
			lines = append(lines, str)
		}
		fmt.Printf("%-3s %-20s %-35s %s\n", "ID", "Device", "Name", "Phys")
		fmt.Printf(strings.Repeat("-", max) + "\n")
		fmt.Printf(strings.Join(lines, "\n") + "\n")

		var choice int
		choice_max := len(lines) - 1

	ReadChoice:
		fmt.Printf("Select device [0-%d]: ", choice_max)
		_, err := fmt.Scan(&choice)
		if err != nil || choice > choice_max || choice < 0 {
			goto ReadChoice
		}

		return devices[choice], nil
	}

	errmsg := fmt.Sprintf("no accessible input devices found by %s", device_glob)
	return nil, errors.New(errmsg)
}

func format_event(ev *evdev.InputEvent) (string,bool) {
	var res, f, code_name string
    want := false;

	code := int(ev.Code)
	etype := int(ev.Type)

	switch ev.Type {
	case evdev.EV_SYN:
		if ev.Code == evdev.SYN_MT_REPORT {
			f = "time %d.%-8d +++++++++ %s ++++++++"
		} else {
			f = "time %d.%-8d --------- %s --------"
		}
		return fmt.Sprintf(f, ev.Time.Sec, ev.Time.Usec, evdev.SYN[code]), false;
	case evdev.EV_KEY:
		val, haskey := evdev.KEY[code]
		if haskey {
			code_name = val
		} else {
			val, haskey := evdev.BTN[code]
			if haskey {
				code_name = val
			} else {
				code_name = "?"
			}
		}
		want = true;
	default:
		m, haskey := evdev.ByEventType[etype]
		if haskey {
			code_name = m[code]
		} else {
			code_name = "?"
		}
	}

	evfmt := "time %d.%-8d type %d (%s), code %-3d (%s), value %d"
	res = fmt.Sprintf(evfmt, ev.Time.Sec, ev.Time.Usec, etype,
		evdev.EV[int(ev.Type)], ev.Code, code_name, ev.Value)

	return res, want;
}


var vk uinput.VKeyboard;
var alt bool = false;

func main() {
	var dev *evdev.InputDevice
	var events []evdev.InputEvent
	var err error

	switch len(os.Args) {
	case 1:
		dev, err = select_device()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case 2:
		dev, err = evdev.Open(os.Args[1])
		if err != nil {
			fmt.Printf("unable to open input device: %s\n", os.Args[1])
			os.Exit(1)
		}
	default:
		fmt.Printf(usage + "\n")
		os.Exit(1)
	}

	info := fmt.Sprintf("bus 0x%04x, vendor 0x%04x, product 0x%04x, version 0x%04x",
		dev.Bustype, dev.Vendor, dev.Product, dev.Version)

	repeat_info := dev.GetRepeatRate()

	fmt.Printf("Evdev protocol version: %d\n", dev.EvdevVersion)
	fmt.Printf("Device name: %s\n", dev.Name)
	fmt.Printf("Device info: %s\n", info)
	fmt.Printf("Repeat settings: repeat %d. delay %d\n", repeat_info[0], repeat_info[1])
	fmt.Printf("Device capabilities:\n")

	fmt.Printf("Listening for events ...\n")

	// vk := uinput.VKeyboard{};
	err = vk.Create("/dev/uinput");

	if err != nil {
		fmt.Printf("Failed to create the virtual keyboard. Last error was: %s\n", err);
		return;
	};	
	defer func(){
			err = vk.Close()

			if err != nil {
				fmt.Printf("Failed to close device. Last error was: %s\n", err)
			}
	}();

	for {
		events, err = dev.Read();
		for i := range events {
			proc(&events[i]);
		// 	str, want := format_event(&events[i])
		// 	if want{
		// 		fmt.Println(str);
		// 		err = vk.SendKeyPress(uinput.KEY_1)

		// 		if err != nil {
		// 			fmt.Printf("Failed to send key event. Last error was: %s\n", err)
		// 		}

		// 		err = vk.SendKeyRelease(uinput.KEY_1)

		// 		if err != nil {
		// 			fmt.Printf("Failed to send key event. Last error was: %s\n", err)
		// 		}
		// 	}
		}
	}
}

func proc(ev *evdev.InputEvent){
	code := int(ev.Code);
	var code_name string;
	switch ev.Type {
		case evdev.EV_KEY:{
			val, haskey := evdev.KEY[code];
			if haskey {
				code_name = val
			} else {
				val, haskey := evdev.BTN[code]
				if haskey {
					code_name = val
				} else {
					code_name = "?"
				}
			}
			// fmt.Println(code_name);
			if ev.Value == 0 { //放键
				// vk.SendKeyRelease(evdev.EV_KEY);
				if alt==false {
					if run {
						carl(code_name);
					}
				}else{ //按着alt时
					switch_run(code_name);
				}	
				if code_name == "KEY_LEFTALT"{ //放开alt
					// fmt.Println("alt=false");
					alt = false;
				}

				break;
			}else{ //按键
				if code_name == "KEY_LEFTALT"{ //按下alt时标识
					// fmt.Println("alt=true");
					alt = true;
					break;
				}
			}

			break;
		}
	}
}

func switch_run(code_name string){
	if code_name == "KEY_P"{  //pause
			run = !run;
	}
}

func carl(code_name string){
	switch code_name{
		case "KEY_E":{
			// attack();
			break;
		}
		case "KEY_V":{
			技能V();
			break;
		}
		case "KEY_B":{
			技能B();
			break;
		}
		case "KEY_1":{
			飓风();
			break;
		}	
		case "KEY_2":{
			磁冲();
			break;
		}
		case "KEY_T":{
			隐形();
			break;
		}
		case "KEY_4":{
			天火();
			break;
		}
		case "KEY_5":{
			精灵();
			break;
		}
		case "KEY_Q":{
			灵迅();
			break;
		}
		case "KEY_W":{
			急冷();
			break;
		}
		case "KEY_3":{
			声波();
			break;
		}
		case "KEY_G":{
			陨石();
			break;
		}
		case "KEY_H":{
			冰墙();
			break;
		}
	}

}

func attack(){
	vk.SendKeyPress(uinput.KEY_R);
	vk.SendKeyPress(uinput.KEY_E);
	// if err := vk.SendKeyPress(uinput.KEY_E); err != nil{
	// 	fmt.Printf("Failed to send key event. Last error was: %s\n", err);
	// }

	// if err := vk.SendKeyRelease(uinput.KEY_E); err != nil{
	// 	fmt.Printf("Failed to send key event. Last error was: %s\n", err)
	// }
}

func press_key(key int){
	if err := vk.SendKeyPress(key); err != nil{
		fmt.Printf("Failed to send key event. Last error was: %s\n", err);
	}

	if err := vk.SendKeyRelease(key); err != nil{
		fmt.Printf("Failed to send key event. Last error was: %s\n", err)
	}	
}

func 技能V(){
	vk.SendKeyPress(uinput.KEY_LEFTALT);
	press_key(uinput.KEY_V);
	vk.SendKeyRelease(uinput.KEY_LEFTALT);
}

func 技能B(){
	vk.SendKeyPress(uinput.KEY_LEFTALT);
	press_key(uinput.KEY_B);
	vk.SendKeyRelease(uinput.KEY_LEFTALT);
}


func 急冷(){
	vk.SendKeyPress(uinput.KEY_LEFTALT);
	press_key(uinput.KEY_Q);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_Q);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_Q);
	time.Sleep(tm * time.Millisecond);
	vk.SendKeyRelease(uinput.KEY_LEFTALT);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_SPACE);
}
func 天火(){
	// time.Sleep(1 * time.Millisecond);
	vk.SendKeyPress(uinput.KEY_LEFTALT);
	press_key(uinput.KEY_T);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_T);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_T);
	time.Sleep(tm * time.Millisecond);
	vk.SendKeyRelease(uinput.KEY_LEFTALT);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_SPACE);
}
func 飓风(){
	// time.Sleep(1 * time.Millisecond);
	vk.SendKeyPress(uinput.KEY_LEFTALT);
	press_key(uinput.KEY_W);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_W);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_Q);
	time.Sleep(tm * time.Millisecond);
	vk.SendKeyRelease(uinput.KEY_LEFTALT);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_SPACE);
}

func 磁冲(){
	// time.Sleep(1 * time.Millisecond);
	vk.SendKeyPress(uinput.KEY_LEFTALT);
	press_key(uinput.KEY_W);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_W);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_W);
	time.Sleep(tm * time.Millisecond);
	vk.SendKeyRelease(uinput.KEY_LEFTALT);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_SPACE);
}

func 隐形(){
	// time.Sleep(1 * time.Millisecond);
	vk.SendKeyPress(uinput.KEY_LEFTALT);
	press_key(uinput.KEY_Q);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_Q);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_W);
	time.Sleep(tm * time.Millisecond);
	vk.SendKeyRelease(uinput.KEY_LEFTALT);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_SPACE);
}

func 冰墙(){
	// time.Sleep(1 * time.Millisecond);
	vk.SendKeyPress(uinput.KEY_LEFTALT);
	press_key(uinput.KEY_Q);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_Q);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_T);
	time.Sleep(tm * time.Millisecond);
	vk.SendKeyRelease(uinput.KEY_LEFTALT);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_SPACE);
}

func 灵迅(){
	// time.Sleep(1 * time.Millisecond);
	vk.SendKeyPress(uinput.KEY_LEFTALT);
	press_key(uinput.KEY_W);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_W);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_T);
	time.Sleep(tm * time.Millisecond);
	vk.SendKeyRelease(uinput.KEY_LEFTALT);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_SPACE);
}

func 精灵(){
	// time.Sleep(1 * time.Millisecond);
	vk.SendKeyPress(uinput.KEY_LEFTALT);
	press_key(uinput.KEY_T);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_T);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_Q);
	time.Sleep(tm * time.Millisecond);
	vk.SendKeyRelease(uinput.KEY_LEFTALT);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_SPACE);
}

func 陨石(){
	vk.SendKeyPress(uinput.KEY_LEFTALT);
	press_key(uinput.KEY_W);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_T);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_T);
	time.Sleep(tm * time.Millisecond);
	vk.SendKeyRelease(uinput.KEY_LEFTALT);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_SPACE);
}

func 声波(){
	vk.SendKeyPress(uinput.KEY_LEFTALT);
	press_key(uinput.KEY_Q);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_W);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_T);
	time.Sleep(tm * time.Millisecond);
	vk.SendKeyRelease(uinput.KEY_LEFTALT);
	time.Sleep(tm * time.Millisecond);
	press_key(uinput.KEY_SPACE);
}
