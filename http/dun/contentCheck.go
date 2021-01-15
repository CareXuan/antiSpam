package dun

import (
	"antispam/base"
	"antispam/models"
	"antispam/src"
	"github.com/gin-gonic/gin"
	"strings"
)

func ApiDunPostContentCheck(c *gin.Context) {
	requestBody := models.Check{}
	c.BindJSON(&requestBody)
	base.Info("taskId:" + requestBody.TaskId + " start content check,models:" + strings.Join(requestBody.Model, ","))
	preData := src.ContentCheckFirstStep(requestBody.Data)
	result := src.DunContentCheckSecondStep(preData, requestBody.Model)
	finalResult := src.ContentCheckThirdStep(result)
	base.PostOk(c, "check finish", finalResult)
	return
}
