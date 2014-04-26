eazye
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
// GenerateAll will find all emails in the email folder and pass them along to the responses channel.
func GenerateAll(info MailboxInfo, markAsRead, delete bool, responses chan Response)
```

##### ... or all unread mail...
```go
// GetUnread will find all unread emails in the folder and return them as a list.
func GetUnread(info MailboxInfo, markAsRead, delete bool) ([]Email, error)
```

```go
// GenerateUnread will find all unread emails in the folder and pass them along to the responses channel.
func GenerateUnread(info MailboxInfo, markAsRead, delete bool, responses chan Response)
```


#####  ... or all mail received since a particular date.
```go
// GetSince will pull all emails that have an internal date after the given time.
func GetSince(info MailboxInfo, since time.Time, markAsRead, delete bool)
```

```go
// GenerateSince will find all emails that have an internal date after the given time and pass them along to the responses channel.
func GenerateSince(info MailboxInfo, since time.Time, markAsRead, delete bool, responses chan Response)
```

##### `eazye` will pull out the most common headers and bits but also provides the `mail.Message` in case you want to pull additional data.

```go
type Email struct {
    Message *mail.Message

    From         string    `json:"from"`
    To           []string  `json:"to"`
    InternalDate time.Time `json:"internal_date"`
    Precedence   string    `json:"precednce"`
    Subject      string    `json:"subject"`
    Body         string    `json:"body"`
}
```

##### If you have a lot of messages and do not want to load everything into memory, use the GenerateXXX functions and the emails will be passed along on a channel of `eazye.Response`s
```go
// Response is a helper struct to wrap the email responses and possible errors.
type Response struct {
    Email Email
    Err   error
}
```

This package has one dependency you will need to `go get`: `github.com/mxk/go-imap/imap`.
