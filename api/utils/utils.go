package utils

import "os"

func IsInSlice(needle string, list []string) bool {
	for _, entry := range list {
		if entry == needle {
			return true
		}
	}
	return false
}

func FileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
