package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"prot/models"
	"prot/utils"
)

func RenderListByTag(c *gin.Context) {
	tag := c.Param("id")
	pageIndex, pageSize := utils.GetPage(c)

	a := models.Article{}

	articles, err := a.GetArticlesByTag(tag, pageIndex, pageSize)
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

	t := models.Tag{}
	tagObj, err := t.GetTagByID(tag)

	if err != nil {
		utils.HandleError(c, "html", 599, err)
		return
	}

	session := sessions.Default(c)
	userInfo := session.Get("userInfo")

	c.HTML(http.StatusOK, "list", gin.H{
		"pageCode":  tag,
		"userInfo": userInfo,
		"pageTitle": tagObj.Name,
		"articles":  articles,
		"pagination": pagination,
	})
}

func GetTags(c *gin.Context) {
	tag := models.Tag{}
	tags, err := tag.GetTags()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    tags,
	})
}
