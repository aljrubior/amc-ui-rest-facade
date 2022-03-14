package model

type Target struct {
	Type     string `json:"type"`
	Provider string `json:"provider"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
}
