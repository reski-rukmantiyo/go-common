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
func Get(url string) string {
	log.Debugf("Request: %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		log.Error(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Error(err)
	}
	log.Debugf("Response: " + string(content))
	return string(content)
}

// PostForm : PostForm values thru HTTP PostForm
func PostJSON(url string, data []byte, uniqueID string) (string, error) {
	log.Debugf("Request: %s\n", url)
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 		InsecureSkipVerify: true,
	// 	},
	// }
	client := &http.Client{
		// Transport: tr,
	}
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		// WriteErrorLog(uniqueID, "PostJSON", fmt.Sprintf("Request: %s|Response: %s", string(data), err.Error()), "")
		return "", err
	}
	r.Header.Add("Content-Type", "application/json")
	// r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		// WriteErrorLog(uniqueID, "PostJSON", fmt.Sprintf("Request: %s|Response: %s", string(data), err.Error()), "")
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// WriteErrorLog(uniqueID, "PostJSON", fmt.Sprintf("Request: %s|Response: %s", string(data), err.Error()), "")
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
