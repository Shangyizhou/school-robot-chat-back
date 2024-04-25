package mysql

import (
	"errors"
	"github.com/jinzhu/gorm"
	"school-robot-chat/model"
)

func GetUserByID(id string) (user model.User, err error) {
	err = db.Where("user_id = ?", id).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return model.User{}, errors.New("user not found")
	}
	return user, err
}

func GetUserByName(name string) (user model.User, err error) {
	err = db.Where("user_name =?", name).First(&user).Error
	//if err == gorm.ErrRecordNotFound {
	//	return model.User{}, errors.New("user not found")
	//}
	return user, err
}

// 插入一条新的User记录，外部传入指针
func InsertUser(user *model.User) (err error) {
	err = db.Create(user).Error
	if err != nil {
		panic(errors.New("failed to insert data"))
	}
	return nil
}

// Save()默认会更新所有字段，外部传入指针
func UpdateUser(user *model.User) (err error) {
	err = db.Save(user).Error
	if err != nil {
		panic(errors.New("failed to save data"))
	}
	return nil
}

// 设置User为登录状态，外部传入指针
func SetUserOnline(user *model.User) (err error) {
	db.Model(user).Update("state", "online")
	if err != nil {
		panic(errors.New("failed to set user online"))
	}
	return nil
}

// 查询所有登录状态为 online 的 user 列表
func GetOnlineUserList() (users []model.User, err error) {
	err = db.Where("status =?", "online").Find(&users).Error
	if err != nil {
		panic(errors.New("failed to get users online"))
	}
	return users, nil
}
