package network

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

// PostForm : PostForm values thru HTTP PostForm
func PostForm(url string, data url.Values, headers map[string]string, options ...string) (string, error) {
	log.Debugf("Request: %s\n", url)
	client := &http.Client{}
	if len(options) != 0 && options[0] == "true" {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	} else {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}
	r, err := http.NewRequest("POST", url, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Error(err)
		return "", err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	for key, value := range headers {
		r.Header.Add(key, value)
	}

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

// PostForm : PostForm values thru HTTP PostForm
func PostJSON(url string, data []byte, headers map[string]string, uniqueID string, options ...string) (string, error) {
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
	for key, value := range headers {
		r.Header.Add(key, value)
	}
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
		log.Debugf("Error Response Code : ", res.StatusCode)
		err := fmt.Errorf("%d", res.StatusCode)
		log.SetReportCaller(false)
		return string(body), err
	}
	log.SetReportCaller(false)
	return string(body), nil
}
