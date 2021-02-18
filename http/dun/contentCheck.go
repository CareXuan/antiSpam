package dun

import (
	"antispam/base"
	"antispam/models"
	"antispam/src"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func ApiDunPostContentCheck(c *gin.Context) {
	requestBody := models.Check{}
	c.BindJSON(&requestBody)
	requestJson, _ := json.Marshal(requestBody)
	base.Conf.ContentRequestChan <- string(requestJson)
	preData := src.ContentCheckFirstStep(requestBody.Data)
	result := src.DunContentCheckSecondStep(preData, requestBody.Model)
	finalResult := src.ContentCheckThirdStep(result)
	responseJson, _ := json.Marshal(finalResult)
	base.Conf.ContentResponseChan <- string(responseJson)
	base.PostOk(c, "check finish", finalResult)
	return
}
