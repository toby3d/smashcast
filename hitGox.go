package hitGox

const (
	// API is a URL for all API requests.
	API = "https://api.hitbox.tv"
	// Images is a URL for all images resources.
	Images = "https://edge.sf.hitbox.tv"
)

type (
	Application struct {
		Name   string
		Token  string
		Secret string
	}

	Status struct {
		Success        bool   `json:"success"`
		Error          bool   `json:"error"`
		SuccessMessage string `json:"success_msg,ommitempty"`
		ErrorMessage   string `json:"error_msg,ommitempty"`
		Message        string `json:"message,ommitempty"`
		MSG            string `json:"msg,ommitempty"`
	}

	Token struct {
		Token string
	}

	Request struct {
		This string `json:"this"`
	}
)

// NewApplication create a new simple application object.
func NewApplication(name string, token string, secret string) Application {
	app := Application{
		Name:   name,
		Token:  token,
		Secret: secret,
	}
	return app
}
