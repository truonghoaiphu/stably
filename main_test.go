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

	var jsonStr = []byte(`{"number":55}`)
	req, err := http.NewRequest("POST", "/find-highest-prime-number", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	expected := `{"result":53}`
	assert.Equal(t, rr.Body.String(), expected)
}
