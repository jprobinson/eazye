package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jprobinson/eazye"
)

func main() {
	info := eazye.MailboxInfo{
		"imap.gmail.com",
		true,
		"your.email.address@gmail.com",
		"pa$$w0rd",
		"INBOX",
	}

	getSince(info)
	getUnread(info)
	getAll(info)
}

func getUnread(info eazye.MailboxInfo) {
	emails, err := eazye.GetUnread(info, true, false)
	if err != nil {
		log.Print("get unread err: ", resp.Err.Error())
	}

	for _, email := range emails {
		fmt.Print(email)
	}

	return
}

func getSince(info eazye.MailboxInfo) {
	responses := make(chan eazye.Response)

	// all mail from the last week
	since := time.Now().AddDate(0, 0, -7)
	go eazye.GenerateSince(info, since, false, false, responses)

	for resp := range responses {
		if resp.Err != nil {
			log.Print("gen since err: ", resp.Err.Error())
		}
		log.Print(resp.Email)
	}

	return
}

func getAll(info eazye.MailboxInfo) {
	responses := make(chan eazye.Response)
	// get all the mails!!
	go eazye.GenerateAll(info, false, false, responses)

	for resp := range responses {
		if resp.Err != nil {
			log.Print("get all err: ", resp.Err.Error())
		}
		log.Print(resp.Email)
	}

	return
}
