package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/smtp"
	"os/exec"
	"strings"
	"time"
)

var TargetUrl = "https://www.douyu.com/606118"
var chooice string

func main() {
	resp, err := http.Get(TargetUrl)
	if err != nil {
		// 异常
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 异常
	}

	//fmt.Println(string(body))
	//fmt.Println(strings.Contains(string(body),"show_status = 2"))
	if strings.Contains(string(body), "show_status = 1") == false {
		fmt.Print("当前直播间未开播\n5秒后退出程序")
		time.Sleep(5 * time.Second)
	} else if strings.Contains(string(body), "show_status = 1") == true {
		fmt.Print("大司马直播中\t请输入：\ngo直接进入捞马房间\t\tno关闭程序\n")
		fmt.Scanf("%s", &chooice)
		switch chooice {
		case "go":
			cmd := exec.Command("cmd.exe", "/c", "start microsoft-edge:https://www.douyu.com/606118")
			_ = cmd.Run()
		case "no":

		}

	}
}

