package webapi

import (
	"fmt"
	"math"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Number float64 `json:"number" form:"number"`
}

var primes []int

func Sieve(n float64) {
	nNew := int(math.Sqrt(n))
	marked := make([]int, int(n/2+500))

	m := int((nNew-1)/2 + 1)

	for i := 1; i <= m; i++ {
		for j := (i * (i + 1)) << 1; j <= int(n)/2; j = j + 2*i + 1 {
			marked[j] = 1
		}
	}

	primes = append(primes, 2)

	for i := 1; i <= int(n)/2; i++ {
		if marked[i] == 0 {
			primes = append(primes, 2*i+1)
		}
	}

}

func binarySearch(left int, right int, n int) int {
	if left <= right {
		mid := int((left + right) / 2)

		if mid == 0 || mid == len(primes)-1 {
			return primes[mid]
		}

		if primes[mid] == n {
			return primes[mid-1]
		}

		if primes[mid] < n && primes[mid+1] > n {
			return primes[mid]
		}
		if n < primes[mid] {
			return binarySearch(left, mid-1, n)
		} else {
			return binarySearch(mid+1, right, n)
		}
	}

	return 0
}

func FindHighestPrimeNumber(c *gin.Context) {
	var input CreateRequest
	if err := c.Bind(&input); err != nil {
		fmt.Println(err)
	}
	number := input.Number
	if number < 3 {
		c.JSON(200, gin.H{
			"error": "Please input number from 3 or more",
		})

		return
	}

	Sieve(number)

	result := binarySearch(0, len(primes)-1, int(number))

	c.JSON(200, gin.H{
		"result": result,
	})
}
