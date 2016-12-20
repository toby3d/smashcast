package hitGox

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

const (
	// API is a URL for all API requests.
	API = "https://api.hitbox.tv"

	// ImagesHost is a URL for all images resources.
	ImagesHost = "https://edge.sf.hitbox.tv"
)

type (
	// Status is a response body about successful or corrupted requests.
	Status struct {
		Success bool
		Error   bool
		Message string
	}

	// Application is simple structure of hitbox OAuth Application for authenticate actions.
	Application struct {
		Name   string
		Token  string
		Secret string
	}
)

// NewApplication create Application structure for functions based on this.
func NewApplication(appName string, appToken string, appSecret string) *Application {
	app := &Application{
		Name:   appName,
		Token:  appToken,
		Secret: appSecret,
	}
	return app
}

func get(url string, args *fasthttp.Args) ([]byte, error) {
	return request("GET", nil, url, args)
}

func post(dst []byte, url string, args *fasthttp.Args) ([]byte, error) {
	return request("POST", dst, url, args)
}

func put(dst []byte, url string, args *fasthttp.Args) ([]byte, error) {
	return request("PUT", dst, url, args)
}

func delete(url string, args *fasthttp.Args) ([]byte, error) {
	return request("DELETE", nil, url, args)
}

func update(url string, args *fasthttp.Args) ([]byte, error) {
	return request("UPDATE", nil, url, args)
}

func request(method string, dst []byte, url string, args *fasthttp.Args) ([]byte, error) {
	if args != nil {
		url += "?" + args.String()
	}

	var req fasthttp.Request
	req.Header.SetUserAgent("hitGox")
	req.Header.SetMethod(method)
	req.Header.SetContentType("application/json; charset=utf-8")
	req.SetRequestURI(url)
	if dst != nil {
		req.SetBody(dst)
	}

	var resp fasthttp.Response
	err := fasthttp.Do(&req, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func fuckYouNeedDecodeStatusFirst(ass []byte) (*Status, error) {
	var shit = struct {
		Success        bool   `json:"success"`
		Error          bool   `json:"error"`
		SuccessMessage string `json:"success_msg,omitempty"`
		ErrorMessage   string `json:"error_msg,omitempty"`
		ShortMessage   string `json:"msg,omitempty"`
		Message        string `json:"message,omitempty"`
	}{}
	if err := json.NewDecoder(bytes.NewReader(ass)).Decode(&shit); err != nil {
		return nil, err
	}

	var msg string
	switch {
	case shit.SuccessMessage != "":
		msg = shit.SuccessMessage
	case shit.ErrorMessage != "":
		msg = shit.ErrorMessage
	case shit.ShortMessage != "":
		msg = shit.ShortMessage
	case shit.Message != "":
		msg = shit.Message
	}

	return &Status{shit.Success, shit.Error, msg}, nil
}
