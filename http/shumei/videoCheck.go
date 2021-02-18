package shumei

import (
	"antispam/base"
	"antispam/models"
	"antispam/src"
	"antispam/src/videoSpam"
	"github.com/gin-gonic/gin"
)

func ApiShuMeiPostVideoCheck(c *gin.Context) {
	requestBody := models.Check{}
	c.BindJSON(&requestBody)
	perData := src.VideoCheckFirstStep(requestBody.Data)
	result, err := src.ShuMeiVideoCheckSecondStep(perData, requestBody.Model)
	if err != nil {
		base.Forbidden(c, "调用三方sdk失败", err)
	}
	finalResult := src.VideoThirdStep(result)
	base.PostOk(c, "check finish", finalResult)
	return
}

func ApiShuMeiPostVideoResult(c *gin.Context) {
		requestBody := models.Check{}
		c.BindJSON(&requestBody)
		result, err := videoSpam.ShuMeiVideoContentResult(requestBody.Data)
		if err != nil {
			base.Forbidden(c, "调用三方sdk失败", err)
		}
		base.PostOk(c, "get result ok", result)
		return
}
