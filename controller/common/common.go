package common

import (
	"net/http"

	"github.com/konger/ckgo/common/codes"
	"github.com/konger/ckgo/common/setting"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseModelBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ret := ResponseModel{Code: codes.SUCCESS, Message: "ok", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应成功
func ResSuccessMsg(c *gin.Context) {
	ret := ResponseModelBase{Code: codes.SUCCESS, Message: "ok"}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应失败
func ResFail(c *gin.Context, msg string) {
	ret := ResponseModelBase{Code: codes.ERROR, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应失败
func ResFailCode(c *gin.Context, msg string, code int) {
	ret := ResponseModelBase{Code: code, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	c.JSON(status, v)
	c.Abort()
}

// 响应错误-服务端故障
func ResErrSrv(c *gin.Context, err error) {
	ret := ResponseModelBase{Code: codes.ERROR, Message: "服务端故障"}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应错误-用户端故障
func ResErrCli(c *gin.Context, err error) {
	ret := ResponseModelBase{Code: codes.ERROR, Message: "err"}
	ResJSON(c, http.StatusOK, &ret)
}

type ResponsePageData struct {
	Total uint64      `json:"total"`
	Items interface{} `json:"items"`
}

type ResponsePage struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    ResponsePageData `json:"data"`
}

// 响应成功-分页数据
func ResSuccessPage(c *gin.Context, total uint64, list interface{}) {
	ret := ResponsePage{Code: codes.SUCCESS, Message: "ok", Data: ResponsePageData{Total: total, Items: list}}
	ResJSON(c, http.StatusOK, &ret)
}

// 获取页码
func GetPageIndex(c *gin.Context) uint64 {
	return GetQueryToUint64(c, "page", 1)
}

// 获取每页记录数
func GetPageLimit(c *gin.Context) uint64 {
	limit := GetQueryToUint64(c, "limit", 20)
	if limit > 500 {
		limit = 20
	}
	return limit
}

//GetPage 获取每页数量
func GetPage(c *gin.Context) (page, pagesize int) {
	page, _ = strconv.Atoi(c.Query("page"))
	pagesize, _ = strconv.Atoi(c.Query("limit"))
	if pagesize == 0 {
		pagesize = setting.App["Pagesize"].(int)
	}
	if page == 0 {
		page = 1
	}
	return
}

// 获取排序信息
func GetPageSort(c *gin.Context) string {
	return GetQueryToStr(c, "sort")
}

// 获取搜索关键词信息
func GetPageKey(c *gin.Context) string {
	return GetQueryToStr(c, "key")
}
