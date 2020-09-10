package example

import "time"

type User struct {
	ID       int
	Name     string
	Age      int
	Birthday time.Duration
}

func (u *User) Table() string {
	return "t_user"
}

func (u *User) Columns() []string {
	return []string{"id", "name", "age", "birthday"}
}

func (u *User) Fields() []interface{} {
	return []interface{}{&u.ID, &u.Name, &u.Age, &u.Birthday}
}
