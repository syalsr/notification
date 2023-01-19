package model

type PersonalizedEmail struct {
	Email   string
	Name    string
	Subject string
	Detail  DetailEmail
}

type CommonEmail struct {
	Emails []InfoCommonRequest
	Detail DetailEmail
}

type InfoCommonRequest struct {
	Email   string
	Name    string
	Subject string
}

type DetailEmail struct {
	Text       string
	Attachment [][]byte
}
