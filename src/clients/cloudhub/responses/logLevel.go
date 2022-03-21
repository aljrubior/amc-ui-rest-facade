package responses

type LogLevel struct {
	PackageName string `json:"packageName"`
	Level       string `json:"level"`
}
