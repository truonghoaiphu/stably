package webapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var MAX = 7

var primes []int

func Sieve() {

	var marked [4]int
	fmt.Println(marked)

}

func FindHighestPrimeNumber(c *gin.Context) {
	go Sieve()
}
