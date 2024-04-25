package mysql

import (
	"school-robot-chat/model"
)

func FindSessionsForUser(userID int64) []model.Session {
	var sessions []model.Session
	db.Where("user_id = ? AND is_del != ?", userID, true).Find(&sessions)
	return sessions
}
