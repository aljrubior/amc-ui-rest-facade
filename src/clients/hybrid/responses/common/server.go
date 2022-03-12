package common

type Server struct {
	Id                        int                `json:"id"`
	TimeCreated               int64              `json:"timeCreated"`
	TimeUpdated               int64              `json:"timeUpdated"`
	Name                      string             `json:"name"`
	Type                      string             `json:"type"`
	ServerType                string             `json:"serverType"`
	MuleVersion               string             `json:"muleVersion"`
	GatewayVersion            string             `json:"gatewayVersion"`
	AgentVersion              string             `json:"agentVersion"`
	LicenseExpirationDate     int64              `json:"licenseExpirationDate"`
	CertificateExpirationDate int64              `json:"certificateExpirationDate"`
	Status                    string             `json:"status"`
	Addresses                 []Address          `json:"addresses"`
	ClusterId                 int                `json:"clusterId,omitempty"`
	ClusterName               string             `json:"clusterName"`
	CurrentClusteringIp       string             `json:"currentClusteringIp"`
	CurrentClusteringPort     int                `json:"currentClusteringPort"`
	RuntimeInformation        RuntimeInformation `json:"runtimeInformation"`
}
