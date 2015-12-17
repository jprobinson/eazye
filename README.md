eazye [![GoDoc](https://godoc.org/github.com/jprobinson/eazye?status.svg)](https://godoc.org/github.com/jprobinson/eazye) ![Travis CI](https://travis-ci.org/jprobinson/eazye.svg?branch=master)
======
### _The `Gangsta Gangsta` way to pull email._

#### Getting your emails is eazy...

##### Start by putting credentials and mailbox info into a MailboxInfo:
```go
type MailboxInfo struct {
    Host   string
    TLS    bool
    User   string
    Pwd    string
    Folder string
}
```

##### Then connect and pull all mail...
```go
// GetAll will pull all emails from the email folder and return them as a list.
func GetAll(info MailboxInfo, markAsRead, delete bool) ([]Email, error)
```

```go
// GenerateAll will find all emails in the email folder and pass them along to the response channel.
func GenerateAll(info MailboxInfo, markAsRead, delete bool) (chan Response, error)
```

##### ... or all unread mail...
```go
// GetUnread will find all unread emails in the folder and return them as a list.
func GetUnread(info MailboxInfo, markAsRead, delete bool) ([]Email, error)
```

```go
// GenerateUnread will find all unread emails in the folder and pass them along to the response channel.
func GenerateUnread(info MailboxInfo, markAsRead, delete bool) (chan Response, error)
```


#####  ... or all mail received since a particular date.
```go
// GetSince will pull all emails that have an internal date after the given time.
func GetSince(info MailboxInfo, since time.Time, markAsRead, delete bool)
```

```go
// GenerateSince will find all emails that have an internal date after the given time and pass them along to the responses channel.
func GenerateSince(info MailboxInfo, since time.Time, markAsRead, delete bool) (chan Response, error)
```

##### `eazye` will pull out the most common headers and bits but also provides the `mail.Message` in case you want to pull additional data.

```go
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
```

##### The `eazye` Email type also has a handy `func (e *Email) VisibleText() ([][]byte, error)` that will return all the visible text from an HTML email or the body of a Text email if HTML is not available.

##### If you have a lot of messages and do not want to load everything into memory, use the GenerateXXX functions and the emails will be passed along on a channel of `eazye.Response`s. To configure the buffer size of the response channel, you can use the exported `GenerateBufferSize` variable, which is defaulted to 100.
```go
// Response is a helper struct to wrap the email responses and possible errors.
type Response struct {
    Email Email
    Err   error
}
```

This package has several dependencies: 
* github.com/mxk/go-imap/imap
* github.com/paulrosania/charset
* github.com/paulrosania/go-charset/data
* github.com/sloonz/go-qprintable
* golang.org/x/net/html
