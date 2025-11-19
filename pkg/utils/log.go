package utils

import (
	"encoding/json"
	"fmt"
)

func LogPretty(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}
	fmt.Printf("%s\n", jsonData)
}
