package main

import (
	"stably/internal/webapi"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	webapi.Route(r)
	r.Run()
}
