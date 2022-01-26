package tools

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

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

// PostForm : PostForm values thru HTTP PostForm
func PostJSON(url string, data []byte, uniqueID string) (string, error) {
	log.SetReportCaller(true)
	log.Debugf("Request: %s\n", url)
	client := &http.Client{}
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Debugf("Error : ", err.Error())
		log.SetReportCaller(false)
		return "", err
	}
	r.Header.Add("Content-Type", "application/json")
	res, err := client.Do(r)
	if err != nil {
		log.Debugf("Error : ", err.Error())
		log.SetReportCaller(false)
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Debugf("Error : ", err.Error())
		log.SetReportCaller(false)
		return "", err
	}
	log.Debugf("Response: " + string(body))
	if res.StatusCode != 200 && res.StatusCode != 201 {
		log.Debugf("Error : ", err.Error())
		err := fmt.Errorf("%d", res.StatusCode)
		log.SetReportCaller(false)
		return string(body), err
	}
	log.SetReportCaller(false)
	return string(body), nil
}

// PostForm : PostForm values thru HTTP PostForm
func PostForm(url string, data url.Values) (string, error) {
	log.Debugf("Request: %s\n", url)
	client := &http.Client{}
	r, err := http.NewRequest("POST", url, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Error(err)
		return "", err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		log.Error(err)
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return "", err
	}
	log.Debugf("Response: " + string(body))
	if res.StatusCode != 200 && res.StatusCode != 201 {
		err := fmt.Errorf("%d", res.StatusCode)
		return string(body), err
	}
	return string(body), nil
}
