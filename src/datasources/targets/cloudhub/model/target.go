package model

import (
	"sync"
)

var lock = &sync.Mutex{}

var (
	instance *Target
)

const (
	CloudhubType = "CLOUDHUB"
	CloudhubName = "CloudHub"
)

func NewTarget() Target {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &Target{
			Type: CloudhubType,
			Name: CloudhubName,
		}
	}

	return *instance
}

type Target struct {
	Type string `json:"type"`
	Name string `json:"name"`
}
