package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/yin5th/gin-api/pkg/setting"
)

func InitRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    gin.SetMode(setting.ServerSetting.RunMode)

    r.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
	    "message": "hello world!",
	})
    })

    return r
}
