package controller

import (
	"sync"
	"time"
)

type Session struct {
	Name       string
	Desc       string
	RobotId    string
	UserId     string
	Url        string
	SessionId  string
	IsDel      bool
	CreateTime time.Time
	UpdateTime time.Time
}

type Message struct {
	// 根据实际情况填写
}

var (
	instance     *SessionManager
	instanceOnce sync.Once
)

type SessionManager struct {
	SessionsList   []*Session
	MessageMap     map[*Session][]*Message
	CurrentSession *Session
}

func GetInstance() *SessionManager {
	instanceOnce.Do(func() {
		instance = &SessionManager{
			SessionsList: make([]*Session, 0),
			MessageMap:   make(map[*Session][]*Message),
		}
		instance.init()
	})

	return instance
}

func (s *SessionManager) init() {
	s.loadSessionList()

	if len(s.SessionsList) == 0 {
		s.addFirstSession()
	}

	s.MessageMap = make(map[*Session][]*Message)
}

func (s *SessionManager) addFirstSession() {
	session := &Session{
		//先赋值默认值，你可以按照你的需要
		//...
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	s.SessionsList = append(s.SessionsList, session)
}

func (s *SessionManager) loadSessionList() {
	//从数据库加载会话到SessionsList
}

func (s *SessionManager) loadSessionMessage(session *Session) {
	//从数据库加载消息到MessageMap

}

func (s *SessionManager) SaveHistoryMessage() {
	//保存SessionsList和MessageMap到数据库
}

func (s *SessionManager) AddNewSession(session *Session) {
	if session != nil {
		s.SessionsList = append(s.SessionsList, session)
	}
}

func (s *SessionManager) getSessionList() []*Session {
	return s.SessionsList
}

func (s *SessionManager) getSessionMessages(session *Session) []*Message {
	return s.MessageMap[session]
}

func (s *SessionManager) getLastSession() *Session {
	//按照 SessionsList 中的session的CreateTime获取离现在最近的一个session
}
