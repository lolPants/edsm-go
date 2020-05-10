package edsm

// SystemBodiesStruct represents the result of SystemBodies()
type SystemBodiesStruct struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`

	Bodies []SystemBody `json:"bodies"`
}

// SystemBody represents a celestial body for a given system
type SystemBody struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	SubType string `json:"subType"`

	DistanceToArrival float64 `json:"distanceToArrival"`
	IsMainStar        bool    `json:"isMainStar"`
	IsScoopable       bool    `json:"isScoopable"`

	Age               float64 `json:"age"`
	Luminosity        string  `json:"luminosity"`
	AbsoluteMagnitude float64 `json:"absoluteMagnitude"`
	SolarMasses       float64 `json:"solarMasses"`
	SolarRadius       float64 `json:"solarRadius"`

	SurfaceTemperature  float64 `json:"surfaceTemperature"`
	OrbitalPeriod       float64 `json:"orbitalPeriod"`
	SemiMajorAxis       float64 `json:"semiMajorAxis"`
	OrbitalEccentricity float64 `json:"orbitalEccentricity"`
	OrbitalInclination  float64 `json:"orbitalInclination"`
	ArgOfPeriapsis      float64 `json:"argOfPeriapsis"`

	RotationalPeriod float64 `json:"rotationalPeriod"`
	IsTidallyLocked  bool    `json:"rotationalPeriodTidallyLocked"`
	AxialTilt        float64 `json:"axialTilt"`

	Rings []SystemBodyRing `json:"rings"`
}

// SystemBodyRing represents an asteroid belt around a celestial body
type SystemBodyRing struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Mass int64  `json:"mass"`

	InnerRadius int64 `json:"innerRadius"`
	OuterRadius int64 `json:"outerRadius"`
}

// SystemEstimatedValueStruct represents the result of SystemEstimatedValue()
type SystemEstimatedValueStruct struct {
	ID   int    `json:"id"`
	ID64 int64  `json:"id64"`
	Name string `json:"name"`
	URL  string `json:"url"`

	EstimatedValue       int `json:"estimatedValue"`
	EstimatedValueMapped int `json:"estimatedValueMapped"`

	ValuableBodies []SystemValuableBody `json:"valuableBodies"`
}

// SystemValuableBody represents a valuable celestial body
type SystemValuableBody struct {
	BodyID   int    `json:"bodyId"`
	BodyName string `json:"bodyName"`
	Distance int    `json:"distance"`
	ValueMax int    `json:"valueMax"`
}

// SystemStationsStruct represents the result of SystemStations()
type SystemStationsStruct struct {
	ID   int    `json:"id"`
	ID64 int64  `json:"id64"`
	Name string `json:"name"`
	URL  string `json:"url"`

	Stations []*SystemStation `json:"stations"`
}

// SystemStationBody represents a planetary body on which a station resides
type SystemStationBody struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// SystemStationFaction represents the controlling faction of a station
type SystemStationFaction struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SystemStationUpdateTime contains information about last update times for a station's services
type SystemStationUpdateTime struct {
	Information Time `json:"information"`
	Market      Time `json:"market"`
	Shipyard    Time `json:"shipyard"`
	Outfitting  Time `json:"outfitting"`
}

// SystemStation represents a station within a system
type SystemStation struct {
	ID       int    `json:"id"`
	MarketID int    `json:"marketId"`
	Type     string `json:"type"`
	Name     string `json:"name"`

	Body              *SystemStationBody `json:"body,omitempty"`
	DistanceToArrival float64            `json:"distanceToArrival"`

	Allegiance string `json:"allegiance"`
	Government string `json:"government"`
	Economy    string `json:"economy"`

	HaveMarket     bool     `json:"haveMarket"`
	HaveShipyard   bool     `json:"haveShipyard"`
	HaveOutfitting bool     `json:"haveOutfitting"`
	OtherServices  []string `json:"otherServices"`

	ControllingFaction SystemStationFaction    `json:"controllingFaction"`
	UpdateTime         SystemStationUpdateTime `json:"updateTime"`

	client *Client
}

// SystemMarketStruct represents the result of SystemMarket()
type SystemMarketStruct struct {
	ID         int    `json:"id"`
	ID64       int64  `json:"id64"`
	SystemName string `json:"name"`
	MarketID   int    `json:"marketId"`

	StationID   int    `json:"sId"`
	StationName string `json:"sName"`
	URL         string `json:"url"`

	Commodities []MarketCommodity `json:"commodities"`
}

// MarketCommodity represents a commodity on the commodities market
type MarketCommodity struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	BuyPrice int `json:"buyPrice"`
	Stock    int `json:"stock"`

	SellPrice int `json:"sellPrice"`
	Demand    int `json:"demand"`

	StockBracket int `json:"stockBracket"`
}

// SystemShipyardStruct represents the result of SystemShipyard()
type SystemShipyardStruct struct {
	ID         int    `json:"id"`
	ID64       int64  `json:"id64"`
	SystemName string `json:"name"`
	MarketID   int    `json:"marketId"`

	StationID   int    `json:"sId"`
	StationName string `json:"sName"`
	URL         string `json:"url"`

	Ships []ShipyardShip `json:"ships"`
}

// ShipyardShip represents a ship available to purchase at a station's shipyard
type ShipyardShip struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SystemOutfittingStruct represents the result of SystemOutfitting()
type SystemOutfittingStruct struct {
	ID         int    `json:"id"`
	ID64       int64  `json:"id64"`
	SystemName string `json:"name"`
	MarketID   int    `json:"marketId"`

	StationID   int    `json:"sId"`
	StationName string `json:"sName"`
	URL         string `json:"url"`

	Outfitting []OutfittingModule `json:"outfitting"`
}

// OutfittingModule represents a module available to buy at a station's outfitting service
type OutfittingModule struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// SystemFactionsStruct represents the result of SystemFactions()
type SystemFactionsStruct struct {
	ID   int    `json:"id"`
	ID64 int64  `json:"id64"`
	Name string `json:"name"`
	URL  string `json:"url"`

	ControllingFaction ControllingFaction `json:"controllingFaction"`
	Factions           []*Faction         `json:"factions"`
}

// ControllingFaction represents a system's controlling faction
type ControllingFaction struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Allegiance string `json:"allegiance"`
	Government string `json:"government"`
}

// FactionActiveState represents an active faction state
type FactionActiveState struct {
	State string `json:"state"`
}

// FactionVolatileState represents a recovering/pending faction state
type FactionVolatileState struct {
	State string `json:"state"`
	Trend int    `json:"trend"`
}

// Faction represents a minor faction in the Elite Dangerous galaxy
type Faction struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Allegiance string  `json:"allegiance"`
	Government string  `json:"government"`
	Influence  float64 `json:"influence"`

	State            string                 `json:"state"`
	ActiveStates     []FactionActiveState   `json:"activeStates"`
	RecoveringStates []FactionVolatileState `json:"recoveringStates"`
	PendingStates    []FactionVolatileState `json:"pendingStates"`
	Happiness        string                 `json:"happiness"`

	IsPlayer   bool `json:"isPlayer"`
	LastUpdate int  `json:"lastUpdate"`
}

// SystemTrafficStruct represents the result of SystemTraffic()
type SystemTrafficStruct struct {
	ID   int    `json:"id"`
	ID64 int64  `json:"id64"`
	Name string `json:"name"`
	URL  string `json:"url"`

	Discovery SystemDiscovery        `json:"discovery"`
	Traffic   SystemTraffic          `json:"traffic"`
	Breakdown SystemTrafficBreakdown `json:"breakdown"`
}

// SystemDiscovery contains details of a system's first discovery
type SystemDiscovery struct {
	Commander string `json:"commander"`
	Date      Time   `json:"date"`
}

// SystemTraffic contains stats about a system's traffic over different time periods
type SystemTraffic struct {
	Total int `json:"total"`
	Week  int `json:"week"`
	Day   int `json:"day"`
}

// SystemTrafficBreakdown is a count of different ships passing through a given system
type SystemTrafficBreakdown struct {
	Adder               int `json:"Adder"`
	AllianceChallenger  int `json:"Alliance Challenger"`
	AllianceChieftain   int `json:"Alliance Chieftain"`
	AllianceCrusader    int `json:"Alliance Crusader"`
	Anaconda            int `json:"Anaconda"`
	AspExplorer         int `json:"Asp Explorer"`
	AspScout            int `json:"Asp Scout"`
	BelugaLiner         int `json:"Beluga Liner"`
	CobraMkIII          int `json:"Cobra MkIII"`
	CobraMkIV           int `json:"Cobra MkIV"`
	DiamondbackExplorer int `json:"Diamondback Explorer"`
	DiamondbackScout    int `json:"Diamondback Scout"`
	Dolphin             int `json:"Dolphin"`
	Eagle               int `json:"Eagle"`
	FederalAssaultShip  int `json:"Federal Assault Ship"`
	FederalCorvette     int `json:"Federal Corvette"`
	FederalDropship     int `json:"Federal Dropship"`
	FederalGunship      int `json:"Federal Gunship"`
	FerDeLance          int `json:"Fer-de-Lance"`
	Hauler              int `json:"Hauler"`
	ImperialClipper     int `json:"Imperial Clipper"`
	ImperialCourier     int `json:"Imperial Courier"`
	ImperialCutter      int `json:"Imperial Cutter"`
	ImperialEagle       int `json:"Imperial Eagle"`
	Keelback            int `json:"Keelback"`
	KraitMkII           int `json:"Krait MkII"`
	KraitPhantom        int `json:"Krait Phantom"`
	Mamba               int `json:"Mamba"`
	Orca                int `json:"Orca"`
	Python              int `json:"Python"`
	Sidewinder          int `json:"Sidewinder"`
	Type10Defender      int `json:"Type-10 Defender"`
	Type6Transporter    int `json:"Type-6 Transporter"`
	Type7Transporter    int `json:"Type-7 Transporter"`
	Type9Heavy          int `json:"Type-9 Heavy"`
	ViperMkIII          int `json:"Viper MkIII"`
	ViperMkIV           int `json:"Viper MkIV"`
	Vulture             int `json:"Vulture"`
}

// SystemDeathsStruct represents the result of SystemDeaths()
type SystemDeathsStruct struct {
	ID   int    `json:"id"`
	ID64 int64  `json:"id64"`
	Name string `json:"name"`
	URL  string `json:"url"`

	Deaths struct {
		Total int `json:"total"`
		Week  int `json:"week"`
		Day   int `json:"day"`
	} `json:"deaths"`
}
