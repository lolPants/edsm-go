package edsm

// EliteServerStatus represents Elite Dangerous Server Status information
type EliteServerStatus struct {
	LastUpdate Time   `json:"lastUpdate"`
	Message    string `json:"message"`
	Status     int    `json:"status"`
}
