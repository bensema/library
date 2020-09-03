package example

import "time"

type User struct {
	ID       int
	Name     string
	Age      int
	Birthday time.Duration
}

func (u *User) FieldsName() []string {
	return []string{"id", "name", "age", "birthday"}
}
func (u *User) FieldsItem() []interface{} {
	return []interface{}{&u.ID, &u.Name, &u.Age, &u.Birthday}
}
