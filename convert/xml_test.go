package convert

import (
	"encoding/xml"
	"testing"

	"github.com/go-playground/assert/v2"
)

type Response struct {
	XMLName  xml.Name `xml:"API"`
	Text     string   `xml:",chardata"`
	Version  string   `xml:"version,attr"`
	Response struct {
		Text      string `xml:",chardata"`
		Operation struct {
			Text   string `xml:",chardata"`
			Name   string `xml:"name,attr"`
			Result struct {
				Text       string `xml:",chardata"`
				Statuscode string `xml:"statuscode"`
				Status     string `xml:"status"`
				Message    string `xml:"message"`
			} `xml:"result"`
		} `xml:"operation"`
	} `xml:"response"`
}

func TestToFromXML(t *testing.T) {
	response := &Response{
		Version: "1.0",
	}
	response.Response.Operation.Result.Statuscode = "200"
	response.Response.Operation.Result.Status = "Success"
	stringXML, err := ToXML(response)
	if err != nil {
		assert.Equal(t, err, nil)
	}
	assert.NotEqual(t, string(stringXML), "")
	objectXML := &Response{}
	err = FromXML(string(stringXML), objectXML)
	if err != nil {
		assert.Equal(t, err, nil)
	}
	assert.Equal(t, objectXML.Version, response.Version)
}
