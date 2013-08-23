package auth

import "github.com/qwait/gorden"

func init() {
    gorden.AddStrategy("password", PasswordStrategy{})
}

type PasswordStrategy struct {}
    
func (PasswordStrategy) Authenticate(arguments interface{}) bool {
    return false
}

func (PasswordStrategy) IsAuthenticated() bool {
    return false
}