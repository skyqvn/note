package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func main() {
	identity := ""
	sender := "abc@abc.com"
	senderName := "abc"
	pwd := "123456"
	host := "r.r.com"
	port := "25"
	sendTo := []string{"r1@r1.com", "r2@r2.com", "r3@r3.com"}
	title := "how are you"
	body := "how are you???"
	contentType := "Content-type: text/html; charset=UTF-8"
	msg := []byte("To: " + strings.Join(sendTo, ",") + "\nFrom: " + senderName + "<" + sender + ">\nSubject: " + title + "\n" + contentType + "\n" + body + "\n")
	auth := smtp.PlainAuth(identity, sender, pwd, host)
	url := host + ":" + port
	err := smtp.SendMail(url, auth, sender, sendTo, msg)
	if err != nil {
		fmt.Println(err.Error())
	}
}
