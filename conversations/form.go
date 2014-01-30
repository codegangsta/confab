package conversations

// The struct that defines the form that will create
// a conversation
type Form struct {
	To        string `form:"to" binding:"required"`
	ToEmail   string `form:"to_email" binding:"required"`
	From      string `form:"from" binding:"required"`
	FromEmail string `form:"from_email" binding:"required"`
	Subject   string `form:"subject" binding:"required"`
	Text      string `form:"text" binding:"required"`
}
