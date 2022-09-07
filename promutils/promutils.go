/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package promutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// QueryInputs is definition for inputs for Prometheus query
type QueryInputs struct {
	Name             string
	PersistentVolume string
	Namespace        string
	Body             *strings.Reader
}

// PrometheusResponse defines the query output response from Prometheus
type PrometheusResponse struct {
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

// PersistentVolumeResponse defines the output response for PersistentVolume
type PersistentVolumeResponse struct {
	Namespace        string `json:"namespace"`
	PersistentVolume string `json:"persistent_volume_name"`
	Value            int    `json:"int"`
}

// GetPersistentVolumeUsagePercentage is a method to get usage percentage of PersistentVolume
func GetPersistentVolumeUsagePercentage(params QueryInputs) (*PersistentVolumeResponse, error) {
	query := fmt.Sprintf("100 * (kubelet_volume_stats_available_bytes{namespace=\"%s\",persistentvolumeclaim=\"%s\"} / kubelet_volume_stats_capacity_bytes{namespace=\"%s\",persistentvolumeclaim=\"%s\"})", params.Namespace, params.PersistentVolume, params.Namespace, params.PersistentVolume)
	queryParams := QueryInputs{
		Body: strings.NewReader(query),
	}
	response, err := executePrometheusQuery(queryParams)
	if err != nil {
		return nil, err
	}
	return response, err
}

// GetPersistentVolumeTotalCapacity is a method to get capacity of PersistentVolume
func GetPersistentVolumeTotalCapacity(params QueryInputs) (*PersistentVolumeResponse, error) {
	query := fmt.Sprintf("query=kubelet_volume_stats_capacity_bytes{namespace=\"%s\",persistentvolumeclaim=\"%s\"} / 1024 / 1024 / 1024 + 1", params.Namespace, params.PersistentVolume)
	queryParams := QueryInputs{
		Body: strings.NewReader(query),
	}
	response, err := executePrometheusQuery(queryParams)
	if err != nil {
		return nil, err
	}
	return response, err
}

// executePrometheusQuery is a method for executing Prometheus query
func executePrometheusQuery(params QueryInputs) (*PersistentVolumeResponse, error) {
	var qeuryResponse PrometheusResponse
	prometheusURL := generatePrometheusURL()
	req, err := http.NewRequest("POST", prometheusURL, params.Body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(output, &qeuryResponse)
	if err != nil {
		return nil, err
	}
	dataValue, err := strconv.ParseFloat(qeuryResponse.Data.Result[0].Value[1], 64)
	if err != nil {
		return nil, err
	}
	persistentVolumeInfo := &PersistentVolumeResponse{
		Namespace:        qeuryResponse.Data.Result[0].Metric.Namespace,
		PersistentVolume: qeuryResponse.Data.Result[0].Metric.Persistentvolumeclaim,
		Value:            int(dataValue),
	}
	return persistentVolumeInfo, nil
}

// generatePrometheusURL is a method to generate Prometheus URL
func generatePrometheusURL() string {
	prometheusURL := fmt.Sprintf("%s%s", os.Getenv("PROMETHEUS_URL"), "/api/v1/query")
	return prometheusURL
}
