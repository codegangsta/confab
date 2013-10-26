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

func Test_SavingConversation(t *testing.T) {
	c := models.Conversation{
		Token:  "foobar",
		Email1: "j1@gmail.com",
		Email2: "j2@gmail.com",
	}

	err := models.Conversations.Insert(&c)
	expect(t, err, nil)

	r := models.Conversation{}
	err = models.Conversations.Find(bson.M{"token": "foobar"}).One(&r)
	expect(t, err, nil)
	expect(t, r.Token, "foobar")
	expect(t, r.Email1, "j1@gmail.com")
	expect(t, r.Email2, "j2@gmail.com")

}

func Test_CreateAndGetConversation(t *testing.T) {
	c, err := models.CreateConversation("j1@gmail.com", "j2@gmail.com")
	expect(t, err, nil)
	refute(t, len(c.Token), 0)
	refute(t, len(c.Email1), 0)
	refute(t, len(c.Email2), 0)

	c2, err := models.GetConversation(c.Token)
	expect(t, err, nil)
	expect(t, c2.Token, c.Token)
}
