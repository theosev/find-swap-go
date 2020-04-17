package user

import "testing"

func TestUserCreation(t *testing.T) {
	user := New()

	if user.State != State {
		t.Fatalf("unexpected error from New: state is not waiting")
	}
}
