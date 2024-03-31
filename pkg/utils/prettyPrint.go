package utils

import (
	"encoding/json"
	"fmt"
)

func PPrint(data any) string {

	info, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println("json Marshal Error")
	}
	return string(info)
}
