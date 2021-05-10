package dsl

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ghodss/yaml"
)

func ParseApiSpec(filename string) (*Context, error) {
	var apiSpec *ApiSpec
	apiSpecFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(apiSpecFile)
	if err != nil {
		return nil, err
	}
	if strings.HasSuffix(filename, ".json") {
		apiSpec, err = parseJsonApiSpec(content)
	} else {
		apiSpec, err = parseYamlApiSpec(content)
	}
	if err != nil {
		return nil, err
	}
	return apiSpec.ToContext()

}

func parseYamlApiSpec(content []byte) (*ApiSpec, error) {
	apiSpec := &ApiSpec{}
	err := yaml.Unmarshal(content, apiSpec)
	if err != nil {
		return nil, err
	}
	return apiSpec, nil
}

func parseJsonApiSpec(content []byte) (*ApiSpec, error) {
	apiSpec := &ApiSpec{}
	err := json.Unmarshal(content, apiSpec)
	if err != nil {
		return nil, err
	}
	return apiSpec, nil
}
