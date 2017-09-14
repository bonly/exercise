/*
auth: bonly
create: 2016.4.21
*/

package main 

import (
"fmt"
"log"
"encoding/json"
"net/http"
"io/ioutil"
"os/exec"
"net/url"
"flag"
"os"
// "syscall"
)

type GitLab struct{
  Object_kind string `json:"object_kind"`;
  Before string `json:"before"`;
  After string `json:"after"`;
  Repository struct{
      Name string `json:"name"`;
      Url string `json:"url"`;
      Description string `json:"description"`;
      Git_http_url string `json:"git_http_url"`;
      Git_ssh_url string `json:"git_ssh_url"`;
  } `json:"repository"`;
  Project struct{
  	  Name string `json:"name"`;
  	  Namespace string `json:"namespace"`;
  	  Path_with_namespace string `json:"Path_with_namespace"`;
  	  Default_branch string `json:"default_branch"`;
  } `json:"project"`;
};

func myUsage(){
	fmt.Println("使用手册:");
	fmt.Printf("%s:\n", os.Args[0]);
	flag.PrintDefaults();
}

var ssh_srv *string = flag.String("s", "a243", "addr for sending command to ssh server");
var git_pre *string = flag.String("p", "/home/gitlab_docker/gitlab/repositories/", "git prefix dir in the ssh server");

func main(){
	flag.Usage = myUsage;

	help := flag.Bool("h", false, "Help");

	flag.Parse();

	if *help{
		flag.Usage();
		return;
	}

	http.HandleFunc("/git", Git_even);

	err := http.ListenAndServe(":8989", nil);
	if err != nil{
		log.Println(err);
	}
}

func Git_even(rw http.ResponseWriter, qry *http.Request){
	log.Println("======  recv git qry  =======");
	defer log.Println("=====  end git qry  ======");

	fmt.Printf("qry: %+v\n", *qry);

	rw.Header().Set("Access-Control-Allow-Origin", "*");
	rw.Header().Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");

	body, err := ioutil.ReadAll(qry.Body);
	if len(body) <= 0 || err != nil{
		log.Println("body err: ", err);
		return;
	}

	var qry_dat GitLab;
	err = json.Unmarshal(body, &qry_dat);
	if err != nil{
		log.Println(err);
		return;
	}

	fmt.Printf("qry: %+v\n", qry_dat);

	mp, err := url.ParseQuery(qry.URL.RawQuery);
	if err != nil{
		log.Println(err);
		return;
	}
	// fmt.Printf("%+v\n", mp);
	fmt.Printf("workpath: %s\n", mp["workpath"][0]);

	if qry_dat.Object_kind != "push"{
		return;
	}
	rm(mp["workpath"][0]);
	git_checkout_laster(qry_dat.Repository.Url, qry_dat.Project.Default_branch, mp["workpath"][0]);
	//git_temp_tag(qry_dat.After, "/home/gitlab_docker/gitlab/repositories/bonly/xb_room.git");
	// git_checkout(qry_dat.Repository.Url, qry_dat.After, mp["workpath"][0]);
}

func git_temp_tag(tag string, git_dir string){
	cmd_argv := []string{"/usr/local/bin/git", "tag", "temporary", tag};
	cmd := exec.Command("sudo", cmd_argv...);
	cmd.Env = []string{"GIT_DIR=" + git_dir};
	// cmd.SysProcAttr = &syscall.SysProcAttr{};
	// cmd.SysProcAttr.Credential = &syscall.Credential{Uid:0, Gid: 0};
	// cmd.Dir = git_dir;
    cout, err := cmd.CombinedOutput();
    fmt.Println(string(cout));
    if err != nil{
    	fmt.Printf("make temp tag: %v", err);
    	return;
    }
}

func git_checkout_laster(git string, branch string, path string) error{
	cmd_argv := []string{"clone", "-b", branch, "--depth", "1", git, path};
	cmd := exec.Command("git", cmd_argv...);
	cmd.Env = []string{"GIT_DIR=" + git};

	out, err := cmd.CombinedOutput();
	
	log.Println(cmd);

	log.Println(string(out));

	if err != nil{
		log.Println("err: ", err);
		return err;	
	}	
	return nil;
}

func git_checkout(git string, tag string, path string) error{
	cmd_argv := []string{"clone", "-b", tag, "--depth", "1", git, path};
	cmd := exec.Command("git", cmd_argv...);
	cmd.Env = []string{"GIT_DIR=" + git};

	out, err := cmd.CombinedOutput();
	
	log.Println(cmd);

	log.Println(string(out));

	if err != nil{
		log.Println("err: ", err);
		return err;	
	}	
	return nil;
}

func rm(dir string)(err error){
	if dir == "/"{
		log.Println("不要删除root！！！！");
		return nil;
	}
	cmd_argv := []string{"rm", "-rf", dir};
	cmd := exec.Command("sudo", cmd_argv...);

	out, err := cmd.CombinedOutput();
	
	log.Println(cmd);

	log.Println(string(out));

	if err != nil{
		log.Println("err: ", err);
		return err;	
	}	
	return nil;
}

/*
{
    "object_kind": "push",
    "before": "ae3a36eab7b87d7d78e067c543ac647696e66f24",
    "after": "f678374944e5081a8efeacd4ca2707ecd3993f8b",
    "ref": "refs/heads/master",
    "checkout_sha": "f678374944e5081a8efeacd4ca2707ecd3993f8b",
    "message": null,
    "user_id": 4,
    "user_name": "bonly",
    "user_email": "hejb@xbed.com.cn",
    "project_id": 25,
    "repository": {
        "name": "xb_room",
        "url": "ssh://git@120.25.106.243:10022/bonly/xb_room.git",
        "description": "房态查询中间层服务",
        "homepage": "http://120.25.106.243:10080/bonly/xb_room",
        "git_http_url": "http://120.25.106.243:10080/bonly/xb_room.git",
        "git_ssh_url": "ssh://git@120.25.106.243:10022/bonly/xb_room.git",
        "visibility_level": 10
    },
    "commits": [{
        "id": "f678374944e5081a8efeacd4ca2707ecd3993f8b",
        "message": "for test\n",
        "timestamp": "2016-04-21T11:46:50+08:00",
        "url": "http://120.25.106.243:10080/bonly/xb_room/commit/f678374944e5081a8efeacd4ca2707ecd3993f8b",
        "author": {
            "name": "bonly",
            "email": "bonly@163.com"
        },
        "added": [],
        "modified": [
            "makefile"
        ],
        "removed": []
    }, {
        "id": "8ab3091f993cfaf8531596f354168acbfe9107ae",
        "message": "查询房间标识为可配置，默认为1\n",
        "timestamp": "2016-04-21T11:19:28+08:00",
        "url": "http://120.25.106.243:10080/bonly/xb_room/commit/8ab3091f993cfaf8531596f354168acbfe9107ae",
        "author": {
            "name": "bonly",
            "email": "bonly@163.com"
        },
        "added": [
            "bin/config.json.sample"
        ],
        "modified": [
            "src/app/configure.go",
            "src/bus/db_op.go",
            "src/bus/exp_room_stat.go"
        ],
        "removed": []
    }, {
        "id": "ae3a36eab7b87d7d78e067c543ac647696e66f24",
        "message": "for test\n",
        "timestamp": "2016-04-19T18:50:04+08:00",
        "url": "http://120.25.106.243:10080/bonly/xb_room/commit/ae3a36eab7b87d7d78e067c543ac647696e66f24",
        "author": {
            "name": "bonly",
            "email": "bonly@163.com"
        },
        "added": [],
        "modified": [
            "makefile"
        ],
        "removed": []
    }],
    "total_commits_count": 3
}
*/
