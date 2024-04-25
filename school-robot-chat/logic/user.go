package logic

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"school-robot-chat/dao/mysql"
	"school-robot-chat/model"
	"school-robot-chat/pkg/snowflake"
)

const secret = "shangyizhou"

func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func IsUserExist(user *model.User) (isExist bool, err error) {
	_, err = mysql.GetUserByName(user.UserName)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		isExist = false
	} else {
		isExist = true
	}
	return isExist, err
}

func Register(user *model.User) (err error) {
	isExist, _ := IsUserExist(user)
	if isExist == true {
		zap.L().Info("user is exist")
		return ErrorUserExit
	}

	// 生成user_id
	userID, err := snowflake.GetID()
	if err != nil {
		return ErrorGenIDFailed
	}

	// 生成加密密码
	password := encryptPassword([]byte(user.Password))
	user.Password = password
	user.ID = userID
	err = mysql.InsertUser(user)
	if err != nil {
		return nil
	}
	return err
}

func Login(user *model.User) (err error) {
	u, err := mysql.GetUserByName(user.UserName)
	if err != nil && err == gorm.ErrRecordNotFound {
		zap.L().Info("no exist user")
		return ErrorUserNotExit
	}
	// 生成加密密码与查询到的密码比较
	password := encryptPassword([]byte(user.Password))
	if u.Password != password {
		zap.L().Info("passwd error", zap.Any("passwd", password))
		return ErrorPasswordWrong
	}

	user.ID = u.ID
	return nil
}
