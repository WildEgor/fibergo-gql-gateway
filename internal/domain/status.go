package domains

// Status - holds the information for the server's status
type StatusDomain struct {
	Status      string `json:"status"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
}
