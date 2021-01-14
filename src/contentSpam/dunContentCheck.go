package contentSpam

import (
	"antispam/base"
	"antispam/common"
	"antispam/models"
	"encoding/json"
	"net/url"
	"sync"
)

type contentLabelDetails struct {
	Hint     []string `json:"hint"`
	HintInfo []string `json:"hintInfo"`
}

type contentCheckLabels struct {
	Label     int64               `json:"label"`
	Level     int64               `json:"level"`
	Details   contentLabelDetails `json:"details"`
	SubLabels []string            `json:"subLabels"`
}

type contentCheckResult struct {
	TaskId       string               `json:"taskId"`
	Action       int64                `json:"action"`
	CensorType   int64                `json:"censorType"`
	IsRelatedHit bool                 `json:"isRelatedHit"`
	Lang         []string             `json:"lang"`
	Label        []contentCheckLabels `json:"labels"`
}

type DunContentCheckResponse struct {
	Code   int64              `json:"code"`
	Msg    string             `json:"msg"`
	Result contentCheckResult `json:"result"`
}

func DunContentCheck(data []models.Data, checkLabels string) map[string]models.ContentResult {
	var responses = map[string]models.ContentResult{}
	wg := sync.WaitGroup{}
	for m := range data {
		wg.Add(1)
		go func(wg *sync.WaitGroup, m int) {
			defer wg.Done()
			var result models.ContentResult
			params := url.Values{
				"content":     []string{data[m].Content},
				"dataId":      []string{data[m].UniqueId},
				"version":     []string{base.Conf.Dun.ContentVersion},
				"businessId":  []string{base.Conf.Dun.BusinessId},
				"checkLabels": []string{checkLabels},
			}
			apiUrl := "https://as.dun.163yun.com/v3/text/check"

			jsonStr, err := common.BaseCheck(params, apiUrl)
			if err != nil {
				result.Status = models.PictureActionMapping[2]
			} else {
				re := DunContentCheckResponse{}
				err = json.Unmarshal([]byte(jsonStr), &re)
				if err != nil {
					result.Status = models.PictureActionMapping[2]
				} else {
					result.Status = models.PictureActionMapping[re.Result.Action]
				}
			}
			responses[data[m].UniqueId] = result
		}(&wg, m)
	}
	wg.Wait()
	return responses
}
