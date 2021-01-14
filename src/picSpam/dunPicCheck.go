package picSpam

import (
	"antispam/base"
	"antispam/common"
	"antispam/models"
	"encoding/json"
	"net/url"
)

type DunImageCheckLabels struct {
	Label int64   `json:"label"`
	Level int64   `json:"level"`
	Rate  float64 `json:"rate"`
}

type DunImageCheckAntispam struct {
	TaskId     string                `json:"taskId"`
	Status     int64                 `json:"status"`
	Action     int64                 `json:"action"`
	CensorType int64                 `json:"censorType"`
	Name       string                `json:"name"`
	Labels     []DunImageCheckLabels `json:"labels"`
}

type DunImageCheckResponse struct {
	Code     int64                   `json:"code"`
	Msg      string                  `json:"msg"`
	Antispam []DunImageCheckAntispam `json:"antispam"`
}

type Image struct {
	Name string `json:"name"`
	Type int64  `json:"type"`
	Data string `json:"data"`
}

func DunImageCheck(images []models.Data, checkLabels []string) (DunImageCheckResponse, error) {
	var imageArr []Image
	for m := range images {
		var imageData Image
		imageData.Name = images[m].UniqueId
		imageData.Type = 1
		imageData.Data = images[m].Content
		imageArr = append(imageArr, imageData)
	}
	result, _ := json.Marshal(imageArr)
	params := url.Values{
		"images":      []string{string(result)},
		"version":     []string{base.Conf.Dun.ImageVersion},
		"businessId":  []string{base.Conf.Dun.MomentImageBusinessId},
		"checkLabels": checkLabels,
	}

	apiUrl := "http://as.dun.163yun.com/v4/image/check"
	jsonStr, err := common.BaseCheck(params, apiUrl)
	if err != nil {
		return DunImageCheckResponse{}, err
	}
	var re DunImageCheckResponse
	err = json.Unmarshal([]byte(jsonStr), &re)
	if err != nil {
		return DunImageCheckResponse{}, err
	}
	return re, nil
}
