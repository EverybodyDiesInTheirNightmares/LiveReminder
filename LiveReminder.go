package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/smtp"
	"os"
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
		fmt.Print("捞马没开播\n5秒后退出程序")
		time.Sleep(5 * time.Second)
		os.Exit(1)
	} else if strings.Contains(string(body), "show_status = 1") == true {
		fmt.Print("大司马直播中\t请输入：\ngo直接进入捞马房间\t\t任意键关闭程序\n")
		fmt.Scanf("%s", &choice)
		switch choice {
		case "go":
			cmd := exec.Command("cmd.exe", "/c", "start microsoft-edge:https://www.douyu.com/606118")
			_ = cmd.Run()
		}
		os.Exit(2)

	}
}

func SendMail() {
	_, err := net.InterfaceAddrs()

	auth := smtp.PlainAuth("", "你的邮箱账户", "SMTP授权码", "SMTP服务器地址")
	to := []string{"收件人邮箱（你的邮箱账户）"}
	nickname := "开播提醒"
	user := "发件人邮箱（你的邮箱账户）"
	subject := "LiveReminder"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := "大司马已开播" + TargetUrl
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	fmt.Print("邮件发送成功")
	err = smtp.SendMail("SMTP服务器+非SSL协议端口号 如：smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("邮件发送出错: %v", err)
	}
}
