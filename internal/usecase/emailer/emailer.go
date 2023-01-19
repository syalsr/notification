package emailer

import (
	"context"
	"strconv"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/rs/zerolog/log"
	"github.com/syalsr/notification/internal/config"
	"github.com/syalsr/notification/internal/model"
)

type Emailer struct {
	Mg     mailgun.Mailgun
	Sender string
}

func NewEmailer(cfg *config.App) Interface {
	return &Emailer{
		Mg:     mailgun.NewMailgun(cfg.MailGunDomain, cfg.MailGunPrivateKey),
		Sender: cfg.MailGunName,
	}
}

func (e *Emailer) SendPersonalizedEmail(emails []model.PersonalizedEmail) error {
	for _, item := range emails {
		message := e.Mg.NewMessage(e.Sender, item.Subject, item.Detail.Text, item.Email)
		
		for _, item := range item.Detail.Attachment {
			message.AddBufferAttachment(item.Name, item.Content)
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		resp, id, err := e.Mg.Send(ctx, message)

		if err != nil {
			log.Err(err).Msgf("cant send email")
			continue
		}
		log.Info().Msgf("email was send - ID: %s, Resp: %s", id, resp)
	}
	return nil
}

func (e *Emailer) SendCommonEmail(emails *model.CommonEmail) error {
	message := e.Mg.NewMessage(e.Sender, emails.Subject, emails.Detail.Text, emails.Emails...)

	for _, item := range emails.Detail.Attachment {
		message.AddBufferAttachment(item.Name, item.Content)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := e.Mg.Send(ctx, message)
	if err != nil {
		log.Err(err).Msgf("cant send email")
		return err
	}
	log.Info().Msgf("email was send - ID: %s, Resp: %s", id, resp)
	return nil
}
