package dun

import (
	"antispam/base"
	"antispam/models"
	"antispam/src"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func ApiDunPostPictureCheck(c *gin.Context) {
	requestBody := models.Check{}
	c.BindJSON(&requestBody)
	base.Info("taskId:" + requestBody.TaskId + " start content check,models:" + strings.Join(requestBody.Model, ","))
	_, err := base.FindMongoOne("carexuan", "task", bson.M{"_id": requestBody.TaskId})
	if err == nil {
		base.Forbidden(c, "task id exists", []string{})
		return
	}
	_, err = base.AddMongoOne("carexuan", "task", bson.M{"_id": requestBody.TaskId, "type": "picture", "models": strings.Join(requestBody.Model, ","), "data": requestBody.Data})
	if err != nil {
		base.AddMongoOne("carexuan", "error", bson.M{"taskId": requestBody.TaskId, "type": "picture", "msg": err})
	}
	perUrls := src.PictureCheckFirstStep(requestBody.Data)
	result, err := src.DunPictureCheckSecondStep(perUrls, requestBody.Model)
	if err != nil {
		base.AddMongoOne("carexuan", "error", bson.M{"taskId": requestBody.TaskId, "type": "picture", "msg": err})
		base.Forbidden(c, "调用三方sdk失败", err)
	}
	finalResult := src.PictureCheckThirdStep(result)
	_, err = base.UpdateMongoOne("carexuan", "task", bson.M{"_id": requestBody.TaskId}, bson.M{"result": finalResult})
	if err != nil {
		base.AddMongoOne("carexuan", "error", bson.M{"taskId": requestBody.TaskId, "type": "picture", "msg": err})
	}
	base.PostOk(c, "check finish", finalResult)
	return
}
