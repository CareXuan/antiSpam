package dun

import (
	"antispam/base"
	"antispam/models"
	"antispam/src"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func ApiDunPostContentCheck(c *gin.Context) {
	requestBody := models.Check{}
	c.BindJSON(&requestBody)
	base.Info("taskId:" + requestBody.TaskId + " start content check,models:" + strings.Join(requestBody.Model, ","))
	_, err := base.FindMongoOne("carexuan", "task", bson.M{"_id": requestBody.TaskId})
	if err == nil {
		base.Forbidden(c, "task id exists", []string{})
		return
	}
	_, err = base.AddMongoOne("carexuan", "task", bson.M{"_id": requestBody.TaskId, "type": "content", "models": strings.Join(requestBody.Model, ","), "data": requestBody.Data})
	if err != nil {
		base.AddMongoOne("carexuan", "error", bson.M{"taskId": requestBody.TaskId, "msg": err})
	}
	preData := src.ContentCheckFirstStep(requestBody.Data)
	result := src.DunContentCheckSecondStep(preData, requestBody.Model)
	finalResult := src.ContentCheckThirdStep(result)
	_, err = base.UpdateMongoOne("carexuan", "task", bson.M{"_id": requestBody.TaskId}, bson.M{"result": finalResult})
	if err != nil {
		fmt.Print(err)
	}
	base.PostOk(c, "check finish", finalResult)
	return
}
