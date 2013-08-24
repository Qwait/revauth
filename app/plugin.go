package revauth

import (
    "fmt"
    "github.com/qwait/gorden"
    "github.com/qwait/revauth/app/controllers"
    "github.com/robfig/revel"
)

func init() {
    revel.OnAppStart(func() {
        gorden.AddStrategy("password", controllers.PasswordStrategy{})
        fmt.Println("Revauth Started")
    })
}