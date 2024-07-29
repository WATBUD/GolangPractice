// Code generated by goa v3.16.2, DO NOT EDIT.
//
// folder HTTP server types
//
// Command:
// $ goa gen mai.today/api/design

package server

import (
	goa "goa.design/goa/v3/pkg"
	folder "mai.today/api/gen/folder"
)

// CreateFolderRequestBody is the type of the "folder" service "CreateFolder"
// endpoint HTTP request body.
type CreateFolderRequestBody struct {
	Folder *FolderRequestBody `form:"folder,omitempty" json:"folder,omitempty" xml:"folder,omitempty"`
}

// UpdateFolderRequestBody is the type of the "folder" service "UpdateFolder"
// endpoint HTTP request body.
type UpdateFolderRequestBody struct {
	// Folder data to update
	Folder *FolderRequestBody `form:"folder,omitempty" json:"folder,omitempty" xml:"folder,omitempty"`
}

// CreateFolderOKResponseBody is the type of the "folder" service
// "CreateFolder" endpoint HTTP response body.
type CreateFolderOKResponseBody struct {
	Command *CommandResponseBody `form:"command" json:"command" xml:"command"`
	// 時間戳記
	Timestamp int64 `form:"timestamp" json:"timestamp" xml:"timestamp"`
	// 資料
	Data *CreatefolderresultdataResponseBody `form:"data" json:"data" xml:"data"`
}

// DeleteFolderOKResponseBody is the type of the "folder" service
// "DeleteFolder" endpoint HTTP response body.
type DeleteFolderOKResponseBody struct {
	Command *CommandResponseBody `form:"command" json:"command" xml:"command"`
	// 時間戳記
	Timestamp int64 `form:"timestamp" json:"timestamp" xml:"timestamp"`
	// 資料
	Data *CreatefolderresultdataResponseBody `form:"data" json:"data" xml:"data"`
}

// UpdateFolderOKResponseBody is the type of the "folder" service
// "UpdateFolder" endpoint HTTP response body.
type UpdateFolderOKResponseBody struct {
	Command *CommandResponseBody `form:"command" json:"command" xml:"command"`
	// 時間戳記
	Timestamp int64 `form:"timestamp" json:"timestamp" xml:"timestamp"`
	// 資料
	Data *CreatefolderresultdataResponseBody `form:"data" json:"data" xml:"data"`
}

// ListFoldersOKResponseBody is the type of the "folder" service "ListFolders"
// endpoint HTTP response body.
type ListFoldersOKResponseBody struct {
	Command *CommandResponseBody `form:"command" json:"command" xml:"command"`
	// 時間戳記
	Timestamp int64 `form:"timestamp" json:"timestamp" xml:"timestamp"`
	// 資料
	Data []*CreatefolderresultdataResponseBody `form:"data" json:"data" xml:"data"`
}

// CreateFolderInvalidTokenResponseBody is the type of the "folder" service
// "CreateFolder" endpoint HTTP response body for the "invalid token" error.
type CreateFolderInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteFolderInvalidTokenResponseBody is the type of the "folder" service
// "DeleteFolder" endpoint HTTP response body for the "invalid token" error.
type DeleteFolderInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateFolderInvalidTokenResponseBody is the type of the "folder" service
// "UpdateFolder" endpoint HTTP response body for the "invalid token" error.
type UpdateFolderInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListFoldersInvalidTokenResponseBody is the type of the "folder" service
// "ListFolders" endpoint HTTP response body for the "invalid token" error.
type ListFoldersInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CommandResponseBody is used to define fields on response body types.
type CommandResponseBody struct {
	// 類型
	Type string `form:"type" json:"type" xml:"type"`
}

// CreatefolderresultdataResponseBody is used to define fields on response body
// types.
type CreatefolderresultdataResponseBody struct {
	// Folder identifier
	FolderID string `form:"folderId" json:"folderId" xml:"folderId"`
	// Base identifier
	BaseID string `form:"baseId" json:"baseId" xml:"baseId"`
	// Parent identifier
	ParentID string `form:"parentId" json:"parentId" xml:"parentId"`
	// Position of the folder
	Position *float64 `form:"position,omitempty" json:"position,omitempty" xml:"position,omitempty"`
	// Time of creation
	CreatedAt string `form:"createdAt" json:"createdAt" xml:"createdAt"`
	// Time of update
	UpdatedAt string `form:"updatedAt" json:"updatedAt" xml:"updatedAt"`
	// Data associated with the folder
	Data *FolderDataResponseBody `form:"data" json:"data" xml:"data"`
	// Type of the folder
	Type string `form:"type" json:"type" xml:"type"`
}

// FolderDataResponseBody is used to define fields on response body types.
type FolderDataResponseBody struct {
	// Folder color
	Color string `form:"color" json:"color" xml:"color"`
	// Folder name
	Name string `form:"name" json:"name" xml:"name"`
}

// FolderRequestBody is used to define fields on request body types.
type FolderRequestBody struct {
	// Folder identifier
	FolderID *string `form:"folderId,omitempty" json:"folderId,omitempty" xml:"folderId,omitempty"`
	// Base identifier
	BaseID *string `form:"baseId,omitempty" json:"baseId,omitempty" xml:"baseId,omitempty"`
	// Parent identifier
	ParentID *string `form:"parentId,omitempty" json:"parentId,omitempty" xml:"parentId,omitempty"`
	// Position of the folder
	Position *float64 `form:"position,omitempty" json:"position,omitempty" xml:"position,omitempty"`
	// Time of creation
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Time of update
	UpdatedAt *string `form:"updatedAt,omitempty" json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// Data associated with the folder
	Data *FolderDataRequestBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// Type of the folder
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// FolderDataRequestBody is used to define fields on request body types.
type FolderDataRequestBody struct {
	// Folder color
	Color *string `form:"color,omitempty" json:"color,omitempty" xml:"color,omitempty"`
	// Folder name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewCreateFolderOKResponseBody builds the HTTP response body from the result
// of the "CreateFolder" endpoint of the "folder" service.
func NewCreateFolderOKResponseBody(res *folder.CreateFolderResult) *CreateFolderOKResponseBody {
	body := &CreateFolderOKResponseBody{
		Timestamp: res.Timestamp,
	}
	if res.Command != nil {
		body.Command = marshalFolderCommandToCommandResponseBody(res.Command)
	}
	if res.Data != nil {
		body.Data = marshalFolderCreatefolderresultdataToCreatefolderresultdataResponseBody(res.Data)
	}
	return body
}

// NewDeleteFolderOKResponseBody builds the HTTP response body from the result
// of the "DeleteFolder" endpoint of the "folder" service.
func NewDeleteFolderOKResponseBody(res *folder.DeleteFolderResult) *DeleteFolderOKResponseBody {
	body := &DeleteFolderOKResponseBody{
		Timestamp: res.Timestamp,
	}
	if res.Command != nil {
		body.Command = marshalFolderCommandToCommandResponseBody(res.Command)
	}
	if res.Data != nil {
		body.Data = marshalFolderCreatefolderresultdataToCreatefolderresultdataResponseBody(res.Data)
	}
	return body
}

// NewUpdateFolderOKResponseBody builds the HTTP response body from the result
// of the "UpdateFolder" endpoint of the "folder" service.
func NewUpdateFolderOKResponseBody(res *folder.UpdateFolderResult) *UpdateFolderOKResponseBody {
	body := &UpdateFolderOKResponseBody{
		Timestamp: res.Timestamp,
	}
	if res.Command != nil {
		body.Command = marshalFolderCommandToCommandResponseBody(res.Command)
	}
	if res.Data != nil {
		body.Data = marshalFolderCreatefolderresultdataToCreatefolderresultdataResponseBody(res.Data)
	}
	return body
}

// NewListFoldersOKResponseBody builds the HTTP response body from the result
// of the "ListFolders" endpoint of the "folder" service.
func NewListFoldersOKResponseBody(res *folder.ListFoldersResult) *ListFoldersOKResponseBody {
	body := &ListFoldersOKResponseBody{
		Timestamp: res.Timestamp,
	}
	if res.Command != nil {
		body.Command = marshalFolderCommandToCommandResponseBody(res.Command)
	}
	if res.Data != nil {
		body.Data = make([]*CreatefolderresultdataResponseBody, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalFolderCreatefolderresultdataToCreatefolderresultdataResponseBody(val)
		}
	} else {
		body.Data = []*CreatefolderresultdataResponseBody{}
	}
	return body
}

// NewCreateFolderInvalidTokenResponseBody builds the HTTP response body from
// the result of the "CreateFolder" endpoint of the "folder" service.
func NewCreateFolderInvalidTokenResponseBody(res *goa.ServiceError) *CreateFolderInvalidTokenResponseBody {
	body := &CreateFolderInvalidTokenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteFolderInvalidTokenResponseBody builds the HTTP response body from
// the result of the "DeleteFolder" endpoint of the "folder" service.
func NewDeleteFolderInvalidTokenResponseBody(res *goa.ServiceError) *DeleteFolderInvalidTokenResponseBody {
	body := &DeleteFolderInvalidTokenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateFolderInvalidTokenResponseBody builds the HTTP response body from
// the result of the "UpdateFolder" endpoint of the "folder" service.
func NewUpdateFolderInvalidTokenResponseBody(res *goa.ServiceError) *UpdateFolderInvalidTokenResponseBody {
	body := &UpdateFolderInvalidTokenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListFoldersInvalidTokenResponseBody builds the HTTP response body from
// the result of the "ListFolders" endpoint of the "folder" service.
func NewListFoldersInvalidTokenResponseBody(res *goa.ServiceError) *ListFoldersInvalidTokenResponseBody {
	body := &ListFoldersInvalidTokenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewReceiveCreateFolderPayload builds a folder service receiveCreateFolder
// endpoint payload.
func NewReceiveCreateFolderPayload(channel string, jwt *string) *folder.ReceiveCreateFolderPayload {
	v := &folder.ReceiveCreateFolderPayload{}
	v.Channel = channel
	v.JWT = jwt

	return v
}

// NewCreateFolderPayload builds a folder service CreateFolder endpoint payload.
func NewCreateFolderPayload(body *CreateFolderRequestBody, jwt *string) *folder.CreateFolderPayload {
	v := &folder.CreateFolderPayload{}
	v.Folder = unmarshalFolderRequestBodyToFolderFolder(body.Folder)
	v.JWT = jwt

	return v
}

// NewReceiveDeleteFolderPayload builds a folder service receiveDeleteFolder
// endpoint payload.
func NewReceiveDeleteFolderPayload(channel string, jwt *string) *folder.ReceiveDeleteFolderPayload {
	v := &folder.ReceiveDeleteFolderPayload{}
	v.Channel = channel
	v.JWT = jwt

	return v
}

// NewDeleteFolderPayload builds a folder service DeleteFolder endpoint payload.
func NewDeleteFolderPayload(folderID string, jwt *string) *folder.DeleteFolderPayload {
	v := &folder.DeleteFolderPayload{}
	v.FolderID = folderID
	v.JWT = jwt

	return v
}

// NewReceiveUpdateFolderPayload builds a folder service receiveUpdateFolder
// endpoint payload.
func NewReceiveUpdateFolderPayload(channel string, jwt *string) *folder.ReceiveUpdateFolderPayload {
	v := &folder.ReceiveUpdateFolderPayload{}
	v.Channel = channel
	v.JWT = jwt

	return v
}

// NewUpdateFolderPayload builds a folder service UpdateFolder endpoint payload.
func NewUpdateFolderPayload(body *UpdateFolderRequestBody, folderID string, jwt *string) *folder.UpdateFolderPayload {
	v := &folder.UpdateFolderPayload{}
	v.Folder = unmarshalFolderRequestBodyToFolderFolder(body.Folder)
	v.FolderID = folderID
	v.JWT = jwt

	return v
}

// NewReceiveListFoldersPayload builds a folder service receiveListFolders
// endpoint payload.
func NewReceiveListFoldersPayload(channel string, jwt *string) *folder.ReceiveListFoldersPayload {
	v := &folder.ReceiveListFoldersPayload{}
	v.Channel = channel
	v.JWT = jwt

	return v
}

// NewListFoldersPayload builds a folder service ListFolders endpoint payload.
func NewListFoldersPayload(jwt *string) *folder.ListFoldersPayload {
	v := &folder.ListFoldersPayload{}
	v.JWT = jwt

	return v
}

// ValidateCreateFolderRequestBody runs the validations defined on
// CreateFolderRequestBody
func ValidateCreateFolderRequestBody(body *CreateFolderRequestBody) (err error) {
	if body.Folder == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("folder", "body"))
	}
	if body.Folder != nil {
		if err2 := ValidateFolderRequestBody(body.Folder); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateFolderRequestBody runs the validations defined on
// UpdateFolderRequestBody
func ValidateUpdateFolderRequestBody(body *UpdateFolderRequestBody) (err error) {
	if body.Folder == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("folder", "body"))
	}
	if body.Folder != nil {
		if err2 := ValidateFolderRequestBody(body.Folder); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateFolderRequestBody runs the validations defined on FolderRequestBody
func ValidateFolderRequestBody(body *FolderRequestBody) (err error) {
	if body.FolderID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("folderId", "body"))
	}
	if body.BaseID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("baseId", "body"))
	}
	if body.ParentID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("parentId", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("createdAt", "body"))
	}
	if body.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updatedAt", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.Data != nil {
		if err2 := ValidateFolderDataRequestBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateFolderDataRequestBody runs the validations defined on
// FolderDataRequestBody
func ValidateFolderDataRequestBody(body *FolderDataRequestBody) (err error) {
	if body.Color == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("color", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}