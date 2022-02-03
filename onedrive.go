package onedrive

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	version   = "0.1"
	baseURL   = "https://graph.microsoft.com/v1.0"
	userAgent = "Sendergram"
)

// OneDrive is the entry point for the client. It manages the communication with
// Microsoft OneDrive API
type OneDrive struct {
	Client *http.Client
	// When debug is set to true, the JSON response is formatted for better readability
	Debug     bool
	BaseURL   string
	UserAgent string
	// Services
	Drives           *DriveService
	Items            *ItemService
	DrivePermissions *PermissionService
	// Private
	throttle time.Time
	// User
	User *User
}

// NewOneDrive returns a new OneDrive client to enable you to communicate with
// the API
func NewOneDrive(c *http.Client, debug bool) *OneDrive {
	drive := OneDrive{
		Client:   c,
		BaseURL:  baseURL,
		Debug:    debug,
		throttle: time.Now(),
	}
	drive.Drives = &DriveService{&drive}
	drive.Items = &ItemService{&drive}
	drive.DrivePermissions = &PermissionService{&drive}
	return &drive
}

func (od *OneDrive) throttleRequest(time time.Time) {
	od.throttle = time
}

type errorReply struct {
	Error *Error `json:"error"`
}

// CheckResponse returns an error (of type *Error) if the response
// status code is not 2xx.
func CheckResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	slurp, err := ioutil.ReadAll(res.Body)
	if err == nil {
		jerr := new(errorReply)
		err = json.Unmarshal(slurp, jerr)
		if err == nil && jerr.Error != nil {
			if jerr.Error.Code == "" {
				jerr.Error.Code = res.Status
			}
			jerr.Error.Message = string(slurp)
			return jerr.Error
		}
	}
	return &Error{
		innerError: innerError{
			Code:    res.Status,
			Message: string(slurp),
		},
	}
}
