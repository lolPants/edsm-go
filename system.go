package edsm

import (
	"fmt"
)

// SystemBodies fetches information about celestial bodies in a system
func (c *Client) SystemBodies(systemName string) (*SystemBodiesStruct, error) {
	req, err := c.newGetRequest("/api-system-v1/bodies")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("systemName", systemName)
	req.URL.RawQuery = q.Encode()

	body := &SystemBodiesStruct{}
	_, err = c.do(req, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Bodies fetches information about celestial bodies for this system
func (s *SystemStruct) Bodies() (*SystemBodiesStruct, error) {
	return s.client.SystemBodies(s.Name)
}

// SystemEstimatedValue fetches estimated scan values of a system
func (c *Client) SystemEstimatedValue(systemName string) (*SystemEstimatedValueStruct, error) {
	req, err := c.newGetRequest("/api-system-v1/estimated-value")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("systemName", systemName)
	req.URL.RawQuery = q.Encode()

	body := &SystemEstimatedValueStruct{}
	_, err = c.do(req, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// EstimatedValue fetches estimated scan values for this system
func (s *SystemStruct) EstimatedValue() (*SystemEstimatedValueStruct, error) {
	return s.client.SystemEstimatedValue(s.Name)
}

// SystemStations fetches information about stations in a system
func (c *Client) SystemStations(systemName string) (*SystemStationsStruct, error) {
	req, err := c.newGetRequest("/api-system-v1/stations")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("systemName", systemName)
	req.URL.RawQuery = q.Encode()

	body := &SystemStationsStruct{}
	_, err = c.do(req, body)
	if err != nil {
		return nil, err
	}

	for _, station := range body.Stations {
		station.client = c
	}

	return body, nil
}

// Stations fetches information about stations for this system
func (s *SystemStruct) Stations() (*SystemStationsStruct, error) {
	return s.client.SystemStations(s.Name)
}

// SystemMarket fetches information about the commodities commomarket in a station
func (c *Client) SystemMarket(systemName string, stationName string) (*SystemMarketStruct, error) {
	req, err := c.newGetRequest("/api-system-v1/stations/market")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("systemName", systemName)
	q.Set("stationName", stationName)
	req.URL.RawQuery = q.Encode()

	body := &SystemMarketStruct{}
	_, err = c.do(req, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Market fetches information about the commodities commomarket in a station
func (s *SystemStruct) Market(stationName string) (*SystemMarketStruct, error) {
	return s.client.SystemMarket(s.Name, stationName)
}

// SystemMarketFromID fetches information about the commodities commomarket in a station using a Market ID
func (c *Client) SystemMarketFromID(marketID int) (*SystemMarketStruct, error) {
	req, err := c.newGetRequest("/api-system-v1/stations/market")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("marketId", fmt.Sprintf("%d", marketID))
	req.URL.RawQuery = q.Encode()

	body := &SystemMarketStruct{}
	_, err = c.do(req, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Market fetches information about the commodities commomarket for this station
func (s *SystemStation) Market() (*SystemMarketStruct, error) {
	return s.client.SystemMarketFromID(s.MarketID)
}

// SystemShipyard fetches information about the shipyard in a station
func (c *Client) SystemShipyard(systemName string, stationName string) (*SystemShipyardStruct, error) {
	req, err := c.newGetRequest("/api-system-v1/stations/shipyard")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("systemName", systemName)
	q.Set("stationName", stationName)
	req.URL.RawQuery = q.Encode()

	body := &SystemShipyardStruct{}
	_, err = c.do(req, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Shipyard fetches information about the shipyard in a station
func (s *SystemStruct) Shipyard(stationName string) (*SystemShipyardStruct, error) {
	return s.client.SystemShipyard(s.Name, stationName)
}

// SystemShipyardFromID fetches information about the shipyard in a station using a Market ID
func (c *Client) SystemShipyardFromID(marketID int) (*SystemShipyardStruct, error) {
	req, err := c.newGetRequest("/api-system-v1/stations/shipyard")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("marketId", fmt.Sprintf("%d", marketID))
	req.URL.RawQuery = q.Encode()

	body := &SystemShipyardStruct{}
	_, err = c.do(req, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Shipyard fetches information about the shipyard for this station
func (s *SystemStation) Shipyard() (*SystemShipyardStruct, error) {
	return s.client.SystemShipyardFromID(s.MarketID)
}

// SystemOutfitting fetches information about outfitting in a station
func (c *Client) SystemOutfitting(systemName string, stationName string) (*SystemOutfittingStruct, error) {
	req, err := c.newGetRequest("/api-system-v1/stations/outfitting")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("systemName", systemName)
	q.Set("stationName", stationName)
	req.URL.RawQuery = q.Encode()

	body := &SystemOutfittingStruct{}
	_, err = c.do(req, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Outfitting fetches information about outfitting in a station
func (s *SystemStruct) Outfitting(stationName string) (*SystemOutfittingStruct, error) {
	return s.client.SystemOutfitting(s.Name, stationName)
}

// SystemOutfittingFromID fetches information about outfitting in a station using a Market ID
func (c *Client) SystemOutfittingFromID(marketID int) (*SystemOutfittingStruct, error) {
	req, err := c.newGetRequest("/api-system-v1/stations/outfitting")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("marketId", fmt.Sprintf("%d", marketID))
	req.URL.RawQuery = q.Encode()

	body := &SystemOutfittingStruct{}
	_, err = c.do(req, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Outfitting fetches information about outfitting for this station
func (s *SystemStation) Outfitting() (*SystemOutfittingStruct, error) {
	return s.client.SystemOutfittingFromID(s.MarketID)
}

// SystemFactions fetches information about factions in a system
func (c *Client) SystemFactions(systemName string, showHistory bool) (*SystemFactionsStruct, error) {
	req, err := c.newGetRequest("/api-system-v1/factions")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("systemName", systemName)
	q.Set("showHistory", boolToQuery(showHistory))
	req.URL.RawQuery = q.Encode()

	body := &SystemFactionsStruct{}
	_, err = c.do(req, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Factions fetches information about factions in this system
func (s *SystemStruct) Factions(showHistory bool) (*SystemFactionsStruct, error) {
	return s.client.SystemFactions(s.Name, showHistory)
}

// SystemTraffic fetches information about traffic in a system
func (c *Client) SystemTraffic(systemName string) (*SystemTrafficStruct, error) {
	req, err := c.newGetRequest("/api-system-v1/traffic")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("systemName", systemName)
	req.URL.RawQuery = q.Encode()

	body := &SystemTrafficStruct{}
	_, err = c.do(req, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Traffic fetches information about traffic in this system
func (s *SystemStruct) Traffic() (*SystemTrafficStruct, error) {
	return s.client.SystemTraffic(s.Name)
}

// SystemDeaths fetches information about deaths in a system
func (c *Client) SystemDeaths(systemName string) (*SystemDeathsStruct, error) {
	req, err := c.newGetRequest("/api-system-v1/deaths")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("systemName", systemName)
	req.URL.RawQuery = q.Encode()

	body := &SystemDeathsStruct{}
	_, err = c.do(req, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Deaths fetches information about deaths in this system
func (s *SystemStruct) Deaths() (*SystemDeathsStruct, error) {
	return s.client.SystemDeaths(s.Name)
}
