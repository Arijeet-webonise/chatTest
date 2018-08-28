package session

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	uuid "github.com/satori/go.uuid"
)

// SessionManager encapsulates session
type SessionManager interface {
	SetFlash(w http.ResponseWriter, r *http.Request, msg string, alertType string) error
	GetFlash(w http.ResponseWriter, r *http.Request) (*Flash, error)
	SetValue(w http.ResponseWriter, r *http.Request, name string, value interface{}) error
	GetValue(r *http.Request, name string) (interface{}, error)
	DeleteValue(w http.ResponseWriter, r *http.Request) error
	SetSession(w http.ResponseWriter, r *http.Request, userID int) (string, error)
	GetSession(r *http.Request) (*User, error)
	DeleteSession(w http.ResponseWriter, r *http.Request) error
}

// Session implements SessionManager
type Session struct {
	store   *sessions.CookieStore
	Project string
}

// Flash wrapper for flash
type Flash struct {
	Message string
	Type    string
}

// User wrapper for user session
type User struct {
	UUID   string
	UserID int
}

// CreateSessionManager create session manager
func CreateSessionManager(secret string) (SessionManager, error) {
	store := sessions.NewCookieStore([]byte(secret))
	return &Session{store: store}, nil
}

// SetFlash save flash
func (s *Session) SetFlash(w http.ResponseWriter, r *http.Request, msg string, alertType string) error {
	session, err := s.store.Get(r, s.Project+"_session_flash")
	if err != nil {
		return err
	}

	gob.Register(&Flash{})
	flash := &Flash{
		Message: msg,
		Type:    alertType,
	}

	session.AddFlash(flash, "flash")

	return session.Save(r, w)
}

// GetFlash save flash
func (s *Session) GetFlash(w http.ResponseWriter, r *http.Request) (*Flash, error) {
	session, err := s.store.Get(r, s.Project+"_session_flash")
	if err != nil {
		return nil, err
	}

	flashes := session.Flashes("flash")

	if len(flashes) > 0 {
		flash := flashes[0].(*Flash)
		session.Options.MaxAge = -1
		if err := session.Save(r, w); err != nil {
			return nil, err
		}
		return flash, nil
	}

	return nil, nil
}

// SetValue save value in session
func (s *Session) SetValue(w http.ResponseWriter, r *http.Request, name string, value interface{}) error {
	session, err := s.store.Get(r, s.Project+"_session_value")
	if err != nil {
		return err
	}

	session.Values[name] = value

	//set session configurations
	session.Options.MaxAge = int(time.Minute * 30)

	return session.Save(r, w)
}

// GetValue get value in session
func (s *Session) GetValue(r *http.Request, name string) (interface{}, error) {
	session, err := s.store.Get(r, s.Project+"_session_value")
	if err != nil {
		return nil, err
	}

	return session.Values[name], nil
}

// DeleteValue delete session
func (s *Session) DeleteValue(w http.ResponseWriter, r *http.Request) error {
	session, err := s.store.Get(r, s.Project+"_session_value")
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	return session.Save(r, w)
}

// SetSession save value in session
func (s *Session) SetSession(w http.ResponseWriter, r *http.Request, userID int) (string, error) {
	session, err := s.store.Get(r, s.Project+"_session_user")
	if err != nil {
		return "", err
	}

	gob.Register(&User{})

	UUID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	user := &User{
		UserID: userID,
		UUID:   UUID.String(),
	}

	session.Values["user"] = user

	//set session configurations
	session.Options.MaxAge = int(time.Minute * 30)

	return user.UUID, session.Save(r, w)
}

// GetSession get value in session
func (s *Session) GetSession(r *http.Request) (*User, error) {
	session, err := s.store.Get(r, s.Project+"_session_user")
	if err != nil {
		return nil, err
	}

	return session.Values["user"].(*User), nil
}

// DeleteSession delete session
func (s *Session) DeleteSession(w http.ResponseWriter, r *http.Request) error {
	session, err := s.store.Get(r, s.Project+"_session_user")
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	return session.Save(r, w)
}
