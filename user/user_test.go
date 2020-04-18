package user

import "testing"

func TestUserCreation(t *testing.T) {
	user := New(1)

	if user.State != State {
		t.Fatalf("unexpected error from New: state is not waiting")
	}
}
