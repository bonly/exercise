package main

import (
"fmt"
"os/exec"
)

func git_get_tag(){
	cmd_argv := []string{"tag"};
	cmd := exec.Command("git", cmd_argv...);
	cmd.Env = []string{"GIT_DIR=/home/bonly/xbed/xb_door/.git"};

	out, err := cmd.Output();

	fmt.Println(string(out));

	if err != nil{
		fmt.Println("err: ", err);	
	}
}

func git_checkout(){
	cmd_argv := []string{"checkout", "-f", "1.6"};
	cmd := exec.Command("git", cmd_argv...);
	cmd.Env = []string{"GIT_DIR=/home/bonly/xbed/xb_door/.git",
					   "GIT_WORK_TREE=/tmp/xb_door"};

	out, err := cmd.Output();

	fmt.Println(string(out));

	if err != nil{
		fmt.Println("err: ", err);	
	}	
}

func make(){
	cmd_argv := []string{"-C", "/tmp/xb_door"};
	cmd := exec.Command("make", cmd_argv...);
	// cmd.Env = []string{"GIT_DIR=/home/bonly/xbed/xb_door/.git",
					   // "GIT_WORK_TREE=/tmp/xb_door"};

	out, err := cmd.Output();

	fmt.Println(string(out));

	if err != nil{
		fmt.Println("err: ", err);	
	}	
}

func cmd(){
	// cmd_argv := []string{"-C", "/tmp/xb_door"};
	cmd := exec.Command("make");
	cmd.Dir = "/tmp/xb_door";
	// cmd.Env = []string{"GIT_DIR=/home/bonly/xbed/xb_door/.git",
					   // "GIT_WORK_TREE=/tmp/xb_door"};

	out, err := cmd.Output();

	fmt.Println(string(out));

	if err != nil{
		fmt.Println("err: ", err);	
	}	
}

func main(){
	// git_get_tag();
	// git_checkout();
	// make();
	cmd();
}
