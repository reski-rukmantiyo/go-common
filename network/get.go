package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// GetWithHeaders : Get values thru HTTP Get
// 1st Option = JSON or not
func GetWithHeaders(url string, headers map[string]string, options ...string) (string, error) {
	log.SetReportCaller(true)
	log.Debugf("Request: %s\n", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	if strings.ToLower(options[0]) == "json" {
		req.Header.Add("Content-Type", "application/json")
	}
	if err != nil {
		log.Debugf("Error : ", err.Error())
		log.SetReportCaller(false)
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		log.Debugf("Error : ", err.Error())
		log.SetReportCaller(false)
		return "", err
	}
	content, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Debugf("Error : ", err.Error())
		log.SetReportCaller(false)
		return "", err
	}
	log.Debugf("Response: " + string(content))
	if res.StatusCode != 200 && res.StatusCode != 201 {
		err := fmt.Errorf("%d", res.StatusCode)
		log.Debugf("Error Response Code : ", res.StatusCode)
		log.SetReportCaller(false)
		return string(content), err
	}
	log.SetReportCaller(false)
	return string(content), nil
}

// Get : Get values thru HTTP Get
func Get(url string) (string, error) {
	log.Debugf("Request: %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return "", err
	}
	content, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Error(err)
		return "", err
	}
	log.Debugf("Response: " + string(content))
	if res.StatusCode != 200 && res.StatusCode != 201 {
		err := fmt.Errorf("%d", res.StatusCode)
		return string(content), err
	}
	return string(content), nil
}

// GetWithOptions : Get Function with a lot of customizations
// url : for address
// headers : headers to be implement in the GET
// Options[0] : for content typeb
// Options[1] : for timeout
func GetWithOptions(url string, headers map[string]string, options ...string) (string, error) {
	var timeSecond time.Duration
	timeSecond = 5
	if options[1] != "" {
		timeFromString, err := strconv.Atoi(options[1])
		if err == nil {
			timeSecond = time.Duration(timeFromString)
		}
	}
	log.SetReportCaller(true)
	log.Debugf("Request: %s\n", url)
	client := &http.Client{
		Timeout: timeSecond * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	if strings.ToLower(options[0]) == "json" {
		req.Header.Add("Content-Type", "application/json")
	}
	if err != nil {
		log.Debugf("Error : ", err.Error())
		log.SetReportCaller(false)
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		log.Debugf("Error : ", err.Error())
		log.SetReportCaller(false)
		return "", err
	}
	content, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Debugf("Error : ", err.Error())
		log.SetReportCaller(false)
		return "", err
	}
	log.Debugf("Response: " + string(content))
	if res.StatusCode != 200 && res.StatusCode != 201 {
		err := fmt.Errorf("%d", res.StatusCode)
		log.Debugf("Error Response Code : ", res.StatusCode)
		log.SetReportCaller(false)
		return string(content), err
	}
	log.SetReportCaller(false)
	return string(content), nil
}
