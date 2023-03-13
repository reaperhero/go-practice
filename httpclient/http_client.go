package httpclient

import (
	"bytes"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	MaxIdleConn        int = 100
	MaxIdleConnPerHost int = 100
	IdleConnTimeout    int = 90
)

var (
	httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConn,
			MaxIdleConnsPerHost: MaxIdleConnPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},
	}
)

type Config struct {
	Domain string
	Token  string
}

func (c *Config) Get(api string, data map[string]interface{}) (json *simplejson.Json, err error) {
	var (
		endPoint = c.Domain + api
		params   = url.Values{}
	)
	endpoint, err := url.Parse(endPoint)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	for k, v := range data {
		params.Set(k, fmt.Sprintf("%v", v))
	}
	endpoint.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.AddCookie(&http.Cookie{Name: "token_test", Value: c.Token, HttpOnly: true})

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	js, _ := simplejson.NewJson(b)
	return js, nil
}

func (c *Config) PostJson(api string, jsonData []byte) (json *simplejson.Json, err error) {
	var (
		endPoint = c.Domain + api
	)
	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.AddCookie(&http.Cookie{Name: "token_test", Value: c.Token, HttpOnly: true})


	resp, err := httpClient.Do(req)
	if err != nil || resp == nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	js, _ := simplejson.NewJson(b)
	return js, nil
}

func (c *Config) PostForm(api string, data map[string]interface{}) (json *simplejson.Json, err error) {
	var (
		endPoint = c.Domain + api
	)

	req, err := http.NewRequest("POST", endPoint, strings.NewReader("name=cjb"))
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	req.Form = make(url.Values)
	for k, v := range data {
		req.Form.Add(k, fmt.Sprintf("%v", v))
	}
	req.AddCookie(&http.Cookie{Name: "token_test", Value: c.Token, HttpOnly: true})

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	js, _ := simplejson.NewJson(b)
	return js, nil
}

func (c *Config) PostFormBytes(api string, data map[string]interface{}) (json *simplejson.Json, err error) {
	var (
		endPoint = c.Domain + api
	)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	for s, i := range data {
		_ = writer.WriteField(s, i.(string))
	}

	err = writer.Close()
	if err != nil {
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", endPoint, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "multipart/form-data; boundary=----WebKitFormBoundaryVzHZonpLIM9ywinP")

	req.Header.Set("Content-Type", writer.FormDataContentType())


	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	js, _ := simplejson.NewJson(body)
	return js, nil
}


