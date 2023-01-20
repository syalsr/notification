package model

// PersonalizedEmail - request
type PersonalizedEmail struct {
	Email   string
	Name    string
	Subject string
	Detail  DetailEmail
}

// CommonEmail - request
type CommonEmail struct {
	Subject string
	Emails []string
	Detail DetailEmail
}

// DetailEmail - detail
type DetailEmail struct {
	Text       string
	Attachment []Attachment
}

// Attachment - detail
type Attachment struct {
	Name string
	Content []byte
}