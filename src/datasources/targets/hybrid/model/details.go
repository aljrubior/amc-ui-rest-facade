package model

type Details struct {
	RuntimeVersion      string         `json:"runtimeVersion,omitempty"`
	Type                string         `json:"type,omitempty"`
	AgentVersion        string         `json:"agentVersion,omitempty"`
	Addresses           []Address      `json:"addresses,omitempty"`
	Servers             []Target       `json:"servers,omitempty"`
	CurrentClusteringIp string         `json:"currentClusteringIp,omitempty"`
	MulticastEnabled    bool           `json:"multicastEnabled,omitempty"`
	PrimaryNodeId       int            `json:"primaryNodeId,omitempty"`
	VisibilityMap       *VisibilityMap `json:"visibilityMap,omitempty"`
}
