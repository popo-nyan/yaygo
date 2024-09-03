package yaygo

import (
	"net/http"
	"time"
)

const VERSION = "0.0.1"

type Session struct {
	client *http.Client

	Auth         *AuthApi
	Call         *CallApi
	Chat         *ChatApi
	Group        *GroupApi
	Hidden       *HiddenApi
	Misc         *MiscApi
	Notification *NotificationApi
	Post         *PostApi
	Review       *ReviewApi
	Thread       *ThreadApi
	User         *UserApi
}

type SessionConfig struct {
	client *http.Client
}

func newSessionConfig() *SessionConfig {
	return &SessionConfig{
		client: &http.Client{Timeout: (20 * time.Second)},
	}
}

type SessionOption func(cfg *SessionConfig)

func WithClient(client *http.Client) SessionOption {
	return func(cfg *SessionConfig) {
		if client != nil {
			cfg.client = client
		}
	}
}

func New(email, password string, options ...SessionOption) (s *Session, err error) {
	cfg := newSessionConfig()
	for _, opt := range options {
		opt(cfg)
	}

	s = &Session{
		client: &http.Client{Timeout: (20 * time.Second)},
	}

	s.Auth = newAuthApi(s)
	s.Call = newCallApi(s)
	s.Chat = newChatApi(s)
	s.Group = newGroupApi(s)
	s.Hidden = newHiddenApi(s)
	s.Misc = newMiscApi(s)
	s.Notification = newNotificationApi(s)
	s.Post = newPostApi(s)
	s.Review = newReviewApi(s)
	s.Thread = newThreadApi(s)
	s.User = newUserApi(s)

	return
}
