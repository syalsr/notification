package emailer

type Emailer struct {

}

func NewEmailer() Interface {
	return &Emailer{}
}