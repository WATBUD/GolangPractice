// Code generated by goa v3.16.2, DO NOT EDIT.
//
// HTTP request path constructors for the folder service.
//
// Command:
// $ goa gen mai.today/api/design

package server

import (
	"fmt"
)

// ReceiveCreateFolderFolderPath returns the URL path to the folder service receiveCreateFolder HTTP endpoint.
func ReceiveCreateFolderFolderPath(channel string) string {
	return fmt.Sprintf("/%v/createFolder", channel)
}

// CreateFolderFolderPath returns the URL path to the folder service CreateFolder HTTP endpoint.
func CreateFolderFolderPath() string {
	return "/folders"
}

// ReceiveDeleteFolderFolderPath returns the URL path to the folder service receiveDeleteFolder HTTP endpoint.
func ReceiveDeleteFolderFolderPath(channel string) string {
	return fmt.Sprintf("/%v/deleteFolder", channel)
}

// DeleteFolderFolderPath returns the URL path to the folder service DeleteFolder HTTP endpoint.
func DeleteFolderFolderPath(folderID string) string {
	return fmt.Sprintf("/folders/%v", folderID)
}

// ReceiveUpdateFolderFolderPath returns the URL path to the folder service receiveUpdateFolder HTTP endpoint.
func ReceiveUpdateFolderFolderPath(channel string) string {
	return fmt.Sprintf("/%v/updateFolder", channel)
}

// UpdateFolderFolderPath returns the URL path to the folder service UpdateFolder HTTP endpoint.
func UpdateFolderFolderPath(folderID string) string {
	return fmt.Sprintf("/folders/%v", folderID)
}

// ReceiveListFoldersFolderPath returns the URL path to the folder service receiveListFolders HTTP endpoint.
func ReceiveListFoldersFolderPath(channel string) string {
	return fmt.Sprintf("/%v/listFolders", channel)
}

// ListFoldersFolderPath returns the URL path to the folder service ListFolders HTTP endpoint.
func ListFoldersFolderPath() string {
	return "/folders"
}
