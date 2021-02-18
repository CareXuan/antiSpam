package dun

import (
	"antispam/base"
	"antispam/models"
	"antispam/src"
	"github.com/gin-gonic/gin"
)

func ApiDunPostContentCheck(c *gin.Context) {
	requestBody := models.Check{}
	c.BindJSON(&requestBody)
	preData := src.ContentCheckFirstStep(requestBody.Data)
	result := src.DunContentCheckSecondStep(preData, requestBody.Model)
	finalResult := src.ContentCheckThirdStep(result)
	base.PostOk(c, "check finish", finalResult)
	return
}
