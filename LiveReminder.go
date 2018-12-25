package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

var TargetUrl = "https://www.douyu.com/606118"
var choice string

func main() {
	resp, _ := http.Get(TargetUrl)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if strings.Contains(string(body), "show_status = 1") == false {
		fmt.Print("当前直播间未开播\n5秒后退出程序")
		time.Sleep(5 * time.Second)
	} else if strings.Contains(string(body), "show_status = 1") == true {
		fmt.Print("大司马直播中\t请输入：\ngo直接进入捞马房间\t\t任意键关闭程序\n")
		fmt.Scanf("%s", &choice)
		switch choice {
		case "go":
			cmd := exec.Command("cmd.exe", "/c", "start microsoft-edge:https://www.douyu.com/606118")
			_ = cmd.Run()
		}

	}
}
