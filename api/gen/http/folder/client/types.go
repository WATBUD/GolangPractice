// Code generated by goa v3.16.2, DO NOT EDIT.
//
// folder HTTP client types
//
// Command:
// $ goa gen mai.today/api/design

package client

import (
	goa "goa.design/goa/v3/pkg"
	folder "mai.today/api/gen/folder"
)

// CreateFolderRequestBody is the type of the "folder" service "CreateFolder"
// endpoint HTTP request body.
type CreateFolderRequestBody struct {
	Folder *FolderRequestBody `form:"folder" json:"folder" xml:"folder"`
}

// UpdateFolderRequestBody is the type of the "folder" service "UpdateFolder"
// endpoint HTTP request body.
type UpdateFolderRequestBody struct {
	// Folder data to update
	Folder *FolderRequestBody `form:"folder" json:"folder" xml:"folder"`
}

// CreateFolderOKResponseBody is the type of the "folder" service
// "CreateFolder" endpoint HTTP response body.
type CreateFolderOKResponseBody struct {
	Command *CommandResponseBody `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// 時間戳記
	Timestamp *int64 `form:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 資料
	Data *CreatefolderresultdataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// DeleteFolderOKResponseBody is the type of the "folder" service
// "DeleteFolder" endpoint HTTP response body.
type DeleteFolderOKResponseBody struct {
	Command *CommandResponseBody `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// 時間戳記
	Timestamp *int64 `form:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 資料
	Data *CreatefolderresultdataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// UpdateFolderOKResponseBody is the type of the "folder" service
// "UpdateFolder" endpoint HTTP response body.
type UpdateFolderOKResponseBody struct {
	Command *CommandResponseBody `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// 時間戳記
	Timestamp *int64 `form:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 資料
	Data *CreatefolderresultdataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// ListFoldersOKResponseBody is the type of the "folder" service "ListFolders"
// endpoint HTTP response body.
type ListFoldersOKResponseBody struct {
	Command *CommandResponseBody `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// 時間戳記
	Timestamp *int64 `form:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 資料
	Data []*CreatefolderresultdataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// CreateFolderInvalidTokenResponseBody is the type of the "folder" service
// "CreateFolder" endpoint HTTP response body for the "invalid token" error.
type CreateFolderInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// DeleteFolderInvalidTokenResponseBody is the type of the "folder" service
// "DeleteFolder" endpoint HTTP response body for the "invalid token" error.
type DeleteFolderInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateFolderInvalidTokenResponseBody is the type of the "folder" service
// "UpdateFolder" endpoint HTTP response body for the "invalid token" error.
type UpdateFolderInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ListFoldersInvalidTokenResponseBody is the type of the "folder" service
// "ListFolders" endpoint HTTP response body for the "invalid token" error.
type ListFoldersInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ReceiveCreateFolderOKResponseBody is used to define fields on response body
// types.
type ReceiveCreateFolderOKResponseBody struct {
	Command *CommandResponseBody `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// 時間戳記
	Timestamp *int64 `form:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 資料
	Data *CreatefolderresultdataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// CommandResponseBody is used to define fields on response body types.
type CommandResponseBody struct {
	// 類型
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// CreatefolderresultdataResponseBody is used to define fields on response body
// types.
type CreatefolderresultdataResponseBody struct {
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
	Data *FolderDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// Type of the folder
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// FolderDataResponseBody is used to define fields on response body types.
type FolderDataResponseBody struct {
	// Folder color
	Color *string `form:"color,omitempty" json:"color,omitempty" xml:"color,omitempty"`
	// Folder name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// FolderRequestBody is used to define fields on request body types.
type FolderRequestBody struct {
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
	Data *FolderDataRequestBody `form:"data" json:"data" xml:"data"`
	// Type of the folder
	Type string `form:"type" json:"type" xml:"type"`
}

// FolderDataRequestBody is used to define fields on request body types.
type FolderDataRequestBody struct {
	// Folder color
	Color string `form:"color" json:"color" xml:"color"`
	// Folder name
	Name string `form:"name" json:"name" xml:"name"`
}

// ReceiveDeleteFolderOKResponseBody is used to define fields on response body
// types.
type ReceiveDeleteFolderOKResponseBody struct {
	Command *CommandResponseBody `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// 時間戳記
	Timestamp *int64 `form:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 資料
	Data *CreatefolderresultdataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// ReceiveUpdateFolderOKResponseBody is used to define fields on response body
// types.
type ReceiveUpdateFolderOKResponseBody struct {
	Command *CommandResponseBody `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// 時間戳記
	Timestamp *int64 `form:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 資料
	Data *CreatefolderresultdataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// ReceiveListFoldersOKResponseBody is used to define fields on response body
// types.
type ReceiveListFoldersOKResponseBody struct {
	Command *CommandResponseBody `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// 時間戳記
	Timestamp *int64 `form:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 資料
	Data []*CreatefolderresultdataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// NewCreateFolderRequestBody builds the HTTP request body from the payload of
// the "CreateFolder" endpoint of the "folder" service.
func NewCreateFolderRequestBody(p *folder.CreateFolderPayload) *CreateFolderRequestBody {
	body := &CreateFolderRequestBody{}
	if p.Folder != nil {
		body.Folder = marshalFolderFolderToFolderRequestBody(p.Folder)
	}
	return body
}

// NewUpdateFolderRequestBody builds the HTTP request body from the payload of
// the "UpdateFolder" endpoint of the "folder" service.
func NewUpdateFolderRequestBody(p *folder.UpdateFolderPayload) *UpdateFolderRequestBody {
	body := &UpdateFolderRequestBody{}
	if p.Folder != nil {
		body.Folder = marshalFolderFolderToFolderRequestBody(p.Folder)
	}
	return body
}

// NewReceiveCreateFolderCreateFolderResultSwitchingProtocols builds a "folder"
// service "receiveCreateFolder" endpoint result from a HTTP
// "SwitchingProtocols" response.
func NewReceiveCreateFolderCreateFolderResultSwitchingProtocols() *folder.CreateFolderResult {
	v := &folder.CreateFolderResult{}

	return v
}

// NewCreateFolderResultOK builds a "folder" service "CreateFolder" endpoint
// result from a HTTP "OK" response.
func NewCreateFolderResultOK(body *CreateFolderOKResponseBody) *folder.CreateFolderResult {
	v := &folder.CreateFolderResult{
		Timestamp: *body.Timestamp,
	}
	v.Command = unmarshalCommandResponseBodyToFolderCommand(body.Command)
	v.Data = unmarshalCreatefolderresultdataResponseBodyToFolderCreatefolderresultdata(body.Data)

	return v
}

// NewCreateFolderInvalidToken builds a folder service CreateFolder endpoint
// invalid token error.
func NewCreateFolderInvalidToken(body *CreateFolderInvalidTokenResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewReceiveDeleteFolderDeleteFolderResultSwitchingProtocols builds a "folder"
// service "receiveDeleteFolder" endpoint result from a HTTP
// "SwitchingProtocols" response.
func NewReceiveDeleteFolderDeleteFolderResultSwitchingProtocols() *folder.DeleteFolderResult {
	v := &folder.DeleteFolderResult{}

	return v
}

// NewDeleteFolderResultOK builds a "folder" service "DeleteFolder" endpoint
// result from a HTTP "OK" response.
func NewDeleteFolderResultOK(body *DeleteFolderOKResponseBody) *folder.DeleteFolderResult {
	v := &folder.DeleteFolderResult{
		Timestamp: *body.Timestamp,
	}
	v.Command = unmarshalCommandResponseBodyToFolderCommand(body.Command)
	v.Data = unmarshalCreatefolderresultdataResponseBodyToFolderCreatefolderresultdata(body.Data)

	return v
}

// NewDeleteFolderInvalidToken builds a folder service DeleteFolder endpoint
// invalid token error.
func NewDeleteFolderInvalidToken(body *DeleteFolderInvalidTokenResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewReceiveUpdateFolderUpdateFolderResultSwitchingProtocols builds a "folder"
// service "receiveUpdateFolder" endpoint result from a HTTP
// "SwitchingProtocols" response.
func NewReceiveUpdateFolderUpdateFolderResultSwitchingProtocols() *folder.UpdateFolderResult {
	v := &folder.UpdateFolderResult{}

	return v
}

// NewUpdateFolderResultOK builds a "folder" service "UpdateFolder" endpoint
// result from a HTTP "OK" response.
func NewUpdateFolderResultOK(body *UpdateFolderOKResponseBody) *folder.UpdateFolderResult {
	v := &folder.UpdateFolderResult{
		Timestamp: *body.Timestamp,
	}
	v.Command = unmarshalCommandResponseBodyToFolderCommand(body.Command)
	v.Data = unmarshalCreatefolderresultdataResponseBodyToFolderCreatefolderresultdata(body.Data)

	return v
}

// NewUpdateFolderInvalidToken builds a folder service UpdateFolder endpoint
// invalid token error.
func NewUpdateFolderInvalidToken(body *UpdateFolderInvalidTokenResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewReceiveListFoldersListFoldersResultSwitchingProtocols builds a "folder"
// service "receiveListFolders" endpoint result from a HTTP
// "SwitchingProtocols" response.
func NewReceiveListFoldersListFoldersResultSwitchingProtocols() *folder.ListFoldersResult {
	v := &folder.ListFoldersResult{}

	return v
}

// NewListFoldersResultOK builds a "folder" service "ListFolders" endpoint
// result from a HTTP "OK" response.
func NewListFoldersResultOK(body *ListFoldersOKResponseBody) *folder.ListFoldersResult {
	v := &folder.ListFoldersResult{
		Timestamp: *body.Timestamp,
	}
	v.Command = unmarshalCommandResponseBodyToFolderCommand(body.Command)
	v.Data = make([]*folder.Createfolderresultdata, len(body.Data))
	for i, val := range body.Data {
		v.Data[i] = unmarshalCreatefolderresultdataResponseBodyToFolderCreatefolderresultdata(val)
	}

	return v
}

// NewListFoldersInvalidToken builds a folder service ListFolders endpoint
// invalid token error.
func NewListFoldersInvalidToken(body *ListFoldersInvalidTokenResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateCreateFolderOKResponseBody runs the validations defined on
// CreateFolderOKResponseBody
func ValidateCreateFolderOKResponseBody(body *CreateFolderOKResponseBody) (err error) {
	if body.Command == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("command", "body"))
	}
	if body.Timestamp == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timestamp", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Command != nil {
		if err2 := ValidateCommandResponseBody(body.Command); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Timestamp != nil {
		if *body.Timestamp < 1.719188007e+18 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.timestamp", *body.Timestamp, 1.719188007e+18, true))
		}
	}
	if body.Data != nil {
		if err2 := ValidateCreatefolderresultdataResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateDeleteFolderOKResponseBody runs the validations defined on
// DeleteFolderOKResponseBody
func ValidateDeleteFolderOKResponseBody(body *DeleteFolderOKResponseBody) (err error) {
	if body.Command == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("command", "body"))
	}
	if body.Timestamp == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timestamp", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Command != nil {
		if err2 := ValidateCommandResponseBody(body.Command); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Timestamp != nil {
		if *body.Timestamp < 1.719188007e+18 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.timestamp", *body.Timestamp, 1.719188007e+18, true))
		}
	}
	if body.Data != nil {
		if err2 := ValidateCreatefolderresultdataResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateFolderOKResponseBody runs the validations defined on
// UpdateFolderOKResponseBody
func ValidateUpdateFolderOKResponseBody(body *UpdateFolderOKResponseBody) (err error) {
	if body.Command == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("command", "body"))
	}
	if body.Timestamp == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timestamp", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Command != nil {
		if err2 := ValidateCommandResponseBody(body.Command); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Timestamp != nil {
		if *body.Timestamp < 1.719188007e+18 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.timestamp", *body.Timestamp, 1.719188007e+18, true))
		}
	}
	if body.Data != nil {
		if err2 := ValidateCreatefolderresultdataResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateListFoldersOKResponseBody runs the validations defined on
// ListFoldersOKResponseBody
func ValidateListFoldersOKResponseBody(body *ListFoldersOKResponseBody) (err error) {
	if body.Command == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("command", "body"))
	}
	if body.Timestamp == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timestamp", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Command != nil {
		if err2 := ValidateCommandResponseBody(body.Command); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Timestamp != nil {
		if *body.Timestamp < 1.719188007e+18 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.timestamp", *body.Timestamp, 1.719188007e+18, true))
		}
	}
	for _, e := range body.Data {
		if e != nil {
			if err2 := ValidateCreatefolderresultdataResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateCreateFolderInvalidTokenResponseBody runs the validations defined on
// CreateFolder_invalid token_Response_Body
func ValidateCreateFolderInvalidTokenResponseBody(body *CreateFolderInvalidTokenResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateDeleteFolderInvalidTokenResponseBody runs the validations defined on
// DeleteFolder_invalid token_Response_Body
func ValidateDeleteFolderInvalidTokenResponseBody(body *DeleteFolderInvalidTokenResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateFolderInvalidTokenResponseBody runs the validations defined on
// UpdateFolder_invalid token_Response_Body
func ValidateUpdateFolderInvalidTokenResponseBody(body *UpdateFolderInvalidTokenResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateListFoldersInvalidTokenResponseBody runs the validations defined on
// ListFolders_invalid token_Response_Body
func ValidateListFoldersInvalidTokenResponseBody(body *ListFoldersInvalidTokenResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateReceiveCreateFolderOKResponseBody runs the validations defined on
// ReceiveCreateFolderOKResponseBody
func ValidateReceiveCreateFolderOKResponseBody(body *ReceiveCreateFolderOKResponseBody) (err error) {
	if body.Command == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("command", "body"))
	}
	if body.Timestamp == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timestamp", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Command != nil {
		if err2 := ValidateCommandResponseBody(body.Command); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Timestamp != nil {
		if *body.Timestamp < 1.719188007e+18 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.timestamp", *body.Timestamp, 1.719188007e+18, true))
		}
	}
	if body.Data != nil {
		if err2 := ValidateCreatefolderresultdataResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateCommandResponseBody runs the validations defined on
// CommandResponseBody
func ValidateCommandResponseBody(body *CommandResponseBody) (err error) {
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.Type != nil {
		if !(*body.Type == "sync" || *body.Type == "createBase" || *body.Type == "deleteBase" || *body.Type == "reorderBaseNavState" || *body.Type == "updateBaseInfo") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.type", *body.Type, []any{"sync", "createBase", "deleteBase", "reorderBaseNavState", "updateBaseInfo"}))
		}
	}
	return
}

// ValidateCreatefolderresultdataResponseBody runs the validations defined on
// CreatefolderresultdataResponseBody
func ValidateCreatefolderresultdataResponseBody(body *CreatefolderresultdataResponseBody) (err error) {
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
		if err2 := ValidateFolderDataResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateFolderDataResponseBody runs the validations defined on
// FolderDataResponseBody
func ValidateFolderDataResponseBody(body *FolderDataResponseBody) (err error) {
	if body.Color == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("color", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

// ValidateFolderRequestBody runs the validations defined on FolderRequestBody
func ValidateFolderRequestBody(body *FolderRequestBody) (err error) {
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	return
}

// ValidateReceiveDeleteFolderOKResponseBody runs the validations defined on
// ReceiveDeleteFolderOKResponseBody
func ValidateReceiveDeleteFolderOKResponseBody(body *ReceiveDeleteFolderOKResponseBody) (err error) {
	if body.Command == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("command", "body"))
	}
	if body.Timestamp == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timestamp", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Command != nil {
		if err2 := ValidateCommandResponseBody(body.Command); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Timestamp != nil {
		if *body.Timestamp < 1.719188007e+18 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.timestamp", *body.Timestamp, 1.719188007e+18, true))
		}
	}
	if body.Data != nil {
		if err2 := ValidateCreatefolderresultdataResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateReceiveUpdateFolderOKResponseBody runs the validations defined on
// ReceiveUpdateFolderOKResponseBody
func ValidateReceiveUpdateFolderOKResponseBody(body *ReceiveUpdateFolderOKResponseBody) (err error) {
	if body.Command == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("command", "body"))
	}
	if body.Timestamp == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timestamp", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Command != nil {
		if err2 := ValidateCommandResponseBody(body.Command); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Timestamp != nil {
		if *body.Timestamp < 1.719188007e+18 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.timestamp", *body.Timestamp, 1.719188007e+18, true))
		}
	}
	if body.Data != nil {
		if err2 := ValidateCreatefolderresultdataResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateReceiveListFoldersOKResponseBody runs the validations defined on
// ReceiveListFoldersOKResponseBody
func ValidateReceiveListFoldersOKResponseBody(body *ReceiveListFoldersOKResponseBody) (err error) {
	if body.Command == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("command", "body"))
	}
	if body.Timestamp == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timestamp", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Command != nil {
		if err2 := ValidateCommandResponseBody(body.Command); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Timestamp != nil {
		if *body.Timestamp < 1.719188007e+18 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.timestamp", *body.Timestamp, 1.719188007e+18, true))
		}
	}
	for _, e := range body.Data {
		if e != nil {
			if err2 := ValidateCreatefolderresultdataResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
