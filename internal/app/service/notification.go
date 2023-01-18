package service

import "github.com/syalsr/notification/internal/config"

type Notificator struct {

}

func NewNotificator(cfg *config.App) *Notificator {
	return &Notificator{}
}