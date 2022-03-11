package responses

type Artifact struct {
	Id           int    `json:"id"`
	StorageId    int    `json:"storageId"`
	Name         string `json:"name"`
	FileName     string `json:"fileName"`
	FileChecksum string `json:"fileChecksum"`
	FileSize     int    `json:"fileSize"`
	TimeUpdated  int64  `json:"timeUpdated"`
}
