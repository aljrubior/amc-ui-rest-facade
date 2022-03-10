package responses

type Worker struct {
	Type                WorkerType `json:"type"`
	Amount              int        `json:"amount"`
	RemainingOrgWorkers float64    `json:"remainingOrgWorkers"`
	TotalOrgWorkers     float64    `json:"totalOrgWorkers"`
}
