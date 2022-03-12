package cluster

type Node struct {
	ServerId       int      `json:"serverId"`
	VisibleNodeIds []int    `json:"visibleNodeIds"`
	UnknownNodeIps []string `json:"unknownNodeIps"`
}
