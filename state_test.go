package yaygo

import (
	"os"
	"testing"
)

const baseDir = "./cmd/.test/"

func TestState(t *testing.T) {
	state := newState(baseDir, "test@example.com")

	t.Log("making directory")

	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		err := os.MkdirAll(baseDir, 0700)
		if err != nil {
			t.Fatal(err)
		}
	}

	t.Log("reading file")

	if err := state.Read(); os.IsNotExist(err) {
		t.Log("file not found. creating")
		err := state.Write()
		if err != nil {
			t.Fatal(err)
		}
	}

	if state.IsUnauthorized() {
		t.Log("is unauthorized")

		state.AccessToken = "access_token"
		state.RefreshToken = "refresh_token"
		state.UserID = 123

		t.Log("updating file")

		if err := state.Write(); err != nil {
			t.Fatal(err)
		}
	}

	t.Log(state)

	if err := state.Remove(); err != nil {
		t.Fatal(err)
	}

	t.Log("file removed.")
}
