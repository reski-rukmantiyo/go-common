package tools

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/olivere/elastic/v7"
	"github.com/reski-rukmantiyo/go-common/config"
	logger "github.com/sirupsen/logrus"
)

func getUniqueID() string {
	uniqueID, _ := uuid.NewV4()
	return fmt.Sprintf("%s", uniqueID)
}

type Fields struct {
	Apps     string `json:"Apps"`
	ID       string `json:"Id"`
	Key      string `json:"Key"`
	Request  string `json:"Request"`
	Response string `json:"Response"`
}

type AppLog struct {
	Fields          Fields    `json:"fields"`
	Level           string    `json:"level"`
	Message         string    `json:"message"`
	MessageTemplate string    `json:"messageTemplate"`
	Timestamp       time.Time `json:"timestamp"`
}

type ES struct {
	username string
	password string
	hosts    string
	client   *elastic.Client
	err      error
	name     string
	caFile   string
}

func NewLogWithCA(hosts, name, username, password, caFile string) *ES {
	es := ES{
		hosts:    hosts,
		name:     name,
		username: username,
		password: password,
		caFile:   caFile,
	}
	client, err := es.getESClient()
	if err != nil {
		client = nil
		es.err = err
	}
	es.client = client
	return &es
}

func NewLog(hosts, name, username, password string) *ES {
	es := ES{
		hosts:    hosts,
		name:     name,
		username: username,
		password: password,
	}
	client, err := es.getESClient()
	if err != nil {
		client = nil
		es.err = err
	}
	es.client = client
	return &es
}

func (es *ES) getESClient() (*elastic.Client, error) {
	hosts := strings.Split(es.hosts, ",")
	var client *elastic.Client
	var err error
	if es.caFile != "" && es.username != "" && es.password != "" {
		// Get the SystemCertPool, continue with an empty pool on error
		rootCAs, _ := x509.SystemCertPool()
		if rootCAs == nil {
			rootCAs = x509.NewCertPool()
		}
		// Read in the cert file
		path, err := config.GetPath()
		if err != nil {
			return nil, err
		}
		certs, err := ioutil.ReadFile(path + es.caFile)
		if err != nil {
			log.Fatalf("Failed to append %q to RootCAs: %v", es.caFile, err)
		}
		// Append our cert to the system pool
		if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
			log.Println("No certs appended, using system certs only")
		}
		// Trust the augmented cert pool in our client
		tlsConfig := &tls.Config{
			RootCAs: rootCAs,
		}
		tlsConfig.BuildNameToCertificate()
		transport := &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSClientConfig:     tlsConfig,
			TLSHandshakeTimeout: 5 * time.Second,
		}
		httpClient := &http.Client{
			Transport: transport,
			Timeout:   time.Second * 10,
		}
		client, err = elastic.NewClient(
			elastic.SetURL(hosts...),
			elastic.SetSniff(false),
			elastic.SetHealthcheckInterval(10*time.Second),
			elastic.SetMaxRetries(5),
			elastic.SetErrorLog(log.New(os.Stderr, "ES Error: ", log.LstdFlags)),
			elastic.SetInfoLog(log.New(os.Stdout, "ES Info: ", log.LstdFlags)),
			elastic.SetHttpClient(httpClient),
			elastic.SetBasicAuth(es.username, es.password),
		)
		if err != nil {
			// client = nil
			es.err = err
			return nil, err
		}
	} else if es.username != "" && es.password != "" {
		client, err = elastic.NewClient(
			elastic.SetURL(hosts...),
			elastic.SetSniff(false),
			elastic.SetHealthcheckInterval(10*time.Second),
			elastic.SetMaxRetries(5),
			elastic.SetErrorLog(log.New(os.Stderr, "ES Error: ", log.LstdFlags)),
			elastic.SetInfoLog(log.New(os.Stdout, "ES Info: ", log.LstdFlags)),
			elastic.SetBasicAuth(es.username, es.password),
		)
		if err != nil {
			// client = nil
			es.err = err
			return nil, err
		}
	} else if es.caFile != "" {
		// Get the SystemCertPool, continue with an empty pool on error
		rootCAs, _ := x509.SystemCertPool()
		if rootCAs == nil {
			rootCAs = x509.NewCertPool()
		}
		// Read in the cert file
		path, err := config.GetPath()
		if err != nil {
			return nil, err
		}
		certs, err := ioutil.ReadFile(path + es.caFile)
		if err != nil {
			log.Fatalf("Failed to append %q to RootCAs: %v", es.caFile, err)
		}
		// Append our cert to the system pool
		if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
			log.Println("No certs appended, using system certs only")
		}
		// Trust the augmented cert pool in our client
		tlsConfig := &tls.Config{
			RootCAs: rootCAs,
		}
		tlsConfig.BuildNameToCertificate()
		transport := &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSClientConfig:     tlsConfig,
			TLSHandshakeTimeout: 5 * time.Second,
		}
		httpClient := &http.Client{
			Transport: transport,
			Timeout:   time.Second * 10,
		}
		client, err = elastic.NewClient(
			elastic.SetURL(hosts...),
			elastic.SetSniff(false),
			elastic.SetHealthcheckInterval(10*time.Second),
			elastic.SetMaxRetries(5),
			elastic.SetErrorLog(log.New(os.Stderr, "ES Error: ", log.LstdFlags)),
			elastic.SetInfoLog(log.New(os.Stdout, "ES Info: ", log.LstdFlags)),
			elastic.SetHttpClient(httpClient),
		)
		if err != nil {
			// client = nil
			es.err = err
			return nil, err
		}
	} else {
		client, err = elastic.NewClient(
			elastic.SetURL(hosts...),
			elastic.SetSniff(false),
			elastic.SetHealthcheckInterval(10*time.Second),
			elastic.SetMaxRetries(5),
			elastic.SetErrorLog(log.New(os.Stderr, "ES Error: ", log.LstdFlags)),
			elastic.SetInfoLog(log.New(os.Stdout, "ES Info: ", log.LstdFlags)),
		)
		if err != nil {
			// client = nil
			es.err = err
			return nil, err
		}
	}
	return client, nil
}

func (es *ES) WriteInfoLog(key, request, message, response string) {
	id := getUniqueID()
	info := "info"
	if err := es.writeLog(id, key, request, info, message, response); err != nil {
		logger.Errorf("%s", err.Error())
	}
	logger.Infof("%s: %s", key, message)
}

func (es *ES) WriteDebugLog(key, request, message, response string) {
	id := getUniqueID()
	info := "debug"
	if err := es.writeLog(id, key, request, info, message, response); err != nil {
		logger.Errorf("%s", err.Error())
	}
	logger.Debugf("%s: %s", key, message)
}

func (es *ES) WriteSimpleErrorLog(key, message string) {
	id := getUniqueID()
	info := "error"
	if err := es.writeLog(id, key, "", info, message, ""); err != nil {
		logger.Errorf("%s", err.Error())
	}
	logger.Errorf("%s: %s", key, message)
}

func (es *ES) WriteErrorLog(key, request, message, response string) {
	id := getUniqueID()
	info := "error"
	if err := es.writeLog(id, key, request, info, message, response); err != nil {
		logger.Errorf("%s", err.Error())
	}
	logger.Errorf("%s: %s", key, message)
}

func (es *ES) writeLog(id, key, request, info, message, response string) error {
	if es.err != nil {
		client, err := es.getESClient()
		if err != nil {
			err := fmt.Errorf("WriteLog: %s\n", err.Error())
			return err
		}
		es.client = client
	}
	appLog := AppLog{
		Fields: Fields{
			Apps:     es.name,
			ID:       id,
			Key:      key,
			Request:  request,
			Response: response,
		},
		Level:           info,
		Message:         message,
		MessageTemplate: "{Apps}{ID}{Key}{Request}{Response}",
		Timestamp:       time.Now(),
	}

	ctx := context.Background()
	esclient := es.client
	dataJSON, err := json.Marshal(appLog)
	if err != nil {
		err := fmt.Errorf("WriteLog: %s\n", err.Error())
		return err
	}
	js := string(dataJSON)
	indexName := fmt.Sprintf("%s-%s", strings.ToLower(appLog.Fields.Apps), time.Now().Format("2006-01-02"))
	_, err = esclient.Index().
		Index(indexName).
		Id(id).
		BodyJson(js).
		Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
