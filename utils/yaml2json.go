package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)
	/* 
	Converts given YAML file to json
	Takes filename as input and returns results as interface.
	*/
func yamlToJson(filename string) []map[string]interface{} {
	

	file, err := ioutil.ReadFile(filename)
	checkError(err)

	data, err := yaml.YAMLToJSON(file)
	checkError(err)

	var results []map[string]interface{}
	json.Unmarshal(data, &results)
	return results
}

	// Checks for any error
func checkError(err error) bool {
	

	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

	// Gets configFile location , converts it to json and returns the same.
func GetConfigurations() []map[string]interface{} {
 

	configFile := os.Getenv("CONFIG_FILE")
	config := yamlToJson(configFile)
	return config
}
