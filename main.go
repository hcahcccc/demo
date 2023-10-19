package main

import (
	"demo/config"
	"encoding/json"
	"fmt"
)

func main() {
	reConfig, err := config.NewReConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	marshal, _ := json.Marshal(reConfig)
	fmt.Println(string(marshal))
}
