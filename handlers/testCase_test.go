package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var contentType string = "application/json"

func setupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/test_case", GetTestCase)
	r.POST("/test_case", PostTestCase)
	r.PUT("/test_case", PutTestCase)
	r.DELETE("/test_case", DeleteTestCase)

	return r
}

func TestGetTestCase(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	url := fmt.Sprintf("%s/test_case", ts.URL)
	resp, _ := http.Get(url)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPostTestCase(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	url := fmt.Sprintf("%s/test_case", ts.URL)
	resp, _ := http.Post(url, contentType, nil)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdateTestCase(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	url := fmt.Sprintf("%s/test_case", ts.URL)
	requestPut, _ := http.NewRequest(http.MethodPut, url, nil)
	client := &http.Client{}
	resp, _ := client.Do(requestPut)

	assert.Equal(t, resp.StatusCode, http.StatusOK)
}

func TestDeleteTestCase(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	url := fmt.Sprintf("%s/test_case", ts.URL)
	requestPut, _ := http.NewRequest(http.MethodDelete, url, nil)
	client := &http.Client{}
	resp, _ := client.Do(requestPut)

	assert.Equal(t, resp.StatusCode, http.StatusOK)
}
