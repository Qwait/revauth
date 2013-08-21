package revauth

import (
    "fmt"
    "github.com/robfig/revel"
)

func init() {
    revel.OnAppStart(func() {
        fmt.Println("Revauth Started")
    })
}