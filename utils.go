package edsm

func boolToQuery(value bool) string {
	if value == true {
		return "1"
	}

	return "0"
}
