package mock

import (
	root "github.com/timhugh/ticket_service"
)

var _ root.LocationRepository = &LocationRepository{}

type LocationRepository struct {
	FindFn      func(string) (*root.Location, error)
	FindInvoked bool

	CreateFn      func(root.Location) error
	CreateInvoked bool
}

func (r *LocationRepository) Find(id string) (*root.Location, error) {
	r.FindInvoked = true
	return r.FindFn(id)
}

func (r *LocationRepository) Create(location root.Location) error {
	r.CreateInvoked = true
	return r.CreateFn(location)
}
