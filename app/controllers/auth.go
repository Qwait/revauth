package controllers

import (
    "github.com/robfig/revel"
    "github.com/qwait/revauth/app/models"
)

type Auth struct {
    ApplicationController
}

func (c Auth) Login() revel.Result {
    return c.Render()
}

func (c Auth) HandleLogin(user *models.User) revel.Result {
    user.ValidateLogin(c.Validation)

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        c.Flash.Error("Incorrect username or password.")
        return c.Redirect("login")
    }

    c.Flash.Success("Login Worked")
    return c.Redirect("login")
}

func (c Auth) ForgotPassword() revel.Result {
    return c.Render()
}

func (c Auth) HandleForgotPassword(email string) revel.Result {
    c.Validation.Required(email)
    c.Validation.Email(email)

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        c.Flash.Error("Account not found.")
        return c.Redirect("forgot")
    }

    c.Flash.Success("Password reset email sent.")
    return c.Redirect("forgot")
}

func (c Auth) ResetPassword(user_id int, hmac string) revel.Result {
    return c.Render()
}

func (c Auth) HandleResetPassword(user_id int, hmac string) revel.Result {
    return c.Render()
}

func (c Auth) Logout() revel.Result {
    return c.Redirect("login")
}

func (c Auth) Register() revel.Result {
    return c.Render()
}

func (c Auth) HandleRegister(user *models.User) revel.Result {
    user.ValidateCreate(c.Validation)

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect("register")
    }

    return c.Redirect("register")
}

func (c Auth) ChangePassword() revel.Result {
    return c.Render()
}

func (c Auth) HandleChangePassword() revel.Result {
    c.Flash.Success("Password reset.")
    return c.Redirect("login")
}