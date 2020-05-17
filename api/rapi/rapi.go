package rapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gnt-cc/config"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func NewGanetiInstance(instanceDetails CreateInstanceParameters) InstanceCreate {
	var inst InstanceCreate
	inst.InstanceName = instanceDetails.InstanceName
	inst.DiskTemplate = instanceDetails.DiskTemplate

	if len(instanceDetails.Nics) == 0 {
		inst.Nics = make([]GanetiNic, 0)
	} else {
		for _, nic := range instanceDetails.Nics {
			inst.Nics = append(inst.Nics, nic)
		}
	}

	if len(instanceDetails.Disks) == 0 {
		inst.Disks = make([]GanetiDisk, 0)
	} else {
		for _, disk := range instanceDetails.Disks {
			inst.Disks = append(inst.Disks, disk)
		}
	}

	if instanceDetails.Vcpus > 0 {
		inst.BeParams.Vcpus = instanceDetails.Vcpus
	}

	if instanceDetails.MemoryInMegabytes > 0 {
		inst.BeParams.Memory = instanceDetails.MemoryInMegabytes
	}

	inst.Version = 1
	inst.Mode = "create"
	inst.Hypervisor = "fake"
	inst.Iallocator = "hail"
	inst.OsType = "noop"
	inst.ConflictsCheck = false
	inst.IPCheck = false
	inst.NameCheck = false
	inst.NoInstall = true
	inst.WaitForSync = false
	return inst
}

func Get(clusterName string, resource string) (string, error) {
	url, netClient := getRapiConnection(clusterName)

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

func Post(clusterName string, resource string, postData interface{}) (string, error) {
	url, netClient := getRapiConnection(clusterName)

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

func getRapiConnection(clusterName string) (string, *http.Client) {
	cluster := config.GetClusterConfig(clusterName)
	var url string
	if cluster.SSL {
		url = "https://"
	} else {
		url = "http://"
	}
	url = url + fmt.Sprintf("%s:%s@%s:%d", cluster.Username, cluster.Password, cluster.Hostname, cluster.Port)

	var tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}

	return url, netClient
}
