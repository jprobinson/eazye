package eazye

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/mail"
	"strings"
	"time"

	"code.google.com/p/go.net/html"
	"github.com/mxk/go-imap/imap"
)

// MailboxInfo holds onto the credentials and other information
// needed for connecting to an IMAP server.
type MailboxInfo struct {
	Host   string
	TLS    bool
	User   string
	Pwd    string
	Folder string
}

// GetAll will pull all emails from the email folder and return them as a list.
func GetAll(info MailboxInfo, markAsRead, delete bool) ([]Email, error) {
	// call chan, put 'em in a list, return
	var emails []Email
	responses := make(chan Response)

	go GenerateAll(info, markAsRead, delete, responses)

	for resp := range responses {
		if resp.Err != nil {
			return emails, resp.Err
		}
		emails = append(emails, resp.Email)
	}

	return emails, nil
}

// GenerateAll will find all emails in the email folder and pass them along to the responses channel.
func GenerateAll(info MailboxInfo, markAsRead, delete bool, responses chan Response) {
	generateMail(info, "ALL", nil, markAsRead, delete, responses)
}

// GetUnread will find all unread emails in the folder and return them as a list.
func GetUnread(info MailboxInfo, markAsRead, delete bool) ([]Email, error) {
	// call chan, put 'em in a list, return
	var emails []Email
	responses := make(chan Response)

	go GenerateUnread(info, markAsRead, delete, responses)

	for resp := range responses {
		if resp.Err != nil {
			return emails, resp.Err
		}
		emails = append(emails, resp.Email)
	}

	return emails, nil
}

// GenerateUnread will find all unread emails in the folder and pass them along to the responses channel.
func GenerateUnread(info MailboxInfo, markAsRead, delete bool, responses chan Response) {
	generateMail(info, "UNSEEN", nil, markAsRead, delete, responses)
}

// GetSince will pull all emails that have an internal date after the given time.
func GetSince(info MailboxInfo, since time.Time, markAsRead, delete bool) ([]Email, error) {
	var emails []Email
	responses := make(chan Response)

	go GenerateSince(info, since, markAsRead, delete, responses)

	for resp := range responses {
		if resp.Err != nil {
			return emails, resp.Err
		}
		emails = append(emails, resp.Email)
	}

	return emails, nil
}

// GenerateSince will find all emails that have an internal date after the given time and pass them along to the
// responses channel.
func GenerateSince(info MailboxInfo, since time.Time, markAsRead, delete bool, responses chan Response) {
	generateMail(info, "", &since, markAsRead, delete, responses)
}

// Email is a simplified email struct containing the basic pieces of an email. If you want more info,
// it should all be available within the Message attribute.
type Email struct {
	Message *mail.Message

	From         string    `json:"from"`
	To           []string  `json:"to"`
	InternalDate time.Time `json:"internal_date"`
	Precedence   string    `json:"precedence"`
	Subject      string    `json:"subject"`
	Body         []byte    `json:"body"`
}

var (
	styleTag       = []byte("style")
	scriptTag      = []byte("script")
	headTag        = []byte("head")
	metaTag        = []byte("meta")
	doctypeTag     = []byte("doctype")
	shapeTag       = []byte("v:shape")
	imageDataTag   = []byte("v:imagedata")
	commentTag     = []byte("!")
	nonVisibleTags = [][]byte{
		styleTag,
		scriptTag,
		headTag,
		metaTag,
		doctypeTag,
		shapeTag,
		imageDataTag,
		commentTag,
	}

	htmlCommentPrefix = []byte("<!--")
	htmlIfBlock       = []byte("[if")
	newLine           = []byte("\n")
)

// VisibleText will return any visible text from an HTML
// email body.
func (e *Email) VisibleText() ([][]byte, error) {
	z := html.NewTokenizer(bytes.NewReader(e.Body))

	var text [][]byte
	skip := false
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			err := z.Err()
			if err == io.EOF {
				return text, nil
			}
			return text, err
		case html.TextToken:
			if !skip {
				tmp := bytes.TrimSpace(z.Text())
				tagText := make([]byte, len(tmp))
				copy(tagText, tmp)

				if len(tagText) == 0 {
					continue
				}
				if bytes.HasPrefix(tagText, htmlCommentPrefix) ||
					bytes.HasPrefix(tagText, htmlIfBlock) {
					continue
				}
				text = append(text, tagText)
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			for _, nvTag := range nonVisibleTags {
				if bytes.Equal(tn, nvTag) {
					skip = (tt == html.StartTagToken)
					break
				}
			}
		}
	}
	return text, nil
}

// String is to spit out a somewhat pretty version of the email.
func (e *Email) String() string {
	return fmt.Sprintf(`
----------------------------
From:           %s
To:             %s
Internal Date:  %s 
Precedence:     %s
Subject:        %s
Body:           %s
----------------------------

`,
		e.From,
		e.To,
		e.InternalDate,
		e.Precedence,
		e.Subject,
		string(e.Body),
	)
}

// Response is a helper struct to wrap the email responses and possible errors.
type Response struct {
	Email Email
	Err   error
}

// newIMAPClient will initiate a new IMAP connection with the given creds.
func newIMAPClient(info MailboxInfo) (*imap.Client, error) {
	var client *imap.Client
	var err error
	if info.TLS {
		client, err = imap.DialTLS(info.Host, new(tls.Config))
		if err != nil {
			return client, err
		}
	} else {
		client, err = imap.Dial(info.Host)
		if err != nil {
			return client, err
		}
	}

	_, err = client.Login(info.User, info.Pwd)
	if err != nil {
		return client, err
	}

	_, err = imap.Wait(client.Select(info.Folder, false))
	if err != nil {
		return client, err
	}

	return client, nil
}

const dateFormat = "02-Jan-2006"

// findEmails will run a find the UIDs of any emails that match the search.:
func findEmails(client *imap.Client, search string, since *time.Time) (*imap.Command, error) {
	var specs []imap.Field
	if len(search) > 0 {
		specs = append(specs, search)
	}

	if since != nil {
		sinceStr := since.Format(dateFormat)
		specs = append(specs, "SINCE", sinceStr)
	}

	// get headers and UID for UnSeen message in src inbox...
	cmd, err := imap.Wait(client.UIDSearch(specs...))
	if err != nil {
		return &imap.Command{}, err
	}
	return cmd, nil
}

func generateMail(info MailboxInfo, search string, since *time.Time, markAsRead, delete bool, responses chan Response) {
	client, err := newIMAPClient(info)
	if err != nil {
		responses <- Response{Err: err}
		return
	}
	defer client.Close(true)

	var cmd *imap.Command
	// find all the UIDs
	cmd, err = findEmails(client, search, since)
	// gotta fetch 'em all
	getEmails(client, cmd, markAsRead, delete, responses)
}

func getEmails(client *imap.Client, cmd *imap.Command, markAsRead, delete bool, responses chan Response) {
	defer close(responses)

	seq := &imap.SeqSet{}
	msgCount := 0
	for _, rsp := range cmd.Data {
		for _, uid := range rsp.SearchResults() {
			msgCount++
			seq.AddNum(uid)
		}
	}

	// nothing to request?! why you even callin me, foolio?
	if seq.Empty() {
		return
	}

	fCmd, err := imap.Wait(client.UIDFetch(seq, "INTERNALDATE", "BODY[]", "UID", "RFC822.HEADER"))
	if err != nil {
		responses <- Response{Err: err}
		return
	}

	var email Email
	for _, msgData := range fCmd.Data {
		msgFields := msgData.MessageInfo().Attrs
		email, err = newEmail(msgFields)
		if err != nil {
			responses <- Response{Err: err}
			return
		}

		responses <- Response{Email: email}

		if !markAsRead {
			err = removeSeen(client, imap.AsNumber(msgFields["UID"]))
			if err != nil {
				responses <- Response{Err: err}
				return
			}
		}

		if delete {
			err = deleteEmail(client, imap.AsNumber(msgFields["UID"]))
			if err != nil {
				responses <- Response{Err: err}
				return
			}
		}
	}
	return
}

func deleteEmail(client *imap.Client, UID uint32) error {
	return alterEmail(client, UID, "\\DELETED", true)
}

func removeSeen(client *imap.Client, UID uint32) error {
	return alterEmail(client, UID, "\\SEEN", false)
}

func alterEmail(client *imap.Client, UID uint32, flag string, plus bool) error {
	flg := "-FLAGS"
	if plus {
		flg = "+FLAGS"
	}
	fSeq := &imap.SeqSet{}
	fSeq.AddNum(UID)
	_, err := imap.Wait(client.UIDStore(fSeq, flg, flag))
	if err != nil {
		return err
	}

	return nil
}

// newEmailMessage will parse an imap.FieldMap into an Email. This
// will expect the message to container the internaldate and the body with
// all headers included.
func newEmail(msgFields imap.FieldMap) (Email, error) {
	var email Email
	// parse the header
	rawHeader := imap.AsBytes(msgFields["RFC822.HEADER"])
	msg, err := mail.ReadMessage(bytes.NewReader(rawHeader))
	if err != nil {
		return email, err
	}

	email = Email{
		Message:      msg,
		InternalDate: imap.AsDateTime(msgFields["INTERNALDATE"]),
		Body:         imap.AsBytes(msgFields["BODY[]"]),
		Precedence:   msg.Header.Get("Precedence"),
		From:         msg.Header.Get("From"),
		To:           strings.Split(msg.Header.Get("To"), ","),
		Subject:      msg.Header.Get("Subject"),
	}

	return email, nil
}
