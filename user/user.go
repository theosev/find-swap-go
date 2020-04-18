package user

const State = "WAITING"

type User struct {
	State string
	Id    int
}

func New(id int) *User {
	return &User{
		State: State,
		Id:    id,
	}
}
