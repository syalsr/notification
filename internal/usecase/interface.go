package usecase

import "github.com/syalsr/notification/internal/model"

type Notificator interface {
	SendPersonalizedEmail(emails []model.PersonalizedEmail) error
	SendCommonEmail(emails *model.CommonEmail) error
}
