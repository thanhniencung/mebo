package repository

import (
	"mibo/model"
)

type MailRepo interface {
	Save(email model.History) error
	List() ([]model.History, error)
}