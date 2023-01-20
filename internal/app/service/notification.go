package service

import (
	"context"

	"github.com/syalsr/notification/internal/config"
	"github.com/syalsr/notification/internal/model"
	"github.com/syalsr/notification/internal/usecase"
	api "github.com/syalsr/notification/pkg/v1"
)

// Notificator - notificator service
type Notificator struct {
	api.UnimplementedNotificationServiceServer
	notif usecase.Interface
}

// NewNotificator - create new obj which implement Interface Notificator
func NewNotificator(cfg *config.App, n usecase.Interface) *Notificator {
	return &Notificator{
		notif: n,
	}
}

// SendPersonalizedEmail - send personalized email
func (n *Notificator) SendPersonalizedEmail(ctx context.Context, req *api.SendPersonalizedEmailRequest) (*api.SendEmailResponse, error) {
	emails := make([]model.PersonalizedEmail, len(req.Emails))
	for _, idx := range req.Emails {
		attachament := make([]model.Attachment, len(idx.Detail.Attachment))
		for _, item := range idx.Detail.Attachment {
			attachament = append(attachament, model.Attachment{Name: item.Name, Content: item.Content})
		}

		emails = append(
			emails,
			model.PersonalizedEmail{
				Email:   idx.Email,
				Subject: idx.Subject,
				Detail: model.DetailEmail{
					Text:       idx.Detail.Text,
					Attachment: attachament,
				},
			})
	}
	err := n.notif.SendPersonalizedEmail(emails)
	if err != nil {
		return &api.SendEmailResponse{Status: err.Error()}, err
	}
	return &api.SendEmailResponse{Status: "OK"}, nil
}

// SendCommonEmail - send common email
func (n *Notificator) SendCommonEmail(ctx context.Context, req *api.SendCommonEmailRequest) (*api.SendEmailResponse, error) {
	attachment := make([]model.Attachment, len(req.Detail.Attachment))
	for _, item := range req.Detail.Attachment {
		attachment = append(attachment, model.Attachment{Name: item.Name, Content: item.Content})
	}
	err := n.notif.SendCommonEmail(&model.CommonEmail{
		Emails: req.Emails,
		Subject: req.Subject,
		Detail: model.DetailEmail{Text: req.Detail.Text, Attachment: attachment},
	})
	if err != nil {
		return &api.SendEmailResponse{Status: err.Error()}, err
	}
	return &api.SendEmailResponse{Status: "OK"}, nil
}
