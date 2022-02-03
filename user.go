package onedrive

import (
	"fmt"
	"net/http"
)

// User represents an user in Microsoft Live.
type User struct {
	Id                string `json:"id"`
	DisplayName       string `json:"displayName"`
	UserPrincipalName string `json:"userPrincipalName"`
}

// GetCurrentAccountOutput request output.
type GetCurrentAccountOutput struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		Id                string `json:"id"`
		DisplayName       string `json:"displayName"`
		Surname           string `json:"surname"`
		GivenName         string `json:"givenName"`
		UserPrincipalName string `json:"userPrincipalName"`
		JobTitle          string `json:"jobTitle"`
		Mail              string `json:"mail"`
		MobilePhone       string `json:"mobilePhone"`
		OfficeLocation    string `json:"officeLocation"`
		PreferredLanguage string `json:"preferredLanguage"`
		Header            http.Header
	} `json:"value"`
}

// GetCurrentAccount returns information about the current user's account.
func (ds *DriveService) GetCurrentAccount() (*GetCurrentAccountOutput, *http.Response, error) {
	req, err := ds.OneDrive.newRequest("GET", "/users", nil, nil)
	if err != nil {
		return nil, nil, err
	}

	ca := new(GetCurrentAccountOutput)
	resp, err := ds.do(req, ca)
	fmt.Println("resp")
	fmt.Printf("\n%+v\n\n", ca)
	if err != nil {
		return nil, resp, err
	}

	return ca, resp, nil
}
