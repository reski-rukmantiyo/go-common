package tools

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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

// PostForm : PostForm values thru HTTP PostForm
func PostJSON(url string, data []byte, uniqueID string, options ...string) (string, error) {
	log.SetReportCaller(true)
	log.Debugf("Request: %s\n", url)
	if options[0] == "true" {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	} else {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}
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
