package resty__test

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
)

func () {
	profileImgBytes, _ := ioutil.ReadFile("/Users/jeeva/test-img.png")
	notesBytes, _ := ioutil.ReadFile("/Users/jeeva/text-file.txt")

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		SetFileReader("profile_img", "test-img.png", bytes.NewReader(profileImgBytes)).
		SetFileReader("notes", "text-file.txt", bytes.NewReader(notesBytes)).
		SetFormData(map[string]string{
			"first_name": "Jeevanandam",
			"last_name":  "M",
		}).
		Post("http://myapp.com/upload")
	fmt.Println(resp.String(), err)
}
