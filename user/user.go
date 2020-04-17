package user

const State = "WAITING"

type User struct {
	State string
}

func New() *User {
	return &User{
		State: State,
	}
}
