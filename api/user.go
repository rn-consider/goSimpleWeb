package api

import (
	"singo/serializer"
	"singo/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
// @Summary 用户注册
// @Description 用户注册接口用于创建新用户账号
// @Tags 用户
// @Accept json
// @Produce json
// @Param user_name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} serializer.Response
// @Failure 400 {object} serializer.Response
// @Router /api/v1/user/register [post]
func UserRegister(c *gin.Context) {
	var registerService service.UserRegisterService
	if err := c.ShouldBind(&registerService); err == nil {
		res := registerService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
// @Summary 用户登录
// @Description 用户登录接口用于验证用户身份并返回JWT令牌
// @Tags 用户
// @Accept json
// @Produce json
// @Param user_name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} serializer.Response
// @Failure 400 {object} serializer.Response
// @Router /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	var loginService service.UserLoginService
	if err := c.ShouldBind(&loginService); err == nil {
		res := loginService.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
// @Summary 用户详情
// @Description 用户详情接口用于获取当前用户的详细信息
// @Tags 用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} serializer.Response
// @Failure 400 {object} serializer.Response
// @Router /api/v1/user/me [get]
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
// @Summary 用户登出
// @Description 用户登出接口用于清除用户会话并登出
// @Tags 用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} serializer.Response
// @Failure 400 {object} serializer.Response
// @Router /api/v1/user/logout [post]
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
