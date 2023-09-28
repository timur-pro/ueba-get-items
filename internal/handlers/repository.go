//go:generate mockgen -destination=./mocks.go -source=./repository.go -package=handlers
package handlers

import "github.com/timur-pro/ueba-get-items/internal/models"

// UebaRepository интерфейс репозитория
type UebaRepository interface {
	GetItem(string) (models.Record, error)
}
