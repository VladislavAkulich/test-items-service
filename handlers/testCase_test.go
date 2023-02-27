package handlers_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"example/test-items-service/mapping"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var contentType string = "application/json"

// TODO rewrite tests according to new changes
func setupServer() *gin.Engine {
	r := gin.Default()
	mapping.ApplyTestCaseMapping(context.Background(), r)
	return r
}

func TestGetTestCase(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	url := fmt.Sprintf("%s%s", ts.URL, mapping.RootTestCasePath)
	resp, _ := http.Get(url)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetTestCaseById(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	url := fmt.Sprintf("%s%s%s", ts.URL, mapping.RootTestCasePath, "/4a93b714-0252-489d-8854-9b2a1ae2bb36")
	resp, _ := http.Get(url)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPostTestCase(t *testing.T) {
	// ts := httptest.NewServer(setupServer())
	// defer ts.Close()

	//url := fmt.Sprintf("%s%s", ts.URL, mapping.RootTestCasePath)

	url := fmt.Sprintf("%s%s", "http://localhost:8080", mapping.RootTestCasePath)
	body := []byte(`{
		"name": "test case post",
		"steps": [
			"step post 1",
			"step post 2"
		],
		"preconditions": "preconditions post",
		"author": "author post"
	}`)
	fmt.Println(url)
	resp, _ := http.Post(url, contentType, bytes.NewBuffer(body))

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdateTestCase(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	url := fmt.Sprintf("%s%s", ts.URL, mapping.RootTestCasePath)
	requestPut, _ := http.NewRequest(http.MethodPut, url, nil)
	client := &http.Client{}
	resp, _ := client.Do(requestPut)

	assert.Equal(t, resp.StatusCode, http.StatusOK)
}

func TestDeleteTestCase(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	url := fmt.Sprintf("%s%s%s", ts.URL, mapping.RootTestCasePath, "/4a93b714-0252-489d-8854-9b2a1ae2bb36")
	requestDelete, _ := http.NewRequest(http.MethodDelete, url, nil)
	client := &http.Client{}
	resp, _ := client.Do(requestDelete)

	assert.Equal(t, resp.StatusCode, http.StatusOK)
}
