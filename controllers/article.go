package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"html/template"
	"net/http"
	"prot/models"
	"prot/utils"
)

func RenderArticle(c *gin.Context) {
	session := sessions.Default(c)
	userInfo := session.Get("userInfo")
	id := c.Param("id")
	article := models.Article{}
	article.AddViewCount(id)
	a, err := article.Article(id)
	if err != nil {
		utils.HandleError(c, "html", 599, err)
		return
	}
	md := []byte(a.Content)
	content := string(markdown.ToHTML(md, nil, nil))

	c.HTML(http.StatusOK, "article", gin.H{
		"pageCode":  "article",
		"userInfo":  userInfo,
		"id":        a.ID,
		"title":     a.Title,
		"tags":      a.Tags,
		"content":   template.HTML(content),
		"createdAt": a.CreatedAt,
		"updatedAt": a.UpdatedAt,
	})
}

func RenderEditArticle(c *gin.Context) {
	session := sessions.Default(c)
	tag := models.Tag{}
	tags, err := tag.GetTags()
	if err != nil {
		utils.HandleError(c, "html", 599, err)
		return
	}

	id := c.Query("id")
	userInfo := session.Get("userInfo")

	if id == "" {
		c.HTML(http.StatusOK, "edit", gin.H{
			"pageCode": "edit-article",
			"userInfo": userInfo,
			"tags":     tags,
		})
	} else {
		article := models.Article{}
		article.AddViewCount(id)
		a, err := article.Article(id)
		if err != nil {
			utils.HandleError(c, "html", 599, err)
			return
		}

		c.HTML(http.StatusOK, "edit", gin.H{
			"pageCode": "edit-article",
			"userInfo": userInfo,
			"tags":     tags,
			"title":    a.Title,
			"summary":  a.Summary,
			"curTags":  a.Tags,
			"content":  a.Content,
		})
	}
}

func CreateArticle(c *gin.Context) {
	session := sessions.Default(c)
	userInfo := session.Get("userInfo").(models.User)
	userId := userInfo.ID

	var articleReq models.ArticleReq
	var tags []string

	c.BindJSON(&articleReq)
	tags = articleReq.Tags

	tag := &models.Tag{}

	newTags, _ := tag.GetTagsByIDS(tags)

	article := models.Article{
		Title:   articleReq.Title,
		Summary: articleReq.Summary,
		Tags:    newTags,
		Content: articleReq.Content,
		CommonStrID: models.CommonStrID{
			CreatedBy: userId,
			UpdatedBy: userId,
		},
	}
	id, err := article.Insert()

	if err != nil {
		utils.HandleError(c, "json", 599, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    id,
	})
}

func UpdateArticle(c *gin.Context) {
	session := sessions.Default(c)
	userInfo := session.Get("userInfo").(models.User)
	userId := userInfo.ID

	var articleReq models.ArticleReq
	var tags []string

	c.BindJSON(&articleReq)
	tags = articleReq.Tags

	tag := &models.Tag{}

	newTags, _ := tag.GetTagsByIDS(tags)

	article := models.Article{
		Title:   articleReq.Title,
		Summary: articleReq.Summary,
		Tags:    newTags,
		Content: articleReq.Content,
		CommonStrID: models.CommonStrID{
			ID:      articleReq.ID,
			UpdatedBy: userId,
		},
	}
	id, err := article.Update()

	if err != nil {
		utils.HandleError(c, "json", 599, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    id,
	})

}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	a := models.Article{}
	_, err := a.Delete(id)

	if err != nil {
		utils.HandleError(c, "json", 599, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    id,
	})
}
