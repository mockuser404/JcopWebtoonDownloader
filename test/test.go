package main

import (
	"encoding/json"
	"log"
)

func main() {
	data := `[{"id":2329,"episode":1,"title":"교수인형 예고편"}]`
	type EachFiles struct {
		Id int `json:"id"`
	}

	var outputFiles []EachFiles
	json.Unmarshal([]byte(data), &outputFiles)
	log.Println(outputFiles)
}
