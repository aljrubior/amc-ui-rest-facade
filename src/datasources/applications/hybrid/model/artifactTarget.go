package model

type ArtifactTarget struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
