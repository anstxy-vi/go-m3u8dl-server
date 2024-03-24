package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()

	v1_routes(router)
	return router
}

func v1_routes(router *gin.Engine) {
	v1_router := router.Group("/v1")

	groupRoutes(v1_router) // 注册用户auth相关接口（登录、注册...）
}
