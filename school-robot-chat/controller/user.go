package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"school-robot-chat/logic"
	"school-robot-chat/model"
	"school-robot-chat/pkg/jwt"
	"strings"
)

func SignUpHandler(c *gin.Context) {
	// 1.获取请求参数
	// 2.校验数据有效性
	var form model.RegisterForm
	if err := c.ShouldBindJSON(&form); err != nil {
		zap.L().Error("ShouldBindJSON mysql.Register() failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	zap.L().Info("RegisterForm detail", zap.Any("form", form))
	// 3.注册用户
	err := logic.Register(&model.User{
		UserName: form.UserName,
		Password: form.Password,
	})
	if errors.Is(err, logic.ErrorUserExit) {
		zap.L().Error("mysql.Register() failed, user exist", zap.Error(err))
		ResponseError(c, CodeUserExist)
		return
	}
	if err != nil {
		zap.L().Error("mysql.Register() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// ShouldBindJSON直接将JSON字符串映射到User
func LoginHandler(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}
	zap.L().Info("LoginHandler ShouldBindJSON", zap.Any("user", u))
	if err := logic.Login(&u); err != nil {
		zap.L().Error("mysql.Login(&u) failed", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 生成Token
	aToken, rToken, _ := jwt.GenToken(u.ID)
	ResponseSuccess(c, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
		"user_id":       u.ID,
		"user_name":     u.UserName,
	})
}

func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, CodeInvalidToken, "Token格式不对")
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
