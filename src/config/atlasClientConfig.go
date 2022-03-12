package config

func NewAtlasConfigClient() AtlasConfigClient {
	return AtlasConfigClient{
		Protocol:    "http",
		Host:        "atlas.amc.svc",
		Port:        "8080",
		TargetsPath: "/organizations/%s/providers/%s/targets",
	}
}

type AtlasConfigClient struct {
	Protocol,
	Host,
	Port,
	TargetsPath string
}
