package usecase

import (
	"github.com/rs/zerolog/log"
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
	err := n.Emailer.SendPersonalizedEmail(emails)
	if err != nil {
		log.Err(err).Msgf("cant send personlized email: %w", err)
		return err
	}
	return nil
}

func (n *Notificator) SendCommonEmail(emails *model.CommonEmail) error {
	err := n.Emailer.SendCommonEmail(emails)
	if err != nil {
		log.Err(err).Msgf("cant send common email: %w", err)
		return err
	}
	return nil
}
