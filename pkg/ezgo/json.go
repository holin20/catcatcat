package ezgo

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

func ExtractJsonPath(jsonStr string, jsonPath string) (*gjson.Result, error) {
	result := gjson.Get(jsonStr, jsonPath)
	if !result.Exists() {
		return nil, fmt.Errorf("error looking up JSON path: %v", jsonPath)
	}

	return &result, nil
}

func GetFloatFromJSONPath(jsonStr string, jsonPath string) (float64, error) {
	result := gjson.Get(jsonStr, jsonPath)
	if !result.Exists() {
		return 0, fmt.Errorf("error looking up JSON path: %v", jsonPath)
	}

	return result.Float(), nil
}

func ToJsonString(v any) string {
	return string(Arg1(json.MarshalIndent(v, "", "  ")))
}
