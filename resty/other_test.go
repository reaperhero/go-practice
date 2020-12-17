package resty__test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"testing"
)

func Test_delete(t *testing.T) {
	type Error struct {
	}
	// Create a Resty Client
	client := resty.New()

	// DELETE a article
	// No need to set auth token, error, if you have client level settings
	resp, err := client.R().
		SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
		SetError(&Error{}). // or SetError(Error{}).
		Delete("https://myapp.com/articles/1234")
	fmt.Println(resp.String(), err)
	// DELETE a articles with payload/body as a JSON string
	// No need to set auth token, error, if you have client level settings
	resp, err = client.R().
		SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
		SetError(&Error{}). // or SetError(Error{}).
		SetHeader("Content-Type", "application/json").
		SetBody(`{article_ids: [1002, 1006, 1007, 87683, 45432] }`).
		Delete("https://myapp.com/articles")

	// HEAD of resource
	// No need to set auth token, if you have client level settings
	resp, err = client.R().
		SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
		Head("https://myapp.com/videos/hi-res-video")

	// OPTIONS of resource
	// No need to set auth token, if you have client level settings
	resp, err = client.R().
		SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
		Options("https://myapp.com/servers/nyc-dc-01")
}
