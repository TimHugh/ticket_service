package common

import ()

type Document interface{}

type Adapter interface {
	Find(collection string, id string, result *interface{}) error
	Create(collection string, doc interface{}) error
	Close()
}