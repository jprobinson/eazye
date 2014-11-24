package eazye

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"github.com/mxk/go-imap/imap"
	"github.com/sloonz/go-qprintable"
	"golang.org/x/net/html"
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

	From         *mail.Address   `json:"from"`
	To           []*mail.Address `json:"to"`
	InternalDate time.Time       `json:"internal_date"`
	Precedence   string          `json:"precedence"`
	Subject      string          `json:"subject"`
	HTML         []byte          `json:"html"`
	Text         []byte          `json:"text"`
	IsMultiPart  bool            `json:"is_multipart"`
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
	// if theres no HTML, just return text
	if len(e.HTML) == 0 {
		return [][]byte{e.Text}, nil
	}
	z := html.NewTokenizer(bytes.NewReader(e.HTML))

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
				if len(tmp) == 0 {
					continue
				}
				tagText := make([]byte, len(tmp))
				copy(tagText, tmp)

				if bytes.HasPrefix(tagText, htmlCommentPrefix) ||
					bytes.HasPrefix(tagText, htmlIfBlock) ||
					bytes.Equal(tagText, newLine) {
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
HTML:           %s

Text:           %s
----------------------------

`,
		e.From,
		e.To,
		e.InternalDate,
		e.Precedence,
		e.Subject,
		string(e.HTML),
		string(e.Text),
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

func hasEncoding(word string) bool {
	return strings.Contains(word, "=?") && strings.Contains(word, "?=")
}

func isEncodedWord(word string) bool {
	return strings.HasPrefix(word, "=?") && strings.HasSuffix(word, "?=") && strings.Count(word, "?") == 4
}

// decodeRFC2047Word was blatantly ripped off from the net/mail package.
// Hopefully https://code.google.com/p/go/issues/detail?id=4687 will be in 1.4 and we wont need this!
func decodeRFC2047Word(s string) (string, error) {
	fields := strings.Split(s, "?")
	if len(fields) != 5 || fields[0] != "=" || fields[4] != "=" {
		return "", errors.New("address not RFC 2047 encoded")
	}
	charset, enc := strings.ToLower(fields[1]), strings.ToLower(fields[2])
	if charset != "iso-8859-1" && charset != "utf-8" {
		return "", fmt.Errorf("charset not supported: %q", charset)
	}

	in := bytes.NewBufferString(fields[3])
	var r io.Reader
	switch enc {
	case "b":
		r = base64.NewDecoder(base64.StdEncoding, in)
	case "q":
		r = qDecoder{r: in}
	default:
		return "", fmt.Errorf("RFC 2047 encoding not supported: %q", enc)
	}

	dec, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	switch charset {
	case "iso-8859-1":
		b := new(bytes.Buffer)
		for _, c := range dec {
			b.WriteRune(rune(c))
		}
		return b.String(), nil
	case "utf-8":
		return string(dec), nil
	default:
		return s, fmt.Errorf("charset unknown: %s", charset)
	}
}

type qDecoder struct {
	r       io.Reader
	scratch [2]byte
}

func (qd qDecoder) Read(p []byte) (n int, err error) {
	// This method writes at most one byte into p.
	if len(p) == 0 {
		return 0, nil
	}
	if _, err := qd.r.Read(qd.scratch[:1]); err != nil {
		return 0, err
	}
	switch c := qd.scratch[0]; {
	case c == '=':
		if _, err := io.ReadFull(qd.r, qd.scratch[:2]); err != nil {
			return 0, err
		}
		x, err := strconv.ParseInt(string(qd.scratch[:2]), 16, 64)
		if err != nil {
			return 0, fmt.Errorf("mail: invalid RFC 2047 encoding: %q", qd.scratch[:2])
		}
		p[0] = byte(x)
	case c == '_':
		p[0] = ' '
	default:
		p[0] = c
	}
	return 1, nil
}

func parseSubject(subject string) string {
	if !hasEncoding(subject) {
		return subject
	}

	var words []string
	subs := strings.SplitN(subject, " ", -1)
	for _, word := range subs {
		if isEncodedWord(word) {
			// we dont care if it bombed. just take the word
			word, _ = decodeRFC2047Word(word)
		}
		words = append(words, word)
	}

	return strings.Join(words, "")
}

var headerSplitter = []byte("\n")

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

	from, err := mail.ParseAddress(msg.Header.Get("From"))
	if err != nil {
		return email, err
	}

	to, err := mail.ParseAddressList(msg.Header.Get("To"))
	if err != nil {
		return email, err
	}

	email = Email{
		Message:      msg,
		InternalDate: imap.AsDateTime(msgFields["INTERNALDATE"]),
		Precedence:   msg.Header.Get("Precedence"),
		From:         from,
		To:           to,
		Subject:      parseSubject(msg.Header.Get("Subject")),
	}

	// chunk the body up into simple chunks
	rawBody := imap.AsBytes(msgFields["BODY[]"])
	email.HTML, email.Text, email.IsMultiPart, err = parseBody(msg.Header, rawBody)
	return email, err
}

// parseBody will accept a a raw body, break it into all its parts and then convert the
// message to UTF-8 from whatever charset it may have.
func parseBody(header mail.Header, body []byte) (html []byte, text []byte, isMultipart bool, err error) {
	var mediaType string
	var params map[string]string
	mediaType, params, err = mime.ParseMediaType(header.Get("Content-Type"))
	if err != nil {
		return
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		isMultipart = true
		mr := multipart.NewReader(bytes.NewReader(body), params["boundary"])
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}

			slurp, err := ioutil.ReadAll(p)
			if err != nil {
				break
			}

			partMediaType, partParams, err := mime.ParseMediaType(p.Header.Get("Content-Type"))
			if err != nil {
				break
			}

			var htmlT, textT []byte
			htmlT, textT, err = parsePart(partMediaType, partParams["charset"], p.Header.Get("Content-Transfer-Encoding"), slurp)
			if len(htmlT) > 0 {
				html = htmlT
			} else {
				text = textT
			}
		}
	} else {

		splitBody := bytes.SplitN(body, []byte("\r\n\r\n"), 2)
		if len(splitBody) < 2 {
			err = errors.New("unexpected email format. (single part and no \\r\\n\\r\\n separating headers/body")
			return
		}

		body = splitBody[1]
		html, text, err = parsePart(mediaType, params["charset"], header.Get("Content-Transfer-Encoding"), body)
	}
	return
}

func parsePart(mediaType, charsetStr, encoding string, part []byte) (html, text []byte, err error) {
	// deal with charset
	if strings.ToLower(charsetStr) == "iso-8859-1" {
		var cr io.Reader
		cr, err = charset.NewReader("latin1", bytes.NewReader(part))
		if err != nil {
			return
		}

		part, err = ioutil.ReadAll(cr)
		if err != nil {
			return
		}
	}

	// deal with encoding
	var body []byte
	if strings.ToLower(encoding) == "quoted-printable" {
		dec := qprintable.NewDecoder(qprintable.WindowsTextEncoding, bytes.NewReader(part))
		body, err = ioutil.ReadAll(dec)
		if err != nil {
			return
		}
	} else {
		body = part
	}

	// deal with media type
	mediaType = strings.ToLower(mediaType)
	switch {
	case strings.Contains(mediaType, "text/html"):
		html = body
	case strings.Contains(mediaType, "text/plain"):
		text = body
	}
	return
}
