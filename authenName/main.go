package main

import (
	"authenName/config"
	"authenName/mock"
	"authenName/route"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Start serve...!\n")
	config.Connect()
	r := gin.Default()
	mock.Main()
	route.Router(r)
}
