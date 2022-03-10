package responses

type WorkerType struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
	Cpu    string  `json:"cpu"`
	Memory string  `json:"memory"`
}
