package utils

func IsInSlice(needle string, list []string) bool {
	for _, entry := range list {
		if entry == needle {
			return true
		}
	}
	return false
}
