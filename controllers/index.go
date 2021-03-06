package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/models"
	"prot/utils"
)

func RenderIndex(c *gin.Context) {
	session := sessions.Default(c)
	pageIndex, pageSize := utils.GetPage(c)
	a := models.Article{}
	var articles []models.Article
	var err error

	articles, err = a.GetArticlesByTag("", pageIndex, pageSize)
	if err != nil {
		utils.HandleError(c, "html", 599, err)
		return
	}

	count, err := a.GetCountByTag("")
	if err != nil {
		utils.HandleError(c, "html", 599, err)
		return
	}
	pagination := utils.GetPagination(count, pageIndex, pageSize)

	userInfo := session.Get("userInfo")

	c.HTML(http.StatusOK, "index", gin.H{
		"pageCode": "index",
		"userInfo": userInfo,
		"articles": articles,
		"pagination": pagination,
	})
}

