package models

import (
	"github.com/codegangsta/confab/env"
	"github.com/codegangsta/confab/util"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var DB *mgo.Database

func init() {
	session, err := mgo.Dial(env.DatabaseURL)
	if err != nil {
		panic(err)
	}

	DB = session.DB(env.Database)
	Conversations = DB.C("conversations")
}

var Conversations *mgo.Collection

type Conversation struct {
	Token  string
	Email1 string
	Email2 string
}

func CreateConversation(email1 string, email2 string) (*Conversation, error) {
	token := util.NewUUID().String()

	c := Conversation{
		Token:  token,
		Email1: email1,
		Email2: email2,
	}

	err := Conversations.Insert(&c)

	return &c, err
}

func GetConversation(token string) (*Conversation, error) {
	c := Conversation{}
	err := Conversations.Find(bson.M{"token": token}).One(&c)

	return &c, err
}
