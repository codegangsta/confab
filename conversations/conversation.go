package conversations

import (
	"labix.org/v2/mgo"
)

// Conversation stored in the database
type Conversation struct {
	Token  string `json:"token"`
	Email1 string `json:"email1"`
	Email2 string `json:"email2"`
	Name1  string `json:"name1"`
	Name2  string `json:"name2"`
}

func (c Conversation) Save(db *mgo.Database) (Conversation, error) {
  // insert into db
  c.Token = NewUUID().String()
	err := db.C("conversations").Insert(c)
  if err != nil {
    return c, err
  }

  return c, nil
}

func (c Conversation) Mail(from string, subject string, message string) error {
  return nil
}
