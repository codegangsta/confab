package models

import (
	"github.com/codegangsta/confab/env"
	"labix.org/v2/mgo"
)

var DB *mgo.Database
var Conversations *mgo.Collection

type Conversation struct {
	Token       string
	Email1      string
	Email2      string
	EmailToken1 string
	EmailToken2 string
}

func init() {
	session, err := mgo.Dial(env.DatabaseURL)
	if err != nil {
		panic(err)
	}

	DB = session.DB(env.Database)
	Conversations = DB.C("conversations")
}

func Foo() {
}
