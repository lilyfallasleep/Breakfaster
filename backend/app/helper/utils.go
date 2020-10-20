package helper

import (
	"encoding/json"
	"fmt"
)

func printJSON(data interface{}) {
	rawjson, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(rawjson))
}
