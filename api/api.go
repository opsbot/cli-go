package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mitchellh/mapstructure"
	"github.com/opsbot/cli-go/utils"
	log "github.com/sirupsen/logrus"
)

var (
	// Config - hold a reference to configuration
	Config Configuration
)

func getEndpoint(path string) string {
	CheckConfig()
	var url string
	url = fmt.Sprintf("%v%v", Config.Host, path)
	log.Tracef("getEndpoint %v\n", url)
	return url
}

func showFullResponseBody(resp APIResponse) {
	var data map[string]interface{}
	json.Unmarshal([]byte(resp.Body), &data)
	utils.PrettyPrint(data)
}

// CheckConfig - check config is set
func CheckConfig() {
	// check if CurrentSession is an empty struct
	if (Configuration{}) == Config {
		log.Fatalln("Config not set: call SetEnvironment")
	}
}

// Request - execure api request
func Request(method string, url string, body []byte) APIResponse {
	client := &http.Client{}

	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(body))
	if reqErr != nil {
		log.Fatalf("HTTPRequestError: %v", reqErr)
	}
	req.SetBasicAuth(Config.APIKey, Config.APISecret)

	resp, resErr := client.Do(req)
	if resErr != nil {
		log.Fatalf("HTTPResponseError: %v", resErr)
	}
	defer resp.Body.Close()

	log.Info(resp.Status)

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadError: %v", err)
	}

	return APIResponse{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Body:       respBody,
	}
}

// SetEnvironment - set environment config
func SetEnvironment(cfg map[string]interface{}) {
	Config = Configuration{}
	mapstructure.Decode(cfg, &Config)
}
