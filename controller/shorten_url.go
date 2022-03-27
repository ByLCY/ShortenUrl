package controller

import (
	"ByLCY/ShortenUrl/model"
	"ByLCY/ShortenUrl/util"
	"ByLCY/ShortenUrl/util/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type shortenUrlController struct {
}

func NewShortenUrlControlle() shortenUrlController {
	return shortenUrlController{}
}

type LongURL struct {
	Url string `json:"url" binding:"required,url"`
}

func (ctr shortenUrlController) GenShortUrl(c *gin.Context) {
	var url LongURL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400001,
			"message": "Parameter error",
		})
		return
	}

	urlModel := model.Url{
		LongURL: url.Url,
	}

	pdb := db.NewDb()
	pdb.Select("LongURL").Create(&urlModel)

	shortPath := "/" + util.GenShortPath(urlModel.Id, 6)
	var count int64
	pdb.Model(&model.Url{}).Where("ShortURL = ?", shortPath).Count(&count)
}
