package main 
import (
"github.com/ben-bensdevzone/uinput"
"fmt"
)

func main(){
	vk := uinput.VKeyboard{}
	err := vk.Create("/dev/uinput")

	if err != nil {
		fmt.Printf("Failed to create the virtual keyboard. Last error was: %s\n", err)
	}

	err = vk.SendKeyPress(uinput.KEY_1)

	if err != nil {
		fmt.Printf("Failed to send key event. Last error was: %s\n", err)
	}

	err = vk.SendKeyRelease(uinput.KEY_1)

	if err != nil {
		fmt.Printf("Failed to send key event. Last error was: %s\n", err)
	}

	err = vk.Close()

	if err != nil {
		fmt.Printf("Failed to close device. Last error was: %s\n", err)
	}
}