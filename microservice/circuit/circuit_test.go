package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/cep21/circuit/v3"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCircuit(t *testing.T) {
	h := circuit.Manager{}
	c := h.MustCreateCircuit("hello-http", circuit.Config{
		Execution: circuit.ExecutionConfig{
			Timeout: time.Second * 3,
		},
	})

	testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(rw, "hello world")
	}))
	defer testServer.Close()

	var body bytes.Buffer
	runErr := c.Run(context.Background(), func(ctx context.Context) error {
		req, err := http.NewRequest("GET", testServer.URL, nil)
		if err != nil {
			return circuit.SimpleBadRequest{Err: err}
		}
		req = req.WithContext(ctx)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode >= 400 && resp.StatusCode <= 499 {
			return circuit.SimpleBadRequest{Err: errors.New("server found your request invalid")}
		}
		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			return fmt.Errorf("invalid status code: %d", resp.StatusCode)
		}
		if _, err := io.Copy(&body, resp.Body); err != nil {
			return err
		}
		return resp.Body.Close()
	})
	if runErr == nil {
		fmt.Printf("We saw a body\n")
	}
}
