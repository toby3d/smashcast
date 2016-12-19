package hitGox

import (
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
		Success bool `json:"success"`
		Error   bool `json:"error"`

		// Used in: CreateTeam, AcceptTeamInvite, LeaveFromTeam.
		SuccessMessage string `json:"success_msg,omitempty"`

		// Used in: CheckToken, GetStreamKey, ResetStreamKey, RunCommercialBreak, GetEditorsList, EditEditor, GetEditorList, GetHosters, GetLiveMedia, GetVideo, CreateVideo, CreateTeam, LeaveFromTeam, EditModerator, UpdateChatSettings.
		ErrorMessage string `json:"error_msg,omitempty"`

		// Used in: UpdateUserObject, SetDefaultTeam, EditEditor, SendTwitterPost, SendFacebookPost, FollowAChannel, UnfollowAChannel, EditModerator, UpdateChatSettings, UpdateWhisperSetting, UpdateUserAvatar, RemoveDescriptionImage.
		Message string `json:"message,omitempty"`

		// Used only in CheckToken function if token is valid.
		MSG string `json:"msg,omitempty"`
	}

	// Application is simple structure about hitbox app.
	Application struct {
		Name   string
		Token  string
		Secret string
	}
)

// NewApplication create Application structure for functions based on this.
func NewApplication(appName string, appToken string, appSecret string) *Application {
	application := &Application{
		Name:   appName,
		Token:  appToken,
		Secret: appSecret,
	}
	return application
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

func update(dst []byte, url string, args *fasthttp.Args) ([]byte, error) {
	return request("UPDATE", dst, url, args)
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
