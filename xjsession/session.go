package xjsession

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

//未写完, 未正式使用

//func Default(ctx *xjhttp.Context) (*Session){
//
//}
//
//type Session interface {
//	NewSession() *Session
//	// Get returns the session value associated to the given key.
//	Get(key string) interface{}
//	// Set sets the session value associated to the given key.
//	Set(key string, val interface{})
//	// Delete removes the session value associated to the given key.
//	Delete(key string)
//	// Clear deletes all values in the session.
//	Clear()
//}

// SessionMgr session manager
type SessionMgr struct {
	sessionName string
	mLock       sync.RWMutex
	maxLifeTime int64
	sessions    map[string]*Session
}

//Session
type Session struct {
	sessionID string
	lastTime  time.Time
	values    map[interface{}]interface{}
}

// NewSessionMgr create session manager
func NewSessionMgr(sessionName string, maxLifeTime int64) *SessionMgr {
	mgr := &SessionMgr{sessionName: sessionName, maxLifeTime: maxLifeTime, sessions: make(map[string]*Session)}
	go mgr.SessionGC()
	return mgr
}

// NewSession create session
func (mgr *SessionMgr) NewSession(w http.ResponseWriter, r *http.Request) string {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()
	newSessionID := url.QueryEscape(mgr.NewSessionID())
	session := &Session{sessionID: newSessionID, lastTime: time.Now(),
		values: make(map[interface{}]interface{})}
	mgr.sessions[newSessionID] = session
	cookie := http.Cookie{Name: mgr.sessionName, Value: newSessionID,
		Path: "/", HttpOnly: true, MaxAge: int(mgr.maxLifeTime)}
	http.SetCookie(w, &cookie)
	return newSessionID
}

// EndSession
func (mgr *SessionMgr) EndSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(mgr.sessionName)
	if err != nil || cookie.Value == "" {
		return
	}
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()
	delete(mgr.sessions, cookie.Value)
	newCookie := http.Cookie{Name: mgr.sessionName,
		Path: "/", HttpOnly: true,
		Expires: time.Now(), MaxAge: -1}
	http.SetCookie(w, &newCookie)
}

// EndSessionByID end the session by session ID
func (mgr *SessionMgr) EndSessionByID(sessionID string) {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()
	delete(mgr.sessions, sessionID)
}

// SetSessionValue set value fo session
func (mgr *SessionMgr) SetSessionValue(sessionID string, key interface{}, value interface{}) error {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()
	if session, ok := mgr.sessions[sessionID]; ok {
		session.values[key] = value
		return nil
	}
	return errors.New("invalid session ID")
}

// GetSessionValue get value fo session
func (mgr *SessionMgr) GetSessionValue(sessionID string, key interface{}) (interface{}, error) {
	mgr.mLock.RLock()
	defer mgr.mLock.RUnlock()
	if session, ok := mgr.sessions[sessionID]; ok {
		if val, ok := session.values[key]; ok {
			return val, nil
		}
	}
	return nil, errors.New("invalid session ID")
}

//CheckCookieValid check cookie is valid or not
func (mgr *SessionMgr) CheckCookieValid(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie(mgr.sessionName)
	if cookie == nil ||
		err != nil {
		return "", err
	}
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()
	sessionID := cookie.Value
	if session, ok := mgr.sessions[sessionID]; ok {
		session.lastTime = time.Now()
		return sessionID, nil
	}
	return "", errors.New("invalid session ID")
}

// SessionGC maintain session
func (mgr *SessionMgr) SessionGC() {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()
	for id, session := range mgr.sessions {
		if session.lastTime.Unix()+mgr.maxLifeTime < time.Now().Unix() {
			delete(mgr.sessions, id)
		}
	}
	time.AfterFunc(time.Duration(mgr.maxLifeTime)*time.Second, func() {
		mgr.SessionGC()
	})
}

// NewSessionID generate unique ID
func (mgr *SessionMgr) NewSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		nano := time.Now().UnixNano()
		return strconv.FormatInt(nano, 10)
	}
	return base64.URLEncoding.EncodeToString(b)
}
