package model

import (
	"school-robot-chat/pkg/snowflake"
	"time"
)

type Session struct {
	ID         uint64 `gorm:"primary_key;" json:"session_id,omitempty"`
	Name       string `gorm:"not_null;" json:"name,omitempty"`
	Desc       string `gorm:"not_null;" json:"desc,omitempty"`
	RobotID    string `gorm:"not_null;" json:"robot_id,omitempty"`
	UserID     string `gorm:"not_null;" json:"user_id,omitempty"`
	URL        string `gorm:"not_null;" json:"url,omitempty"`
	IsDel      bool   `gorm:"not_null;" json:"is_del,omitempty"`
	CreateTime int64  `gorm:"not_null;" json:"create_time,omitempty"`
	UpdateTime int64  `gorm:"not_null;" json:"update_time,omitempty"`
}

// 创建 Session 的方法
func (session Session) InertSession(name, desc, robotID, userID, url string) Session {
	// 生成 UUID 作为 ID
	ID, err := snowflake.GetID()
	if err != nil {
		panic(err)
	}

	return Session{
		ID:         ID,
		Name:       name,
		Desc:       desc,
		RobotID:    robotID,
		UserID:     userID,
		URL:        url,
		IsDel:      false,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
}
