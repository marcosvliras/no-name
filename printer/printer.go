package printer

import (
	"encoding/json"
	"fmt"
)

func StdoutPrint(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(data))
}
