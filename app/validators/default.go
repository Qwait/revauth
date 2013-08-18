package validators

import "github.com/robfig/revel"

func UniqueUsername(v *revel.Validation, username string) *revel.ValidationResult {
    return v.Check(username, revel.Required{})
}

func UniqueEmail(v *revel.Validation, email string) *revel.ValidationResult {
    return v.Check(email, revel.Required{}, revel.Email{})
}