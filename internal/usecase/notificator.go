package usecase

import (
	"github.com/syalsr/notification/internal/model"
	"github.com/syalsr/notification/internal/usecase/emailer"
)

type Notificator struct {
	Emailer emailer.Interface
}

func NewNotificator(e emailer.Interface) Interface {
	return &Notificator{
		Emailer: e,
	}
}

func (n *Notificator) SendPersonalizedEmail(emails []model.PersonalizedEmail) error {

}

func (n *Notificator) SendCommonEmail(emails *model.CommonEmail) error {

}
