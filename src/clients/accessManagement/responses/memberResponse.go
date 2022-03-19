package responses

import "time"

type MemberResponse struct {
	Id                      string    `json:"id"`
	FirstName               string    `json:"firstName"`
	LastName                string    `json:"lastName"`
	Email                   string    `json:"email"`
	OrganizationId          string    `json:"organizationId"`
	Enabled                 bool      `json:"enabled"`
	IdProviderId            string    `json:"idprovider_id"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
	LastLogin               time.Time `json:"lastLogin"`
	MfaVerifiersConfigured  string    `json:"mfaVerifiersConfigured"`
	MfaVerificationExcluded bool      `json:"mfaVerificationExcluded"`
	IsFederated             bool      `json:"isFederated"`
	Username                string    `json:"username"`
	Type                    string    `json:"type"`
}
