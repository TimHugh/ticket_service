package common

import ()

type Location struct {
	ID           string
	SignatureKey string
}

type LocationRepository interface {
	Find(string) *Location
	Store(Location)
}