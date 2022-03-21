package model

type Permission struct {
	Read   bool `json:"read"`
	Create bool `json:"create"`
	Modify bool `json:"modify"`
	Delete bool `json:"delete"`
}
