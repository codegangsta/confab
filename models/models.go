package models

import (
	"github.com/codegangsta/confab/env"
	"github.com/codegangsta/confab/util"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var DB *mgo.Database

func init() {
	session, err := mgo.Dial(env.Get("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	name := env.Get("DATABASE_NAME")
	println("connecting to db:", name)
	DB = session.DB(name)
	Conversations = DB.C("conversations")
}

var Conversations *mgo.Collection

type Conversation struct {
	Token  string `json:"token"`
	Email1 string `json:"email1"`
	Email2 string `json:"email2"`
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
