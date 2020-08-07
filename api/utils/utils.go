package utils

import "encoding/json"

func IsInSlice(needle string, list []string) bool {
	for _, entry := range list {
		if entry == needle {
			return true
		}
	}
	return false
}

func ConvertMapToStruct(source map[string]interface{}, target interface{}) error {
	data, err := json.Marshal(source)

	if err != nil {
		return err
	}

	return json.Unmarshal(data, target)
}
