package common

import (
	"antispam/base"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func VideoCheck(payload map[string]interface{}, apiUrl string) (string, error) {
	payload["accessKey"] = base.Conf.ShuMei.AccessKey
	b, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	return string(respBytes), nil
}

func GetVideoResult(payload map[string]interface{}, apiUrl string) (string, error) {
	payload["accessKey"] = base.Conf.ShuMei.AccessKey
	b, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	return string(respBytes), nil
}
