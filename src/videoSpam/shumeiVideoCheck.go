package videoSpam

import (
	"antispam/base"
	"antispam/common"
	"antispam/models"
	"encoding/json"
	"strings"
	"sync"
)

type VideoCheckResponse struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
	BtId      string `json:"btId"`
}

type VideoResultResponse struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
	BtId      string `json:"btId"`
	Labels    string `json:"labels"`
	RiskLevel string `json:"riskLevel"`
}

func getVideoImgType(model []string) string {
	if len(model) == 0 {
		return "DEFAULT"
	}
	return strings.Join(model, "_")
}

func ShuMeiVideoContentCheck(data []models.Data, model []string) (map[string]models.ContentResult, error) {
	wg := sync.WaitGroup{}
	var finalResult = map[string]models.ContentResult{}
	for m := range data {
		wg.Add(1)
		go func(wg *sync.WaitGroup, m int) {
			defer wg.Done()
			var r models.ContentResult
			imgType := getVideoImgType(model)
			apiUrl := "http://api-video-bj.fengkongcloud.com/v2/saas/anti_fraud/video"
			payload := map[string]interface{}{
				"imgType":   imgType,
				"audioType": "NONE",
				"appId":     "default",
				"btId":      data[m].UniqueId,
				"data": map[string]interface{}{
					"url": data[m].Content,
				},
			}
			if data[m].Callback != "" {
				payload["callback"] = data[m].Callback
			}
			result, err := common.VideoCheck(payload, apiUrl)
			if err != nil {
				r.Status = "FAIL"
			} else {
				var re VideoCheckResponse
				err = json.Unmarshal([]byte(result), &re)
				base.Info("unique_id:" + data[m].UniqueId + ",content:" + data[m].Content + ",response:" + result + ",callback:" + data[m].Callback)
				if err != nil {
					r.Status = "FAIL"
				} else {
					if re.Code == 1100 {
						r.Status = "SUCCESS"
					} else {
						r.Status = "FAIL"
						r.Msg = re.Message
					}
				}
			}
			finalResult[data[m].UniqueId] = r
		}(&wg, m)
	}
	wg.Wait()
	return finalResult, nil
}

func ShuMeiVideoContentResult(data []models.Data) (map[string]models.ContentResult, error) {
	var finalResult = map[string]models.ContentResult{}
	apiUrl := "http://video-api.fengkongcloud.com/v2/saas/anti_fraud/query_video"
	wg := sync.WaitGroup{}
	for m := range data {
		wg.Add(1)
		go func(wg *sync.WaitGroup, m int) {
			defer wg.Done()
			var r models.ContentResult
			payload := map[string]interface{}{
				"btId": data[m].UniqueId,
			}
			result, err := common.GetVideoResult(payload, apiUrl)
			if err != nil {
				r.Status = "FAIL"
			} else {
				var re VideoResultResponse
				err = json.Unmarshal([]byte(result), &re)
				if err != nil {
					r.Status = "FAIL"
				} else {
					r.Status = re.RiskLevel
				}
			}
			finalResult[data[m].UniqueId] = r
		}(&wg, m)
	}
	wg.Wait()
	return finalResult, nil
}
