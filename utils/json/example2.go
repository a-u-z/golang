package json

import (
	"encoding/json"
	"log"
)

func getExample2(b []byte) {
	m := map[string]map[string][]string{}
	json.Unmarshal(b, &m)
	log.Printf("here is m:%+v", m)
}
