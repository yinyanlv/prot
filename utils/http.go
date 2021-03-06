package utils

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func IsAjax(c *gin.Context) bool {
	key := "x-requested-with"
	val := "XMLHttpRequest"
	if c.Request.Header.Get(key) == val {
		return true
	} else {
		return false
	}
}

func GetPage(c *gin.Context) (uint64, uint64) {
	pageIndex, _ := strconv.ParseUint(c.Query("page"), 10, 64)
	pageSize, _ := strconv.ParseUint(c.Query("size"), 10, 64)

	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}

	return pageIndex, pageSize
}

type Pagination struct {
	IsShowPrev bool
	IsShowNext bool
	PrevPage   uint64
	NextPage   uint64
	PageSize   uint64
}

func GetPagination(count uint64, pageIndex uint64, pageSize uint64) Pagination {
	isShowPrev := true
	isShowNext := true

	if pageIndex == 1 {
		isShowPrev = false
	}

	if pageIndex*pageSize >= count {
		isShowNext = false
	}

	return Pagination{
		IsShowPrev: isShowPrev,
		IsShowNext: isShowNext,
		PrevPage:   pageIndex - 1,
		NextPage:   pageIndex + 1,
		PageSize:   pageSize,
	}
}

func HandleError(c *gin.Context, t string, code int, err error) {
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}
	log.WithFields(log.Fields{
		"url":        c.Request.URL.String(),
		"user_agent": c.Request.UserAgent(),
		"ip":     c.ClientIP(),
	}).Error(errStr)

	if t == "html" {
		c.HTML(code, "error", gin.H{
			"errorCode": code,
			"message":   err,
		})
	} else {
		c.JSON(code, gin.H{
			"success": false,
			"message": err,
		})
	}
}
