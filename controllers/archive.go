package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/models"
	"prot/utils"
)

func RenderArchive(c *gin.Context) {
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

	session := sessions.Default(c)
	userInfo := session.Get("userInfo")

	c.HTML(http.StatusOK, "list", gin.H{
		"pageCode":   "archive",
		"userInfo": userInfo,
		"pageTitle":  "归档",
		"articles":   articles,
		"pagination": pagination,
		"baseUrl":    "/archive",
	})
}

func RenderArchiveByYearMonth(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")

	pageIndex, pageSize := utils.GetPage(c)

	a := models.Article{}
	var articles []models.Article
	var err error

	articles, err = a.GetArticlesByYearMonth(year, month, pageIndex, pageSize)
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

	session := sessions.Default(c)
	userInfo := session.Get("userInfo")

	c.HTML(http.StatusOK, "list", gin.H{
		"pageCode":   "archive",
		"userInfo": userInfo,
		"pageTitle":  "归档 " + year + "-" + month,
		"articles":   articles,
		"pagination": pagination,
		"baseUrl":    "/archive/" + year + "/" + month,
	})
}
