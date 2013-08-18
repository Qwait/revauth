package models

import (
    "github.com/robfig/revel"
    "github.com/qwait/revauth/app/validators"
)

type User struct {
    Username, Email, Password, PasswordConfirm string
}

func (user *User) ValidateCreate(v *revel.Validation) {
    validators.UniqueUsername(v, user.Username).Key("user.Username")
    validators.UniqueEmail(v, user.Email).Key("user.Email")

    v.Required(user.Password)
    v.Required(user.PasswordConfirm)

    v.Required(user.PasswordConfirm == user.Password).Message("Passwords must match")
}

func (user *User) ValidateLogin(v *revel.Validation) {
    v.Required(user.Email)
    v.Required(user.Password)
    v.Email(user.Email)
}