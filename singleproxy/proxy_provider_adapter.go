package singleproxy

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/martian/v3/log"
	"io/ioutil"
	"net/http"
	"strconv"
)

//var baseProxyProviderUrl = "https://alealogic.com:8082/api/"
var baseProxyProviderUrl = "http://localhost:80"
var client = &http.Client{}

type ProxyDescriptor struct {
	Host string      `json:"host"`
	Port int      	 `json:"port"`
	IpId string      `json:"ipId"`
}

func FetchProxy(apiKey string) (ProxyDescriptor, error) {
	req, err := http.NewRequest("GET", baseProxyProviderUrl + "/proxy-descriptor", nil)
	req.Header.Add("key", apiKey)

	resp, err := client.Do(req)

	if err != nil {
		log.Errorf(err.Error())
		return ProxyDescriptor{}, err
	}else if resp.StatusCode == 401 {
		log.Errorf(strconv.Itoa(resp.StatusCode))
		return ProxyDescriptor{}, errors.New("API key unauthorized")
	} else {
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Errorf(err.Error())
		}

		data := ProxyDescriptor{}
		b := []byte(string(body))
		err = json.Unmarshal(b, &data)
		if err != nil {
			log.Errorf(err.Error())
		}

		return data, nil
	}
}

func ReportMalfunctioningProxy(ipId string) {
	values := map[string]string{"ipId": ipId}
	jsonValue, _ := json.Marshal(values)

	_, err := http.Post(
		baseProxyProviderUrl + "failed-proxy",
		"application/json",
		bytes.NewBuffer(jsonValue),
	)

	if err != nil {
		log.Errorf(err.Error())
	}
}
