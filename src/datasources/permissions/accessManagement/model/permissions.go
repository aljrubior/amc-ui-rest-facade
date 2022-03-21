package model

type Permissions struct {
	Applications Permission `json:"applications"`
	Servers      Permission `json:"servers"`
	ServerGroups Permission `json:"serverGroups"`
	Clusters     Permission `json:"clusters"`
	Alerts       Permission `json:"alerts"`
	Services     Permission `json:"services"`
}
