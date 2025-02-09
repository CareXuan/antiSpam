package http

import (
	netease2 "antispam/http/dun"
	ShuMei2 "antispam/http/shumei"
	"github.com/gin-gonic/gin"
)

func RouteInit(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		dun := v1.Group("dun")
		{
			dun.POST("/content", netease2.ApiDunPostContentCheck)
			dun.POST("/picture", netease2.ApiDunPostPictureCheck)
		}
		shumei := v1.Group("shumei")
		{
			shumei.POST("/video", ShuMei2.ApiShuMeiPostVideoCheck)
			shumei.POST("/video/result", ShuMei2.ApiShuMeiPostVideoResult)
		}
		//callback := v1.Group("callback")
		//{
		//	callback.POST("/video")
		//}
	}
}
