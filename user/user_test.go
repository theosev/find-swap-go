package user

import "testing"

func TestUserCreation(t *testing.T) {
	user := New(1)

	if user == nil {
		t.Fatalf("unexpected error from New: user is nil")
	}
}
