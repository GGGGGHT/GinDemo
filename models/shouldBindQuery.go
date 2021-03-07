package models

import "fmt"

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}

type Register struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Phone    int    `form:"phone" json:"phone" binding:"required"`
}

func (r *Register) String() string {
	return fmt.Sprintf("name: " + r.UserName + ",password: " + r.Password + ",phone: " + string(rune(r.Phone)))
}
