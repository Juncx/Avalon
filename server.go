package main

import (
	"Avalon/avalon"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	avalon.Router(r)

	//go r.RunTLS(":443", "../ssl/server.crt", "../ssl/server.key")
	r.Run(":8080")
}
