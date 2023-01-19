package service

import (
	"context"

	"github.com/syalsr/notification/internal/config"
	"github.com/syalsr/notification/internal/model"
	"github.com/syalsr/notification/internal/usecase"
	api "github.com/syalsr/notification/pkg/v1"
)

type Notificator struct {
	api.UnimplementedNotificationServiceServer
	usecase.Notificator
}

func NewNotificator(cfg *config.App, notif usecase.Notificator) *Notificator {
	return &Notificator{
		Notificator: notif,
	}
}

func (n *Notificator) SendPersonalizedEmail(ctx context.Context, req *api.SendPersonalizedEmailRequest) (*api.SendEmailResponse, error) {
	emails := make([]model.PersonalizedEmail, len(req.Emails))
	for _, idx := range req.Emails {
		emails = append(
			emails,
			model.PersonalizedEmail{
				Name:    idx.Name,
				Email:   idx.Email,
				Subject: idx.Subject,
				Detail: model.DetailEmail{
					Text:       idx.Detail.Text,
					Attachment: idx.Detail.Attachment,
				},
			})
	}
	err := n.Notificator.SendPersonalizedEmail(emails)
	if err != nil {
		return &api.SendEmailResponse{Status: err.Error()}, err
	}
	return nil, nil
}

func (n *Notificator) SendCommonEmail(ctx context.Context, req *api.SendCommonEmailRequest) (*api.SendEmailResponse, error) {
	emails := make([]model.InfoCommonRequest, len(req.Emails))
	for _, item := range req.Emails {
		emails = append(emails, model.InfoCommonRequest{Email: item.Email, Name: item.Name, Subject: item.Subject})
	}
	err := n.Notificator.SendCommonEmail(&model.CommonEmail{
		Emails: emails,
		Detail: model.DetailEmail{Text: req.Detail.Text, Attachment: req.Detail.Attachment},
	})
	if err != nil {
		return &api.SendEmailResponse{Status: err.Error()}, err
	}
	return nil, nil
}
