package yaygo

import (
	"net/http"
	"time"
)

const VERSION = "0.0.1"

type Session struct {
	client *http.Client

	User *UserApi
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

	s.User = newUserApi(s)

	return
}
