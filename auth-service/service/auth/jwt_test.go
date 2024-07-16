package auth

import (
	"testing"
)

func TestCreateJWT(t *testing.T) {
	userId := "1f8d850d-f857-4fea-af66-937f54e8750c"

	token, err := CreateJWT(userId)
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Error("expected token to be not empty")
	}
}
