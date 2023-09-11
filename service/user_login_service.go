package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
	"singo/util" // 导入您之前封装好的JWT工具函数
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// GenerateJWT 生成JWT令牌
func (service *UserLoginService) GenerateJWT(user model.User) (string, error) {
	// 使用用户名和权限生成JWT令牌
	permissions := user.Role // 假设User模型中有一个名为Role的字段表示权限
	return util.GenerateJWT(service.UserName, permissions)
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	// 生成JWT令牌
	token, err := service.GenerateJWT(user)
	if err != nil {
		return serializer.Err(serializer.ErrCodeInternalError, "生成JWT令牌失败", err)
	}

	// 设置session
	service.setSession(c, user)

	// 将JWT令牌返回给客户端
	return serializer.Response{
		Code: serializer.CodeSuccess,
		Data: token,
	}
}
