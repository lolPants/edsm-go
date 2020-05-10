package edsm

// SystemStruct represents an EDSM System
type SystemStruct struct {
	Name string `json:"name"`
	ID   *int   `json:"id"`
	ID64 *int64 `json:"id64"`

	Coordinates       *Vector3 `json:"coords"`
	CoordinatesLocked *bool    `json:"coordsLocked"`
	Distance          *float64 `json:"distance"`

	RequirePermit *bool   `json:"requirePermit"`
	PermitName    *string `json:"permitName"`

	Information *SystemInformation `json:"information"`

	client *Client
}

// SystemInformation describes information about an EDSM system
type SystemInformation struct {
	Allegiance   string `json:"allegiance"`
	Government   string `json:"government"`
	Faction      string `json:"faction"`
	FactionState string `json:"factionState"`
	Population   int    `json:"population"`
	Security     string `json:"security"`
	Economy      string `json:"economy"`
}
