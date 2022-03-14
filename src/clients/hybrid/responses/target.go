package responses

import "github.com/aljrubior/amc-ui-rest-facade/clients/hybrid/responses/common"

type Target struct {
	Id                        int                       `json:"id"`
	TimeCreated               int64                     `json:"timeCreated"`
	TimeUpdated               int64                     `json:"timeUpdated"`
	Name                      string                    `json:"name"`
	Type                      string                    `json:"type"`
	ServerType                string                    `json:"serverType"`
	MuleVersion               string                    `json:"muleVersion"`
	GatewayVersion            string                    `json:"gatewayVersion"`
	AgentVersion              string                    `json:"agentVersion"`
	LicenseExpirationDate     int64                     `json:"licenseExpirationDate"`
	CertificateExpirationDate int64                     `json:"certificateExpirationDate"`
	Status                    string                    `json:"status"`
	Addresses                 []common.Address          `json:"addresses"`
	RuntimeInformation        common.RuntimeInformation `json:"runtimeInformation"`
	Servers                   []common.Server           `json:"servers"`
}
