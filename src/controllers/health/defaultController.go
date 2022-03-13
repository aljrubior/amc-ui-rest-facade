package health

type Info struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Version string `json:"version"`
}

type Ping struct {
	Status string `json:"status"`
}

func NewDefaultController() DefaultController {
	return DefaultController{}
}

type DefaultController struct {
}
