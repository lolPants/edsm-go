package edsm

// ServerStatus returns Elite Dangerous server status information
func (c *Client) ServerStatus() (*EliteServerStatus, error) {
	req, err := c.newGetRequest("/api-status-v1/elite-server")
	if err != nil {
		return nil, err
	}

	body := &EliteServerStatus{}
	_, err = c.do(req, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
