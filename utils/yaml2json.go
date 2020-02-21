package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

func yamlToJson(filename string) []map[string]interface{} {
	file, err := ioutil.ReadFile(filename)
	checkError(err)

	data, err := yaml.YAMLToJSON(file)
	checkError(err)

	var results []map[string]interface{}
	json.Unmarshal(data, &results)
	return results
}

func checkError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func GetConfigurations() []map[string]interface{} {
	configFile := os.Getenv("CONFIG_FILE")
	config := yamlToJson(configFile)
	return config
}
