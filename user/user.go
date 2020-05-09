package user


const state = "WAITING"

// User struct
type User struct {
	state string
	id    int
	ComChannel chan byte
}

// New creates a new anonymous user
func New(id int) *User {
	return &User{
		state: state,
		id:    id,
	}
}