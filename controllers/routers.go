package controllers

import "github.com/gin-gonic/gin"

func DefineRouters(router *gin.Engine) {
	bookDefineRouters(router)
}

func bookDefineRouters(router *gin.Engine) {
	group := router.Group("/book")
	{
		group.POST("", CreateBook)
		group.GET("", FindAllBook)
	}
}
