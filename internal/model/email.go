package model

type PersonalizedEmail struct {
	Email   string
	Name    string
	Subject string
	Detail  DetailEmail
}

type CommonEmail struct {
	Subject string
	Emails []string
	Detail DetailEmail
}

type DetailEmail struct {
	Text       string
	Attachment []Attachment
}

type Attachment struct {
	Name string
	Content []byte
}