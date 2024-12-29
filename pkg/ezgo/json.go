package ezgo

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func GetValueFromJSONPath(jsonStr string, jsonPath string) (interface{}, error) {
	result := gjson.Get(jsonStr, jsonPath)
	if !result.Exists() {
		return nil, fmt.Errorf("error looking up JSON path: %v", jsonPath)
	}

	return result.Value(), nil
}

func GetFloatFromJSONPath(jsonStr string, jsonPath string) (float64, error) {
	result := gjson.Get(jsonStr, jsonPath)
	if !result.Exists() {
		return 0, fmt.Errorf("error looking up JSON path: %v", jsonPath)
	}

	return result.Float(), nil
}
