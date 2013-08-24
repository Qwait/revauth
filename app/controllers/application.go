package controllers

import (
    "fmt"
    "reflect"
    r "github.com/robfig/revel"
    "github.com/qwait/gorden"
)

type AuthenticateParams struct {
    Email string
    Password string
}

type PasswordStrategy struct {}

func (strategy PasswordStrategy) Authenticate(arguments interface{}) bool {
    v := reflect.ValueOf(arguments)

    fmt.Println(v.FieldByName("Email"))
    fmt.Println(v.FieldByName("Password"))

    return false
}

func (PasswordStrategy) IsAuthenticated() bool {
    return false
}

type ApplicationController struct {
    *r.Controller
}

func (c *ApplicationController) GordenManager() *gorden.Manager {
    secret, _ := r.Config.String("app.secret")

    return gorden.NewManager("password", gorden.SessionConfig{CookieName: "auth_token", CookieKey: []byte(secret)})
}