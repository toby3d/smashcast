package hitGox

const (
	// API is a URL for all API requests.
	API = "https://api.hitbox.tv"

	// Images is a URL for all images resources.
	Images = "https://edge.sf.hitbox.tv"
)

type (
	// Application is simple structure about hitbox app.
	Application struct {
		Name   string
		Token  string
		Secret string
	}

	// Status is a response body about successful or corrupted requests.
	Status struct {
		Success bool `json:"success"`
		Error   bool `json:"error"`

		// Message about suceffuly actions. Used only in CreateTeam, AcceptTeamInvite, TeamKick and TeamLeave.
		SuccessMessage string `json:"success_msg,omitempty"`

		// Message about error, only in corrupted requests.
		ErrorMessage string `json:"error_msg,omitempty"`

		// Message about suceffuly actions.
		Message string `json:"message,omitempty"`

		// Used only in CheckToken if AuthToken is valid.
		ShortMessage string `json:"msg,omitempty"`
	}

	// Request is a string about current request. Maybe can be used for logs.
	Request struct {
		This string `json:"this"`
	}

	Timestamp struct {
		time.Time
	}

	// Token is a universal structure for AccessToken and AuthToken.
	Token struct {
		Token string
	}
)

// NewApplication create a new simple application object.
func NewApplication(appName string, appToken string, appSecret string) Application {
	app := Application{
		Name:   name,
		Token:  appToken,
		Secret: secret,
	}
	return app
}
