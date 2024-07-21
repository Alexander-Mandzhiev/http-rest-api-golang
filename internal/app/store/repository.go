package store

import "github.com/Alexander-Mandzhiev/http-rest-api-golang/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
