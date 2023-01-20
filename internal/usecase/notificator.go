package usecase

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/syalsr/notification/internal/model"
	"github.com/syalsr/notification/internal/usecase/emailer"
)

// Notificator - obj which implement interface
type Notificator struct {
	Emailer emailer.Interface
}

// NewNotificator - create obj which implement interface
func NewNotificator(e emailer.Interface) Interface {
	return &Notificator{
		Emailer: e,
	}
}

// SendPersonalizedEmail - send personalized email
func (n *Notificator) SendPersonalizedEmail(emails []model.PersonalizedEmail) error {
	err := n.Emailer.SendPersonalizedEmail(emails)
	if err != nil {
		log.Err(err).Msgf("cant send personlized email: %w", err)
		return fmt.Errorf("cant send personlized email: %w", err)
	}
	return nil
}

// SendCommonEmail - send common email
func (n *Notificator) SendCommonEmail(emails *model.CommonEmail) error {
	err := n.Emailer.SendCommonEmail(emails)
	if err != nil {
		log.Err(err).Msgf("cant send common email: %w", err)
		return fmt.Errorf("cant send common email: %w", err)
	}
	return nil
}
