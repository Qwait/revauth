package controllers

import (
    "github.com/qwait/gorden"
    r "github.com/robfig/revel"
)

type ApplicationController struct {
    *r.Controller
}

func (c *ApplicationController) getGordenManager() *gorden.Manager {
    secret, _ := r.Config.String("app.secret")

    return gorden.NewManager("password", gorden.SessionConfig{CookieName: "auth_token", CookieKey: []byte(secret)})
}