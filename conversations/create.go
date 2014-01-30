package conversations

import (
	"github.com/codegangsta/martini-contrib/render"
	"labix.org/v2/mgo"
)

// Handler to Create a new conversation in the database and send
// out the initial message
func Create(db *mgo.Database, form Form, r render.Render) {
	// create a new conversation and save it in the db
	conv, err := Conversation{
		Email1: form.FromEmail,
		Name1: form.From,
		Email2: form.ToEmail,
		Name2: form.To,
	}.Save(db)

	if err != nil {
		r.JSON(500, err)
		return
	}

	// send the initial message
  err = conv.Mail(form.From, form.Subject, form.Text)
	if err != nil {
		r.JSON(500, err)
		return
	}

	r.JSON(200, conv)
}
