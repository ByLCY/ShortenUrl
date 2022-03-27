package router

import (
	"ByLCY/ShortenUrl/controller"

	"github.com/gin-gonic/gin"
)

func (r routes) ShortenUrl(rg *gin.RouterGroup) {

	shortenUrl := controller.NewShortenUrlControlle()

	shortenUrlRg := rg.Group("/ShortenUrl")
	{
		shortenUrlRg.POST("/", shortenUrl.GenShortUrl)
	}
}
