package email

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "dj <XXX@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{"XXX@qq.com"}
	//设置抄送如果抄送多人逗号隔开
	e.Cc = []string{"XXX@qq.com", "XXX@qq.com"}
	//设置秘密抄送
	e.Bcc = []string{"XXX@qq.com"}
	//设置主题
	e.Subject = "这是主题"
	//设置文件发送的内容
	e.Text = []byte("www.topgoer.com是个不错的go语言中文文档")
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "你的邮箱账号", "这块是你的授权码", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
