package webapi

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine) *gin.Engine {
	r.POST("/find-highest-prime-number", FindHighestPrimeNumber)

	return r
}
