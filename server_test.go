package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldShouldSucceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)

	resp, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Hello World!", string(body))
	assert.Equal(t, 200, resp.StatusCode)
}

func TestHelloWorldShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)

	body := strings.NewReader("some body")

	resp, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, resp.StatusCode)
}

func TestHealthShouldSucceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)

	resp, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "OK", string(body))
	assert.Equal(t, 200, resp.StatusCode)
}

func TestHealthShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)

	body := strings.NewReader("some body")

	resp, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, resp.StatusCode)
}

func TestNewEndpointShouldSucceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleNewEndpoint))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)

	resp, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "OK", string(body))
	assert.Equal(t, 200, resp.StatusCode)
}

func TestNewEndpointShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleNewEndpoint))
	defer testServer.Close()

	testClient := testServer.Client()

	fmt.Println(testServer.URL)

	body := strings.NewReader("some body")

	resp, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, resp.StatusCode)
}
