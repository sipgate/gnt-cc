package repository

type rapiJobResponse struct {
	Status     string        `json:"status"`
	ID         int           `json:"id"`
	StartTs    []int         `json:"start_ts"`
	EndTs      []int         `json:"end_ts"`
	ReceivedTs []int         `json:"received_ts"`
	Summary    []string      `json:"summary"`
	OpStatus   []string      `json:"opstatus"`
	Ops        []interface{} `json:"ops"`
}
