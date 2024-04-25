package model

import (
	"encoding/json"
	"errors"
)

//`omitempty` 是 Go 语言中在结构体标签中常用的一个属性。
//
//它的主要作用是：当该字段的值为空（例如空字符串、零值等）时，在将结构体序列化为 JSON 时，不会将该字段包含在 JSON 结果中。
//
//这样做的好处是：
//1. 减少 JSON 数据的大小，提高传输效率。
//2. 使得 JSON 数据更加简洁和易读。
//
//在使用 `omitempty` 时需要注意：
//1. 如果字段的值不为空，那么它会正常出现在序列化后的 JSON 中。
//2. 对于一些默认值为零值的字段，如果希望它们在 JSON 中出现，即使值为默认值，也不要使用 `omitempty`。

type User struct {
	ID            uint64 `gorm:"not null" gorm:"primary_key" json:"user_id,omitempty" db:"user_id"`
	state         string `gorm:"not null" json:"state,omitempty" default:"offline" db:"state"`
	UserName      string `gorm:"user_name" json:"user_name,omitempty" default:"" db:"user_name"`
	Password      string `gorm:"not null" json:"password,omitempty" default:"" db:"password"`
	Photo         string `gorm:"not null" json:"photo,omitempty" default:"" db:"photo"`
	SchoolName    string `gorm:"not null" json:"school_name,omitempty" default:"" db:"school_name"`
	ClassName     string `gorm:"not null" json:"class_name,omitempty" default:"" db:"class_name"`
	Sex           bool   `gorm:"not null" json:"sex,omitempty" default:"true" db:"sex"`
	Desc          string `gorm:"not null" json:"desc,omitempty" default:"" db:"desc"`
	Age           int    `gorm:"not null" json:"age,omitempty" default:"0" db:"age"`
	Birthday      string `gorm:"not null" json:"birthday,omitempty" default:"" db:"birthday"`
	Constellation string `gorm:"not null" json:"constellation,omitempty" default:"" db:"constellation"`
	Hobby         string `gorm:"not null" json:"hobby,omitempty" default:"" db:"hobby"`
}

// 将 User 的表名设置为 `user`
func (User) TableName() string {
	return "user"
}

// 登录相关
// 假设有一个 JSON 字符串 jsonData，你可以这样调用
// var user User
// err := json.Unmarshal([]byte(jsonData), &user)
// if err!= nil {
// 处理解析错误
// }
func (user *User) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return err
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段 user_name")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段 password")
	} else {
		user.UserName = required.UserName
		user.Password = required.Password
	}
	return nil
}

// 注册相关
// 好像postman只可以以raw形式接受
type RegisterForm struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required"`
}

func (r *RegisterForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName        string `json:"user_name"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return err
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else if required.Password != required.ConfirmPassword {
		err = errors.New("两次密码不一致")
	} else {
		r.UserName = required.UserName
		r.Password = required.Password
		r.ConfirmPassword = required.ConfirmPassword
	}
	return nil
}
