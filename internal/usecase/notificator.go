package usecase

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/syalsr/notification/internal/model"
	"github.com/syalsr/notification/internal/usecase/emailer"
	"github.com/syalsr/notification/internal/utils"
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

func (n *Notificator) Run(ctx context.Context, commonEmail <-chan string, personEmail <-chan string) {
	for {
		select{
		case <- ctx.Done():
			return
		case text := <- commonEmail:
			err := n.SendCommonEmail(utils.CommonEmailParse(text))
			if err != nil {
				log.Err(err).Msgf("cant send common email: %w", err)
			}
		case text := <- personEmail:
			err := n.SendPersonalizedEmail(utils.PersonEmailParse(text))
			if err != nil {
				log.Err(err).Msgf("cant send common email: %w", err)
			}
		}
	}
}