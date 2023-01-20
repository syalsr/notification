package usecase

import "github.com/syalsr/notification/internal/model"

// Interface notificator
type Interface interface {
	SendPersonalizedEmail(emails []model.PersonalizedEmail) error
	SendCommonEmail(emails *model.CommonEmail) error
}
