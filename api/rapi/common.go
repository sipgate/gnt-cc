package rapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gnt-cc/config"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func Get(clusterConfig config.ClusterConfig, resource string) (string, error) {
	url, netClient := getRapiConnection(clusterConfig)

	log.Infof("RAPI GET %s", resource)
	response, err := netClient.Get(url + resource)
	if err != nil {
		log.Errorf("HTTP RAPI Connect: %s", err)
		return "", err
	}

	if response.StatusCode != 200 {
		log.Errorf("HTTP RAPI Bad Status Code: %d", response.StatusCode)
		return "", fmt.Errorf("Bad RAPI Status Code: %d", response.StatusCode)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("HTTP RAPI Reading Body: %s", err)
		return "", err
	}

	return string(body), err
}

func Post(clusterConfig config.ClusterConfig, resource string, postData interface{}) (string, error) {
	url, netClient := getRapiConnection(clusterConfig)

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return "", fmt.Errorf("Could not prepare JSON for RAPI Request: %s", err)
	}

	log.Infof("RAPI POST %s", resource)
	response, err := netClient.Post(url+resource, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Errorf("HTTP RAPI Connect: %s", err)
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("HTTP RAPI Reading Body: %s", err)
		return "", err
	}

	if response.StatusCode != 200 {
		log.Errorf("HTTP RAPI Bad Status Code: %d", response.StatusCode)
		log.Errorf("Answer: %s", body)
		return "", fmt.Errorf("Bad RAPI Status Code: %d", response.StatusCode)
	}

	return string(body), err
}

func getRapiConnection(clusterConfig config.ClusterConfig) (string, *http.Client) {
	var protocol string
	if clusterConfig.SSL {
		protocol = "https"
	} else {
		protocol = "http"
	}

	url := fmt.Sprintf("%s://%s:%s@%s:%d", protocol, clusterConfig.Username, clusterConfig.Password, clusterConfig.Hostname, clusterConfig.Port)

	var transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	return url, netClient
}
