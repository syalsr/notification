package usecase

import (
	"context"

	"github.com/syalsr/notification/internal/model"
)

// Interface notificator
type Interface interface {
	SendPersonalizedEmail(emails []model.PersonalizedEmail) error
	SendCommonEmail(emails *model.CommonEmail) error
	Run(ctx context.Context, commonEmail <-chan string, personEmail <-chan string)
}
