package main

import (
	"authenName/config"
	"authenName/route"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Start serve...!\n")
	config.Connect()

	r := gin.Default()
	route.Router(r)
}
