package json

import (
	"encoding/json"
	"log"
)

type example1 struct {
	JobTitle string `json:"job_title"`
	Age      string `json:"age"`
}

func getExample1(b []byte) {
	m := map[string]example1{}
	json.Unmarshal(b, &m)
	log.Printf("here is m:%+v", m)
}
