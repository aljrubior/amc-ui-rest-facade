package model

type Artifact struct {
	Name           string `json:"name"`
	FileName       string `json:"fileName"`
	FileChecksum   string `json:"fileChecksum"`
	LastUpdateTime int64  `json:"lastUpdateTime"`
}
