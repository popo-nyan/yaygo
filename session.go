package yaygo

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

const (
	VERSION = "0.0.1"

	API_VERSION_NAME = "3.51"
	VERSION_NAME     = "3.51.0"
	API_VERSION_KEY  = "8befe6044b504adca957d5c4fb5dba86"
	API_KEY          = "ccd59ee269c01511ba763467045c115779fcae3050238a252f1bd1a4b65cfec6"
)

type Session struct {
	baseDir string

	client *http.Client
	state  *State

	Auth         *AuthAPI
	Call         *CallAPI
	Chat         *ChatAPI
	Group        *GroupAPI
	Hidden       *HiddenAPI
	Misc         *MiscAPI
	Notification *NotificationAPI
	Post         *PostAPI
	Review       *ReviewAPI
	Thread       *ThreadAPI
	User         *UserAPI
}

type SessionConfig struct {
	baseDir string

	client *http.Client
}

// Default session configuration.
func newSessionConfig() *SessionConfig {
	return &SessionConfig{
		baseDir: ".config/",

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

func WithBaseDir(baseDir string) SessionOption {
	return func(cfg *SessionConfig) {
		if baseDir != "" {
			cfg.baseDir = baseDir
		}
	}
}

func New(email, password string, options ...SessionOption) (*Session, error) {
	cfg := newSessionConfig()
	for _, opt := range options {
		opt(cfg)
	}

	s := &Session{
		baseDir: cfg.baseDir,

		client: cfg.client,
		state:  newState(cfg.baseDir, email),
	}

	s.Auth = newAuthAPI(s)
	s.Call = newCallAPI(s)
	s.Chat = newChatAPI(s)
	s.Group = newGroupAPI(s)
	s.Hidden = newHiddenAPI(s)
	s.Misc = newMiscAPI(s)
	s.Notification = newNotificationAPI(s)
	s.Post = newPostAPI(s)
	s.Review = newReviewAPI(s)
	s.Thread = newThreadAPI(s)
	s.User = newUserAPI(s)

	if err := s.init(email, password); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Session) init(email, password string) error {
	if _, err := os.Stat(s.baseDir); os.IsNotExist(err) {
		return os.MkdirAll(s.baseDir, 0700)
	}

	if err := s.state.Read(); os.IsNotExist(err) {
		slog.Debug("State file not found. Creating new one.")
		err := s.state.Write()
		if err != nil {
			return err
		}
	}

	if s.state.IsUnauthorized() {
		resp, err := s.Auth.Login(&LoginParams{Email: email, Password: password})
		if err != nil {
			return err
		}

		s.state.AccessToken = resp.AccessToken
		s.state.RefreshToken = resp.RefreshToken
		s.state.UserID = resp.UserID

		slog.Debug("Saving authenticated state.")

		if err := s.state.Write(); err != nil {
			return err
		}
	}

	slog.Debug("Sesssion intialized.")

	return nil
}
