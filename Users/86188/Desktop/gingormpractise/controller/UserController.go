package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"golangstudy/jike/gingormpractise/common"
	"golangstudy/jike/gingormpractise/dto"
	"golangstudy/jike/gingormpractise/model"
	"golangstudy/jike/gingormpractise/response"
	"golangstudy/jike/gingormpractise/util"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB:=common.GetDB()
	//获取参数
	name:=ctx.PostForm("name")
	telephone:=ctx.PostForm("telephone")
	password:=ctx.PostForm("password")
	//数据验证
	if len(telephone)!=11 {
		response.Respnse(ctx,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")

		return
	}
	if len(password)<6 {
		response.Respnse(ctx,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		return
	}
	//如果名称没有传 那么就给一个随机的十位字符串
	if len(name)==0 {
		name=util.RandomString(10)
	}
	//判断手机号是否存在
	if isTelephoneExist(DB,telephone){
		response.Respnse(ctx,http.StatusUnprocessableEntity,422,nil,"手机号已被注册")
		return
	}
	//创建用户
	hasedPassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)  //密码加密
	if err!=nil {
		response.Respnse(ctx,http.StatusInternalServerError,500,nil,"加密失败")
		return
	}
	newUser:=model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hasedPassword),
	}
	DB.Create(&newUser)
	//返回结果
	log.Println(name,password,telephone)
	response.Success(ctx,nil,"注册成功")
}
func Login(ctx *gin.Context)  {
	//获取参数
	DB:=common.GetDB()
	telephone:=ctx.PostForm("telephone")
	password:=ctx.PostForm("password")
	//数据验证
	if len(telephone)!=11 {
		response.Respnse(ctx,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		return
	}
	if len(password)<6 {
		response.Respnse(ctx,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone =?",telephone).First(&user)
	if user.ID==0 {
		response.Respnse(ctx,http.StatusUnprocessableEntity,422,nil,"用户不存在")
		return
	}
	//判断密码是否正确
    if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil {
		response.Respnse(ctx,http.StatusBadRequest,400,nil,"密码错误")
	} //第一个参数是原始密码,第二个参数是加密后的密码
	//发放token
	token,err:=common.ReleaseToken(user)
	if err!=nil{
		fmt.Println("failed to release token,err:",err)
		response.Fail(ctx,nil,"token发放失败")
		return
	}
	//返回结果
	response.Success(ctx,gin.H{"token":token},"登陆成功")
}
func Info(ctx *gin.Context)  {
	user,_:=ctx.Get("user")
	response.Success(ctx,gin.H{"user": dto.ToUserDto(user.(model.User)),},"验证成功")
}



func isTelephoneExist(db *gorm.DB,telephone string) bool {
	var user model.User
	db.Where("telephone=?",telephone).First(&user)  //查到了第一个就把查询到的记录写到user里
	if user.ID!=0 {
		return true
	}
	return false
}

