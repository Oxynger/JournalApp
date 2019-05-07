package service

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"

	"github.com/Oxynger/JournalApp/model/user"
)

type SessionService struct {
	sessions map[string]*Session
	lock     sync.RWMutex
}

type Session struct {
	Token    string
	Role     user.Role
	Username string
	ExpireAt int64
}

func NewSessionService() *SessionService {
	return &SessionService{
		sessions: make(map[string]*Session),
	}
}

func (srv *SessionService) VerifyToken(token string) bool {
	srv.lock.RLock()
	defer srv.lock.RUnlock()

	session, ok := srv.sessions[token]
	if !ok || session.ExpireAt < time.Now().Unix() {
		return false
	}
	return true
}

func (srv *SessionService) InvalidateToken(token string) {
	srv.lock.Lock()
	defer srv.lock.Unlock()

	delete(srv.sessions, token)
}

func (srv *SessionService) CreateSession(usr *user.User) (*Session, error) {
	token, err := generateTokenString()
	if err != nil {
		return nil, err
	}

	s := &Session{
		Token:    token,
		Role:     usr.Role,
		Username: usr.Username,
		ExpireAt: time.Now().Add(time.Hour).Unix(),
	}
	srv.sessions[token] = s
	return s, nil
}

func generateTokenString() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), err
}
