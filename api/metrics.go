package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"dynamic-pv-scaling/logger"
	log "github.com/sirupsen/logrus"
)

var (
	prometheusURL string
)

// JSONResponse is a struct for json with the below mentioned keys
type JSONResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Endpoint              string `json:"endpoint"`
				Instance              string `json:"instance"`
				Job                   string `json:"job"`
				Namespace             string `json:"namespace"`
				Node                  string `json:"node"`
				Persistentvolumeclaim string `json:"persistentvolumeclaim"`
				Service               string `json:"service"`
			} `json:"metric"`
			Value []string `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

// PersistentVolumeList is a struct for json with the below mentioned keys
type PersistentVolumeList struct {
	PeristentVolumeName string `json:"persistent_volume_name"`
	Namespace           string `json:"namespace"`
	Value               int    `json:"value"`
}

// PersistentVolumeUsage struct for a json with below mentioned keys
type PersistentVolumeUsage struct {
	PeristentVolumeName string `json:"persistent_volume_name"`
	Namespace           string `json:"namespace"`
	Value               int    `json:"value"`
}

// GetPersistentVolumeList takes namespace and PV name as input and returns persistent volume list  */
func GetPersistentVolumeList(nameSpace string, persistentVolumeName string) PersistentVolumeList {
	var qeuryResponse JSONResponse
	var pvList PersistentVolumeList
	logger.LogStdout()

	resp := GetVolumeListQueryResponse(nameSpace, persistentVolumeName)
	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"kubeconfig":        kubeConfig,
			"namespace":         nameSpace,
			"persistent_volume": persistentVolumeName,
		}).Error(err)
	}

	json.Unmarshal(output, &qeuryResponse)

	for _, queryOutput := range qeuryResponse.Data.Result {
		finalValue, _ := strconv.ParseFloat(strings.Join(queryOutput.Value, ""), 64)
		pvLists := PersistentVolumeList{
			PeristentVolumeName: queryOutput.Metric.Persistentvolumeclaim,
			Namespace:           queryOutput.Metric.Namespace,
			Value:               int(finalValue),
		}
		pvList = pvLists
	}
	return pvList
}

// GetPeristentVolumeUsage takes namespace and PV name as input and return persistent volume usage  */
func GetPeristentVolumeUsage(nameSpace string, persistentVolumeName string) PersistentVolumeUsage {
	var qeuryResponse JSONResponse
	var pvList PersistentVolumeUsage
	logger.LogStdout()

	resp := GetVolumeUsageQueryResponse(nameSpace, persistentVolumeName)
	output, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.WithFields(log.Fields{
			"kubeconfig":             kubeConfig,
			"namespace":              nameSpace,
			"persistent_volume_name": persistentVolumeName,
		}).Error(err)
	}

	json.Unmarshal(output, &qeuryResponse)
	for _, queryOutput := range qeuryResponse.Data.Result {
		finalValue, _ := strconv.Atoi(strings.Join(queryOutput.Value, ""))
		gbValue := finalValue/1024/1024/1024 + 1
		pvLists := PersistentVolumeUsage{
			PeristentVolumeName: queryOutput.Metric.Persistentvolumeclaim,
			Namespace:           queryOutput.Metric.Namespace,
			Value:               gbValue,
		}
		pvList = pvLists
	}
	return pvList
}

// GetVolumeListQueryResponse function takes namespace and PV volume name as input and returns list of volumes via prometheus */
func GetVolumeListQueryResponse(nameSpace string, persistentVolumeName string) *http.Response {
	logger.LogStdout()

	req := GenerateVolumeListQuery(nameSpace, persistentVolumeName)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.WithFields(log.Fields{
			"kubeconfig":        kubeConfig,
			"namespace":         nameSpace,
			"persistent_volume": persistentVolumeName,
		}).Error(err)
	}

	log.WithFields(log.Fields{
		"kubeconfig":        kubeConfig,
		"namespace":         nameSpace,
		"persistent_volume": persistentVolumeName,
	}).Info("Successfully connected with prometheus at " + prometheusURL + " to list persistent volume")

	return resp
}

// GetVolumeUsageQueryResponse function takes namespace and PV volume name as input and returns Volume Usage via prometheus */
func GetVolumeUsageQueryResponse(nameSpace string, persistentVolumeName string) *http.Response {
	logger.LogStdout()

	req := GenerateVolumeUsageQuery(nameSpace, persistentVolumeName)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.WithFields(log.Fields{
			"kubeconfig":        kubeConfig,
			"namespace":         nameSpace,
			"persistent_volume": persistentVolumeName,
		}).Error(err)
	}

	log.WithFields(log.Fields{
		"kubeconfig":        kubeConfig,
		"namespace":         nameSpace,
		"persistent_volume": persistentVolumeName,
	}).Info("Successfully connected with prometheus at " + prometheusURL + " to get persistent volume usage")

	return resp
}

// GenerateVolumeListQuery function generates the query to extract volume list from prometheus and returns the same
func GenerateVolumeListQuery(nameSpace string, persistentVolumeName string) *http.Request {
	logger.LogStdout()
	prometheusURL = os.Getenv("PROMETHEUS_URL") + "/api/v1/query"

	body := strings.NewReader(`query=100 * (kubelet_volume_stats_available_bytes{namespace="` + nameSpace + `",persistentvolumeclaim="` + persistentVolumeName + `"} / kubelet_volume_stats_capacity_bytes{namespace="` + nameSpace + `",persistentvolumeclaim="` + persistentVolumeName + `"})`)
	req, err := http.NewRequest("POST", prometheusURL, body)

	if err != nil {
		log.WithFields(log.Fields{
			"kubeconfig":        kubeConfig,
			"namespace":         nameSpace,
			"persistent_volume": persistentVolumeName,
		}).Error(err)
	}

	log.WithFields(log.Fields{
		"kubeconfig":        kubeConfig,
		"namespace":         nameSpace,
		"persistent_volume": persistentVolumeName,
	}).Info("Successfully created the query for prometheus to list persistent volumes")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("sky", "sky")
	return req
}

// GenerateVolumeUsageQuery function generates the volume usage query to extract usage info from prometheus
func GenerateVolumeUsageQuery(nameSpace string, persistentVolumeName string) *http.Request {
	logger.LogStdout()
	prometheusURL = os.Getenv("PROMETHEUS_URL") + "/api/v1/query"

	body := strings.NewReader(`query=kubelet_volume_stats_capacity_bytes{namespace="` + nameSpace + `",persistentvolumeclaim="` + persistentVolumeName + `"}`)
	req, err := http.NewRequest("POST", prometheusURL, body)

	if err != nil {
		log.WithFields(log.Fields{
			"kubeconfig":        kubeConfig,
			"namespace":         nameSpace,
			"persistent_volume": persistentVolumeName,
		}).Error(err)
	}

	log.WithFields(log.Fields{
		"kubeconfig":        kubeConfig,
		"namespace":         nameSpace,
		"persistent_volume": persistentVolumeName,
	}).Info("Successfully created the query for prometheus to check volume usage")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req
}
