package tools

import (
	"bytes"
	"encoding/json"
	"fmt"

	f "github.com/valyala/fasthttp"
)

const (
	// APIEndpoint is a URL for all API requests.
	APIEndpoint = "https://api.hitbox.tv/%s"

	// ImagesEndpoint is a URL for all images resources.
	ImagesEndpoint = "https://edge.sf.hitbox.tv"
)

type (
	Status struct {
		Success bool
		Error   bool
		Message string
	}
)

func GET(url string, args *f.Args) ([]byte, error) {
	return request("GET", nil, url, args)
}

func POST(dst []byte, url string, args *f.Args) ([]byte, error) {
	return request("POST", dst, url, args)
}

func PUT(dst []byte, url string, args *f.Args) ([]byte, error) {
	return request("PUT", dst, url, args)
}

func DELETE(url string, args *f.Args) ([]byte, error) {
	return request("DELETE", nil, url, args)
}

func UPDATE(url string, args *f.Args) ([]byte, error) {
	return request("UPDATE", nil, url, args)
}

func request(method string, dst []byte, url string, args *f.Args) ([]byte, error) {
	if args != nil {
		url += "?" + args.String()
	}

	var req f.Request
	req.Header.SetUserAgent("hitGox")
	req.Header.SetMethod(method)
	req.Header.SetContentType("application/json; charset=utf-8")
	req.SetRequestURI(url)
	if dst != nil {
		req.SetBody(dst)
	}

	var resp f.Response
	err := f.Do(&req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("%s", FixStatus(resp.Body()).Message)
	}

	return resp.Body(), nil
}

func FixStatus(resp []byte) *Status {
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
