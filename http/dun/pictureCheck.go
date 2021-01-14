package dun

import (
	"antispam/base"
	"antispam/models"
	"antispam/src"
	"github.com/gin-gonic/gin"
)

func ApiDunPostPictureCheck(c *gin.Context) {
	requestBody := models.Check{}
	c.BindJSON(&requestBody)
	perUrls := src.PictureCheckFirstStep(requestBody.Data)
	result, err := src.DunPictureCheckSecondStep(perUrls, requestBody.Model)
	if err != nil {
		base.Forbidden(c, "调用三方sdk失败", err)
	}
	finalResult := src.PictureCheckThirdStep(result)
	base.PostOk(c, "check finish", finalResult)
	return
}
