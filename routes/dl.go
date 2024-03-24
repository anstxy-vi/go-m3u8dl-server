package routes

import (
	v1 "gorman/m3u8dl/api/v1"

	"github.com/gin-gonic/gin"
)

func groupRoutes(router *gin.RouterGroup) {

	api := v1.NewDLApi()

	//	$TEST
	router.POST("/dl", api.HandleUrls)
}
