package hitGox

const (
	// API is a URL for all API requests.
	API = "https://api.hitbox.tv"

	// Images is a URL for all images resources.
	Images = "https://edge.sf.hitbox.tv"
)

type (
	// Application is a simple
	Application struct {
		Name   string
		Token  string
		Secret string
	}

	// Status is a response body about successful or corrupted requests.
	Status struct {
		Success bool `json:"success"`
		Error   bool `json:"error"`

		// Message aboud suceffuly actions.
		// Used only in CreateTeam, AcceptTeamInvite, TeamKick and TeamLeave.
		SuccessMessage string `json:"success_msg,ommitempty"`

		// Message about error, only in corrupted requests.
		ErrorMessage string `json:"error_msg,ommitempty"`

		// Message about suceffuly actions.
		Message string `json:"message,ommitempty"`

		// Used only CheckToken if AuthToken is valid.
		MSG string `json:"msg,ommitempty"`
	}

	// Token is a universal structure for AccessToken and AuthToken.
	Token struct {
		Token string
	}

	// Request is a string about current request. Maybe can be used for logs.
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
