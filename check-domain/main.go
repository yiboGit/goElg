package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os/exec"
	"time"
)

func SendMail(to []string, subject, body string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "tech@epeijing.cn", "Eglass2018", "smtp.exmail.qq.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	msg := []byte("To: hb@epeijing.cn\r\n" +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"\r\n" +
		fmt.Sprintf("%s\r\n", body))
	err := smtp.SendMail("smtp.exmail.qq.com:25", auth, "tech@epeijing.cn", to, msg)
	return err
}

var oldDomain = ""
var oldSSH = ""

func checkDomain() {
	log.Println("check domain start")
	domain, error := exec.Command("/bin/bash", "-c", `journalctl -u ng-note | grep 'established' | tail -n 1 | awk '{ print $13}'`).Output()
	if error != nil {
		SendMail([]string{"hb@epeijing.cn"}, "获取domain失败", "")
	}
	dm := string(domain)
	if oldDomain == "" || dm != oldDomain {
		log.Printf("domain change to %s ", domain)
		SendMail([]string{"hb@epeijing.cn"}, "domain change ", dm)
		oldDomain = dm
	}
	domainSSH, error := exec.Command("/bin/bash", "-c", `journalctl -e -u ngrok | grep 'established' | tail -n 1 | awk '{ print $13}'`).Output()
	dm = string(domainSSH)
	if error != nil {
		SendMail([]string{"hb@epeijing.cn"}, "获取domain失败", "")
	}
	if oldSSH == "" || dm != oldSSH {
		log.Printf("ssh domain change to %s ", domain)
		SendMail([]string{"hb@epeijing.cn"}, "ssh port change ", dm)
		oldSSH = dm
	}
}

func main() {
	checkDomain()
	tick := time.Tick(time.Minute)
	for {
		select {
		case <-tick:
			checkDomain()
		}
	}
}
