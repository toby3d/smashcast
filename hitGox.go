package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
)

const (
	// APIEndpoint is a URL for all API requests.
	APIEndpoint = "https://api.hitbox.tv/%s"

	// ImagesEndpoint is a URL for all images resources.
	ImagesEndpoint = "https://edge.sf.hitbox.tv/%s"
)

// Status is a response body about successful or corrupted requests.
type Status struct {
	Success bool
	Error   bool
	Message string
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

	if resp.StatusCode() != 200 {
		return nil, errors.New(fixStatus(resp.Body()).Message)
	}

	return resp.Body(), nil
}

func fixStatus(resp []byte) *Status {
	var obj = struct {
		Success        bool   `json:"success"`
		Error          bool   `json:"error"`
		SuccessMessage string `json:"success_msg,omitempty"`
		ErrorMessage   string `json:"error_msg,omitempty"`
		ShortMessage   string `json:"msg,omitempty"`
		Message        string `json:"message,omitempty"`
	}{}
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	var msg string
	switch {
	case obj.SuccessMessage != "":
		msg = obj.SuccessMessage
	case obj.ErrorMessage != "":
		msg = obj.ErrorMessage
	case obj.ShortMessage != "":
		msg = obj.ShortMessage
	case obj.Message != "":
		msg = obj.Message
	}

	return &Status{
		Success: obj.Success,
		Error:   obj.Error,
		Message: msg,
	}
}
