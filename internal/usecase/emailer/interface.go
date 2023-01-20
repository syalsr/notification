package emailer

import "github.com/syalsr/notification/internal/model"

// Interface emailer
type Interface interface {
	SendPersonalizedEmail(emails []model.PersonalizedEmail) error
	SendCommonEmail(emails *model.CommonEmail) error
}
