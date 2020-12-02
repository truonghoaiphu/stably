package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"stably/internal/webapi"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type App struct {
	lists []webapi.CreateRequest
}

func TestFindHighestPrimeNumber(t *testing.T) {
	r := gin.Default()
	router := webapi.Route(r)
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty number",
			input:    `{"number"}`,
			expected: `{"error":"Please input number from 3 or more"}`,
		},
		{
			name:     "Negative",
			input:    `{"number: -8"}`,
			expected: `{"error":"Please input number from 3 or more"}`,
		},
		{
			name:     "Number less than 3",
			input:    `{"number":2}`,
			expected: `{"error":"Please input number from 3 or more"}`,
		},
		{
			name:     "Number greater than 3",
			input:    `{"number":55}`,
			expected: `{"result":53}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var jsonStr = []byte(testCase.input)
			req, err := http.NewRequest("POST", "/find-highest-prime-number", bytes.NewBuffer(jsonStr))
			if err != nil {
				fmt.Println(err)
			}
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			expected := testCase.expected
			assert.Equal(t, rr.Body.String(), expected)
		})
	}

}
