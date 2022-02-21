package convert

import (
	"encoding/xml"

	log "github.com/sirupsen/logrus"
)

func ToXML(obj interface{}) (jsonString []byte, err error) {
	jsonString, err = xml.Marshal(obj)
	if err != nil {
		log.Errorf("Error parsing XML. " + err.Error())
		return nil, err
	}
	return
}

func FromXML(jsonString string, obj interface{}) error {
	err := xml.Unmarshal([]byte(jsonString), &obj)
	// if strings.ToLower(err.Error()) == "eof" {
	// 	return nil
	// }
	if err != nil {
		log.Errorf("Error parsing XML. " + err.Error())
		return err
	}
	return nil
}
