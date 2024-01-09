package yaygo

import (
	"net/http"
	"time"
)

const VERSION string = "0.0.1"

func New(email string, password string) (s *Session, err error) {
	s = &Session{
		Client: &http.Client{Timeout: (20 * time.Second)},
	}

}
