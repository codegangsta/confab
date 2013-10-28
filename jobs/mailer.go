package jobs

import (
  "errors"
  "fmt"
  "net/http"
  "net/url"
  "os"
)

var mailDomain string
var mailApiKey string

type Mail struct {
  Index string
  Token string
  Name string
  To string
  Subject string
  Body string
}

func SendMail(mail Mail) error {
	endpoint := fmt.Sprintf("https://api:%s@api.mailgun.net/v2/%s/messages", mailApiKey, mailDomain)
  res, err := http.PostForm(endpoint, url.Values{
		"from":    {fmt.Sprintf("%s <%s+%s@%s>", mail.Name, mail.Token, mail.Index, mailDomain)},
		"to":      {mail.To},
		"subject": {mail.Subject},
		"text":    {mail.Body},
	})

  if res.StatusCode != 200 {
    return errors.New(res.Status)
  }

  return err
}

func init() {
  mailDomain = os.Getenv("MAIL_DOMAIN")
  mailApiKey = os.Getenv("MAILGUN_API_KEY")
}
