package test

import (
	"github.com/codegangsta/confab/env"
	"github.com/codegangsta/confab/models"
	"labix.org/v2/mgo/bson"
	"testing"
)

// drop the database
func init() {
	println("Database:", env.Database)
	err := models.DB.DropDatabase()
	if err != nil {
		panic(err)
	}
}

func TestSavingConversation(t *testing.T) {
	c := models.Conversation{
		Token:       "foobar",
		Email1:      "j1@gmail.com",
		Email2:      "j2@gmail.com",
		EmailToken1: "batbaz1",
		EmailToken2: "batbaz2",
	}

	err := models.Conversations.Insert(&c)
	expect(t, err, nil)

	r := models.Conversation{}
	err = models.Conversations.Find(bson.M{"token": "foobar"}).One(&r)
	expect(t, err, nil)
	expect(t, r.Token, "foobar")
	expect(t, r.Email1, "j1@gmail.com")
	expect(t, r.Email2, "j2@gmail.com")
	expect(t, r.EmailToken1, "batbaz1")
	expect(t, r.EmailToken2, "batbaz2")

}
