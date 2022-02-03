package onedrive

import (
	"context"
	"errors"
	"net/http"
	"net/url"
)

// PermissionService handles permission settings of a drive item
type PermissionService struct {
	*OneDrive
}

// Permission is the permission of a drive item.
type Permission struct {
	ID        string      `json:"id"`
	GrantedTo interface{} `json:"grantedTo"`
	Link      SharingLink `json:"link"`
	Roles     []string    `json:"roles"`
}

// CreateShareLinkRequest is the request for creating a share link.
type CreateShareLinkRequest struct {
	Type  string `json:"type"`  // The type of sharing link to create. Either view, edit, or embed.
	Scope string `json:"scope"` // Optional. The scope of link to create. Either anonymous or organization.
}

// CreateShareLink will create a new sharing link if the specified link type doesn't already exist for the calling application.
func (s *PermissionService) CreateShareLink(ctx context.Context, itemId string, permissionType ShareLinkType, permissionScope ShareLinkScope) (*Permission, error) {
	apiURL := "me/drive/items/" + url.PathEscape(itemId) + "/createLink"

	body := &CreateShareLinkRequest{Type: permissionType.toString(), Scope: permissionScope.toString()}
	req, err := s.newRequest(http.MethodPost, apiURL, nil, body)
	if err != nil {
		return nil, err
	}

	var oneDriveResponse *Permission
	_, err = s.do(req, &oneDriveResponse)
	if err != nil {
		return nil, err
	}

	return oneDriveResponse, nil
}

// ListPermissionsResponse is the response of list permissions of a drive item
type ListPermissionsResponse struct {
	Value []Permission `json:"value"`
}

// List lists the effective sharing permissions of on a DriveItem.
func (s *PermissionService) List(ctx context.Context, itemId string) ([]Permission, error) {
	apiURL := "me/drive/items/" + url.PathEscape(itemId) + "/permissions"

	req, err := s.newRequest(http.MethodGet, apiURL, nil, nil)
	if err != nil {
		return nil, err
	}

	var oneDriveResponse *ListPermissionsResponse
	_, err = s.do(req, &oneDriveResponse)
	if err != nil {
		return nil, err
	}

	return oneDriveResponse.Value, nil
}

// Delete will delete a sharing permission from a file or folder.
func (s *PermissionService) Delete(ctx context.Context, driveId string, itemId string, permissionId string) error {
	if itemId == "" {
		return errors.New("Please provide the Item ID of the item to be deleted.")
	}

	apiURL := "me/drive/items/" + url.PathEscape(itemId)
	if driveId != "" {
		apiURL = "me/drives/" + url.PathEscape(driveId) + "/items/" + url.PathEscape(itemId)
	}

	apiURL += "/permissions/" + url.PathEscape(permissionId)

	req, err := s.newRequest("DELETE", apiURL, nil, nil)
	if err != nil {
		return err
	}

	var driveItem *Item
	_, err = s.do(req, &driveItem)
	if err != nil {
		return err
	}

	return nil
}
