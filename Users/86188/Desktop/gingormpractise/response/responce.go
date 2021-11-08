package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Respnse(ctx *gin.Context,httpStatus int,code int,data gin.H,msg string)  {
	ctx.JSON(httpStatus,gin.H{
		"code": code,
		"data": data,
		"msg": msg,
	})
}
func Success(ctx *gin.Context,data gin.H,msg string)  {
	Respnse(ctx,http.StatusOK,200,data,msg)
}
func Fail(ctx *gin.Context,data gin.H,msg string)  {
	Respnse(ctx,http.StatusOK,400,data,msg)
}
