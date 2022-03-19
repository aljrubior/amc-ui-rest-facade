package requests

import "fmt"

func NewAuthorizeRequest(orgId, envId, namespace, action string, resourceNames []string) AuthorizeRequest {

	resources := make([]string, 0)

	for _, v := range resourceNames {
		resources = append(resources, fmt.Sprintf("%s/%s/%s", orgId, envId, v))
		resources = append(resources, fmt.Sprintf("%s/%s/%s/**", orgId, envId, v))
	}

	return AuthorizeRequest{
		Namespace: namespace,
		Action:    action,
		Resources: resources,
	}
}

type AuthorizeRequest struct {
	Namespace string   `json:"namespace"`
	Action    string   `json:"action"`
	Resources []string `json:"resources"`
}
