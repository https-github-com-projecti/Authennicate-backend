package main

import (
	"fmt"
	"authenName/config"
	"github.com/gin-gonic/gin"
	"authenName/route"
)

func main() {
	fmt.Print("String serve...!\n")
	config.Connect()

	r := gin.Default()
	route.Router(r)
}
