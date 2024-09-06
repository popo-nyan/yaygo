package yaygo

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type State struct {
	mu sync.Mutex

	dir string
	key string

	UserID       int
	Email        string
	DeviceUUID   string
	AccessToken  string
	RefreshToken string
	IP           string
}

type Storage struct {
	UserID       int
	Email        string
	DeviceUUID   string
	AccessToken  string
	RefreshToken string
}

func newState(dir, email string) *State {
	return &State{
		dir: dir,
		key: generateKey(email),

		Email:      email,
		UserID:     -1, // Placeholder until API returns user ID
		DeviceUUID: generateUUID(),
	}
}

func generateUUID() string {
	uuid := make([]byte, 16)
	io.ReadFull(rand.Reader, uuid)

	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func (s *State) getFilename() string {
	h := sha256.New()
	h.Write([]byte(s.Email))
	return hex.EncodeToString(h.Sum(nil)) + ".gob"
}

func (s *State) IsUnauthorized() bool {
	return s.AccessToken == "" || s.RefreshToken == ""
}

func (s *State) Read() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	path := filepath.Join(s.dir, s.getFilename())

	if _, err := os.Stat(path); err != nil {
		return err
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var st Storage
	dec := gob.NewDecoder(file)
	if err := dec.Decode(&st); err != nil {
		return err
	}

	s.UserID = st.UserID
	s.Email = st.Email
	s.DeviceUUID = st.DeviceUUID
	s.AccessToken = st.AccessToken
	s.RefreshToken = st.RefreshToken

	return nil
}

func (s *State) Write() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	path := filepath.Join(s.dir, s.getFilename())

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	st := Storage{
		UserID:       s.UserID,
		Email:        s.Email,
		DeviceUUID:   s.DeviceUUID,
		AccessToken:  s.AccessToken,
		RefreshToken: s.RefreshToken,
	}

	enc := gob.NewEncoder(file)
	if err := enc.Encode(st); err != nil {
		return err
	}

	return nil
}

func (s *State) Remove() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	path := filepath.Join(s.dir, s.getFilename())
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}
