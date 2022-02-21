package convert

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func ToJSON(obj interface{}) (jsonString []byte, err error) {
	jsonString, err = json.Marshal(obj)
	if err != nil {
		log.Errorf("Error parsing JSON. " + err.Error())
		return nil, err
	}
	return
}

func FromJSON(jsonString string, obj interface{}) error {
	err := json.Unmarshal([]byte(jsonString), &obj)
	if err != nil {
		log.Errorf("Error parsing JSON. " + err.Error())
		return err
	}
	return nil
}
