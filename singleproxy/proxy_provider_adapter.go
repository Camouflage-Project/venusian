package singleproxy

import (
	"bytes"
	"encoding/json"
	"github.com/google/martian/v3/log"
	"io/ioutil"
	"net/http"
)

var proxyProviderUrl = "https://alealogic.com:8082/api/proxy"

type ProxyDescriptor struct {
	Host string      `json:"host"`
	Port int      	 `json:"port"`
	IpId string      `json:"ipId"`
}

func FetchProxy(apiKey string) (ProxyDescriptor, error) {
	values := map[string]string{"apiKey": apiKey}
	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post(
		proxyProviderUrl,
		"application/json",
		bytes.NewBuffer(jsonValue),
		)

	if err != nil {
		log.Errorf(err.Error())
		return ProxyDescriptor{}, err
	}else {
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
