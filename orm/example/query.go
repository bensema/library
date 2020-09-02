package example

import (
	"fmt"
	"library/orm"
)

func query() (err error) {
	user := &User{}
	o := orm.New()

	err = o.Table("t_user").Query().Where("name = ?", "ben").Gen(user)
	if err != nil {
		return err
	}

	o1 := orm.New()
	err = o1.Table("t_user").Query().Select().Where("name = ?", "ben").Gen(user)
	if err != nil {
		return err
	}

	o2 := orm.New()
	err = o2.Table("t_user").Select().Where("name = ?", "ben").Gen(user)
	if err != nil {
		return err
	}

	o3 := orm.New()
	err = o3.Table("t_user").Select("name", "birthday").Where("name = ?", "ben").Gen(user)
	if err != nil {
		return err
	}

	o4 := orm.New()
	err = o4.Table("t_user").Query().Where("name = ? AND age >= ?", "ben", 18).Gen(user)
	if err != nil {
		return err
	}
	fmt.Println(o4.String())
	return
}
