package formatters

type Artifact struct {
	LastUpdateTime int64  `json:"lastUpdateTime"`
	CreateTime     int64  `json:"createTime"`
	Name           string `json:"name"`
}
