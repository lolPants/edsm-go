package edsm

import (
	"fmt"
	"net/http"
	"time"
)

// SystemOptions defines options for the System method
type SystemOptions struct {
	// Set to `true` to get the internal ID for this system.
	ShowID bool

	// Set to `true` to get the system coordinates.
	// If coordinates are unknown, the coords key will not be returned.
	ShowCoordinates bool

	// Set to `true` to get the system permit if there is one.
	// If the permit is named, also return permitName.
	ShowPermit bool

	// Set to `true` to get the system information like allegiance, government, etc.
	// If no information are stored, an empty array will be returned.
	ShowInformation bool

	// Set to `true` to get the system primary star if known.
	// If no primary star is stored, a `nil` will be returned.
	ShowPrimaryStar bool

	// Set to `true` to get system even if hidden in the database.
	// Hidden system are generally typo errors, renamed system in the game, etc.
	IncludeHidden bool
}

// System fetches information about a system
func (c *Client) System(systemName string, options *SystemOptions) (*SystemStruct, error) {
	req, err := c.newGetRequest("/api-v1/system")
	if err != nil {
		return nil, err
	}

	var o SystemOptions
	if options == nil {
		o = SystemOptions{}
	} else {
		o = *options
	}

	q := req.URL.Query()
	q.Set("systemName", systemName)
	q.Set("showId", boolToQuery(o.ShowID))
	q.Set("showCoordinates", boolToQuery(o.ShowCoordinates))
	q.Set("showPermit", boolToQuery(o.ShowPermit))
	q.Set("showInformation", boolToQuery(o.ShowInformation))
	q.Set("showPrimaryStar", boolToQuery(o.ShowPrimaryStar))
	q.Set("includeHidden", boolToQuery(o.IncludeHidden))

	req.URL.RawQuery = q.Encode()
	sys := &SystemStruct{}

	_, err = c.do(req, sys)
	if err != nil {
		return nil, err
	}

	sys.client = c
	return sys, nil
}

// SystemsOptions defines options for Systems methods
type SystemsOptions struct {
	// Set to `true` to get the internal ID for this system.
	ShowID bool

	// Set to `true` to get the system coordinates.
	// If coordinates are unknown, the coords key will not be returned.
	ShowCoordinates bool

	// Set to `true` to get the system permit if there is one.
	// If the permit is named, also return permitName.
	ShowPermit bool

	// Set to `true` to get the system information like allegiance, government, etc.
	// If no information are stored, an empty array will be returned.
	ShowInformation bool

	// Set to `true` to get the system primary star if known.
	// If no primary star is stored, a `nil` will be returned.
	ShowPrimaryStar bool

	// If you only want to receive systems updated after a specific date & time, use this parameter.
	// That parameter is inclusive. All dates must be UTC.
	//
	// StartDateTime and EndDateTime can be use together or separately, but the maximum interval will be 3 DAY.
	// If both are set, EndDateTime will be used to calculate the interval.
	StartDateTime *time.Time

	// If you only want to receive systems updated before a specific date & time, use this parameter.
	// That parameter is inclusive. All dates must be UTC.
	//
	// StartDateTime and EndDateTime can be use together or separately, but the maximum interval will be 3 DAY.
	// If both are set, EndDateTime will be used to calculate the interval.
	EndDateTime *time.Time

	// Set to `true` to get only systems featured on the EDSM home page.
	// Featured systems are systems here we try to trilatered coordinates with a highest priority.
	OnlyFeatured bool

	// Set to `true` to get only systems with known coordinates.
	OnlyKnownCoordinates bool

	// Set to `true` to get only systems without coordinates.
	OnlyUnknownCoordinates bool

	// Set to `true` to get system even if hidden in the database.
	// Hidden system are generally typo errors, renamed system in the game, etc.
	IncludeHidden bool
}

// Systems fetches information about systems
func (c *Client) Systems(systemNames []string, options *SystemsOptions) ([]*SystemStruct, error) {
	req, err := c.newGetRequest("/api-v1/systems")
	if err != nil {
		return nil, err
	}

	var o SystemsOptions
	if options == nil {
		o = SystemsOptions{}
	} else {
		o = *options
	}

	q := req.URL.Query()
	for _, name := range systemNames {
		q.Add("systemName[]", name)
	}

	req.URL.RawQuery = q.Encode()
	return systemsCommon(c, req, &o)
}

// SystemsSingle fetches information about systems
func (c *Client) SystemsSingle(systemName string, options *SystemsOptions) ([]*SystemStruct, error) {
	req, err := c.newGetRequest("/api-v1/systems")
	if err != nil {
		return nil, err
	}

	var o SystemsOptions
	if options == nil {
		o = SystemsOptions{}
	} else {
		o = *options
	}

	q := req.URL.Query()
	q.Set("systemName", systemName)
	req.URL.RawQuery = q.Encode()

	return systemsCommon(c, req, &o)
}

func systemsCommon(c *Client, req *http.Request, o *SystemsOptions) ([]*SystemStruct, error) {
	q := req.URL.Query()
	q.Set("showId", boolToQuery(o.ShowID))
	q.Set("showCoordinates", boolToQuery(o.ShowCoordinates))
	q.Set("showPermit", boolToQuery(o.ShowPermit))
	q.Set("showInformation", boolToQuery(o.ShowInformation))
	q.Set("showPrimaryStar", boolToQuery(o.ShowPrimaryStar))
	q.Set("onlyFeatured", boolToQuery(o.OnlyFeatured))
	q.Set("onlyKnownCoordinates", boolToQuery(o.OnlyKnownCoordinates))
	q.Set("onlyUnknownCoordinates", boolToQuery(o.OnlyUnknownCoordinates))
	q.Set("includeHidden", boolToQuery(o.IncludeHidden))

	if o.StartDateTime != nil {
		q.Set("startDateTime", formatDateTime(*o.StartDateTime))
	}

	if o.EndDateTime != nil {
		q.Set("endDateTime", formatDateTime(*o.EndDateTime))
	}

	req.URL.RawQuery = q.Encode()
	arr := make([]*SystemStruct, 0)
	sys := &arr

	_, err := c.do(req, sys)
	if err != nil {
		return nil, err
	}

	for _, s := range *sys {
		s.client = c
	}

	return *sys, nil
}

// SphereSystemsOptions defines options for SphereSystems methods
type SphereSystemsOptions struct {
	// Set to the desired radius in Ly.
	// Maximum value is 100.
	Radius float64

	// Set to a value between 0 and radius to reduce the returned results.
	MinRadius float64

	// Set to `true` to get the internal ID for this system.
	ShowID bool

	// Set to `true` to get the system coordinates.
	// If coordinates are unknown, the coords key will not be returned.
	ShowCoordinates bool

	// Set to `true` to get the system permit if there is one.
	// If the permit is named, also return permitName.
	ShowPermit bool

	// Set to `true` to get the system information like allegiance, government, etc.
	// If no information are stored, an empty array will be returned.
	ShowInformation bool

	// Set to `true` to get the system primary star if known.
	// If no primary star is stored, a `nil` will be returned.
	ShowPrimaryStar bool
}

// SphereSystems fetches systems in a sphere radius around `systemName`
func (c *Client) SphereSystems(systemName string, options *SphereSystemsOptions) ([]*SystemStruct, error) {
	req, err := c.newGetRequest("/api-v1/sphere-systems")
	if err != nil {
		return nil, err
	}

	var o SphereSystemsOptions
	if options == nil {
		o = SphereSystemsOptions{}
	} else {
		o = *options
	}

	q := req.URL.Query()
	q.Set("systemName", systemName)
	req.URL.RawQuery = q.Encode()

	return sphereSystemsCommon(c, req, &o)
}

// SphereSystemsCoords fetches systems in a sphere radius around `coords`
func (c *Client) SphereSystemsCoords(coords Vector3, options *SphereSystemsOptions) ([]*SystemStruct, error) {
	req, err := c.newGetRequest("/api-v1/sphere-systems")
	if err != nil {
		return nil, err
	}

	var o SphereSystemsOptions
	if options == nil {
		o = SphereSystemsOptions{}
	} else {
		o = *options
	}

	q := req.URL.Query()
	q.Set("x", fmt.Sprintf("%f", coords.X))
	q.Set("y", fmt.Sprintf("%f", coords.Y))
	q.Set("z", fmt.Sprintf("%f", coords.Z))
	req.URL.RawQuery = q.Encode()

	return sphereSystemsCommon(c, req, &o)
}

func sphereSystemsCommon(c *Client, req *http.Request, o *SphereSystemsOptions) ([]*SystemStruct, error) {
	var radius float64
	if o.Radius == 0 {
		radius = 50
	} else {
		radius = o.Radius
	}

	if radius > 100 {
		radius = 100
	}

	var minRadius float64
	if o.MinRadius > radius {
		minRadius = radius
	} else {
		minRadius = o.MinRadius
	}

	q := req.URL.Query()
	q.Set("radius", fmt.Sprintf("%f", radius))
	q.Set("minRadius", fmt.Sprintf("%f", minRadius))
	q.Set("showId", boolToQuery(o.ShowID))
	q.Set("showCoordinates", boolToQuery(o.ShowCoordinates))
	q.Set("showPermit", boolToQuery(o.ShowPermit))
	q.Set("showInformation", boolToQuery(o.ShowInformation))
	q.Set("showPrimaryStar", boolToQuery(o.ShowPrimaryStar))

	req.URL.RawQuery = q.Encode()
	arr := make([]*SystemStruct, 0)
	sys := &arr

	_, err := c.do(req, sys)
	if err != nil {
		return nil, err
	}

	for _, s := range *sys {
		s.client = c
	}

	return *sys, nil
}

// CubeSystemsOptions defines options for CubeSystems methods
type CubeSystemsOptions struct {
	// Set to the desired size of the cube in Ly.
	// Maximum value is 200.
	Size float64

	// Set to `true` to get the internal ID for this system.
	ShowID bool

	// Set to `true` to get the system coordinates.
	// If coordinates are unknown, the coords key will not be returned.
	ShowCoordinates bool

	// Set to `true` to get the system permit if there is one.
	// If the permit is named, also return permitName.
	ShowPermit bool

	// Set to `true` to get the system information like allegiance, government, etc.
	// If no information are stored, an empty array will be returned.
	ShowInformation bool

	// Set to `true` to get the system primary star if known.
	// If no primary star is stored, a `nil` will be returned.
	ShowPrimaryStar bool
}

// CubeSystems fetches systems in a cube around `systemName`
func (c *Client) CubeSystems(systemName string, options *CubeSystemsOptions) ([]*SystemStruct, error) {
	req, err := c.newGetRequest("/api-v1/cube-systems")
	if err != nil {
		return nil, err
	}

	var o CubeSystemsOptions
	if options == nil {
		o = CubeSystemsOptions{}
	} else {
		o = *options
	}

	q := req.URL.Query()
	q.Set("systemName", systemName)
	req.URL.RawQuery = q.Encode()

	return cubeSystemsCommon(c, req, &o)
}

// CubeSystemsCoords fetches systems in a cube around `coords`
func (c *Client) CubeSystemsCoords(coords Vector3, options *CubeSystemsOptions) ([]*SystemStruct, error) {
	req, err := c.newGetRequest("/api-v1/cube-systems")
	if err != nil {
		return nil, err
	}

	var o CubeSystemsOptions
	if options == nil {
		o = CubeSystemsOptions{}
	} else {
		o = *options
	}

	q := req.URL.Query()
	q.Set("x", fmt.Sprintf("%f", coords.X))
	q.Set("y", fmt.Sprintf("%f", coords.Y))
	q.Set("z", fmt.Sprintf("%f", coords.Z))
	req.URL.RawQuery = q.Encode()

	return cubeSystemsCommon(c, req, &o)
}

func cubeSystemsCommon(c *Client, req *http.Request, o *CubeSystemsOptions) ([]*SystemStruct, error) {
	var size float64
	if o.Size == 0 {
		size = 50
	} else {
		size = o.Size
	}

	if size > 200 {
		size = 200
	}

	q := req.URL.Query()
	q.Set("size", fmt.Sprintf("%f", size))
	q.Set("showId", boolToQuery(o.ShowID))
	q.Set("showCoordinates", boolToQuery(o.ShowCoordinates))
	q.Set("showPermit", boolToQuery(o.ShowPermit))
	q.Set("showInformation", boolToQuery(o.ShowInformation))
	q.Set("showPrimaryStar", boolToQuery(o.ShowPrimaryStar))
	req.URL.RawQuery = q.Encode()

	arr := make([]*SystemStruct, 0)
	sys := &arr

	_, err := c.do(req, sys)
	if err != nil {
		return nil, err
	}

	for _, s := range *sys {
		s.client = c
	}

	return *sys, nil
}
