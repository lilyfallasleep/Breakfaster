package helper

import (
	"encoding/json"
	"fmt"
)

// PrintJSON helps printing out object in JSON
func PrintJSON(data interface{}) {
	rawjson, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(rawjson))
}
