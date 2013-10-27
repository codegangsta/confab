package main

import (
	"labix.org/v2/mgo/bson"
	"testing"
)

func Test_SavingConversation(t *testing.T) {
	c := Conversation{
		Token:  "foobar",
		Email1: "j1@gmail.com",
		Email2: "j2@gmail.com",
	}

	err := Conversations.Insert(&c)
	expect(t, err, nil)

	r := Conversation{}
	err = Conversations.Find(bson.M{"token": "foobar"}).One(&r)
	expect(t, err, nil)
	expect(t, r.Token, "foobar")
	expect(t, r.Email1, "j1@gmail.com")
	expect(t, r.Email2, "j2@gmail.com")

}

func Test_CreateAndGetConversation(t *testing.T) {
	c, err := CreateConversation("j1@gmail.com", "j2@gmail.com")
	expect(t, err, nil)
	refute(t, len(c.Token), 0)
	refute(t, len(c.Email1), 0)
	refute(t, len(c.Email2), 0)

	c2, err := GetConversation(c.Token)
	expect(t, err, nil)
	expect(t, c2.Token, c.Token)
}
