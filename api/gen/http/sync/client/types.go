// Code generated by goa v3.16.2, DO NOT EDIT.
//
// Sync HTTP client types
//
// Command:
// $ goa gen mai.today/api/design

package client

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
	sync "mai.today/api/gen/sync"
)

// SyncRequestBody is the type of the "Sync" service "sync" endpoint HTTP
// request body.
type SyncRequestBody struct {
	// 上次同步的時間戳記
	LastTimestamp *int64 `form:"lastTimestamp,omitempty" json:"lastTimestamp,omitempty" xml:"lastTimestamp,omitempty"`
}

// SyncOKResponseBody is the type of the "Sync" service "sync" endpoint HTTP
// response body.
type SyncOKResponseBody struct {
	Command *CommandResponseBody `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// 時間戳記
	Timestamp *int64 `form:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 資料
	Data *SyncPayloadDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// SyncInvalidTokenResponseBody is the type of the "Sync" service "sync"
// endpoint HTTP response body for the "invalid token" error.
type SyncInvalidTokenResponseBody struct {
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

// ReceiveSyncOKResponseBody is used to define fields on response body types.
type ReceiveSyncOKResponseBody struct {
	Command *CommandResponseBody `form:"command,omitempty" json:"command,omitempty" xml:"command,omitempty"`
	// 時間戳記
	Timestamp *int64 `form:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 資料
	Data *SyncPayloadDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// CommandResponseBody is used to define fields on response body types.
type CommandResponseBody struct {
	// 類型
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// SyncPayloadDataResponseBody is used to define fields on response body types.
type SyncPayloadDataResponseBody struct {
	BaseInfos     []*BaseInfoResponseBody     `form:"baseInfos,omitempty" json:"baseInfos,omitempty" xml:"baseInfos,omitempty"`
	BaseNavStates []*BaseNavStateResponseBody `form:"baseNavStates,omitempty" json:"baseNavStates,omitempty" xml:"baseNavStates,omitempty"`
}

// BaseInfoResponseBody is used to define fields on response body types.
type BaseInfoResponseBody struct {
	// 識別碼
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 基地識別碼
	BaseID *string `form:"baseId,omitempty" json:"baseId,omitempty" xml:"baseId,omitempty"`
	// 主題颜色
	Color *string `form:"color,omitempty" json:"color,omitempty" xml:"color,omitempty"`
	// 標識圖片的 URL
	Logo *string `form:"logo,omitempty" json:"logo,omitempty" xml:"logo,omitempty"`
	// 名稱
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 刪除時間
	DeletedAt *string `form:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
}

// BaseNavStateResponseBody is used to define fields on response body types.
type BaseNavStateResponseBody struct {
	// 識別碼
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 基地識別碼
	BaseID *string `form:"baseId,omitempty" json:"baseId,omitempty" xml:"baseId,omitempty"`
	// 順序
	Index *int `form:"index,omitempty" json:"index,omitempty" xml:"index,omitempty"`
	// 刪除時間
	DeletedAt *string `form:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
}

// NewSyncRequestBody builds the HTTP request body from the payload of the
// "sync" endpoint of the "Sync" service.
func NewSyncRequestBody(p *sync.SyncPayload) *SyncRequestBody {
	body := &SyncRequestBody{
		LastTimestamp: p.LastTimestamp,
	}
	return body
}

// NewReceiveSyncSyncResultSwitchingProtocols builds a "Sync" service
// "receiveSync" endpoint result from a HTTP "SwitchingProtocols" response.
func NewReceiveSyncSyncResultSwitchingProtocols() *sync.SyncResult {
	v := &sync.SyncResult{}

	return v
}

// NewSyncResultOK builds a "Sync" service "sync" endpoint result from a HTTP
// "OK" response.
func NewSyncResultOK(body *SyncOKResponseBody) *sync.SyncResult {
	v := &sync.SyncResult{
		Timestamp: *body.Timestamp,
	}
	v.Command = unmarshalCommandResponseBodyToSyncCommand(body.Command)
	v.Data = unmarshalSyncPayloadDataResponseBodyToSyncSyncPayloadData(body.Data)

	return v
}

// NewSyncInvalidToken builds a Sync service sync endpoint invalid token error.
func NewSyncInvalidToken(body *SyncInvalidTokenResponseBody) *goa.ServiceError {
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

// ValidateSyncOKResponseBody runs the validations defined on SyncOKResponseBody
func ValidateSyncOKResponseBody(body *SyncOKResponseBody) (err error) {
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
		if err2 := ValidateSyncPayloadDataResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSyncInvalidTokenResponseBody runs the validations defined on
// sync_invalid token_response_body
func ValidateSyncInvalidTokenResponseBody(body *SyncInvalidTokenResponseBody) (err error) {
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

// ValidateReceiveSyncOKResponseBody runs the validations defined on
// ReceiveSyncOKResponseBody
func ValidateReceiveSyncOKResponseBody(body *ReceiveSyncOKResponseBody) (err error) {
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
		if err2 := ValidateSyncPayloadDataResponseBody(body.Data); err2 != nil {
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

// ValidateSyncPayloadDataResponseBody runs the validations defined on
// SyncPayloadDataResponseBody
func ValidateSyncPayloadDataResponseBody(body *SyncPayloadDataResponseBody) (err error) {
	for _, e := range body.BaseInfos {
		if e != nil {
			if err2 := ValidateBaseInfoResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range body.BaseNavStates {
		if e != nil {
			if err2 := ValidateBaseNavStateResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateBaseInfoResponseBody runs the validations defined on
// BaseInfoResponseBody
func ValidateBaseInfoResponseBody(body *BaseInfoResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.BaseID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("baseId", "body"))
	}
	if body.Color == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("color", "body"))
	}
	if body.Logo == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("logo", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID != nil {
		if utf8.RuneCountInString(*body.ID) < 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.id", *body.ID, utf8.RuneCountInString(*body.ID), 24, true))
		}
	}
	if body.ID != nil {
		if utf8.RuneCountInString(*body.ID) > 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.id", *body.ID, utf8.RuneCountInString(*body.ID), 24, false))
		}
	}
	if body.BaseID != nil {
		if utf8.RuneCountInString(*body.BaseID) < 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.baseId", *body.BaseID, utf8.RuneCountInString(*body.BaseID), 24, true))
		}
	}
	if body.BaseID != nil {
		if utf8.RuneCountInString(*body.BaseID) > 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.baseId", *body.BaseID, utf8.RuneCountInString(*body.BaseID), 24, false))
		}
	}
	if body.Color != nil {
		if utf8.RuneCountInString(*body.Color) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.color", *body.Color, utf8.RuneCountInString(*body.Color), 1, true))
		}
	}
	if body.Color != nil {
		if utf8.RuneCountInString(*body.Color) > 16 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.color", *body.Color, utf8.RuneCountInString(*body.Color), 16, false))
		}
	}
	if body.Logo != nil {
		if utf8.RuneCountInString(*body.Logo) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.logo", *body.Logo, utf8.RuneCountInString(*body.Logo), 1, true))
		}
	}
	if body.Logo != nil {
		if utf8.RuneCountInString(*body.Logo) > 1024 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.logo", *body.Logo, utf8.RuneCountInString(*body.Logo), 1024, false))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 1, true))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 128, false))
		}
	}
	if body.DeletedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.deleted_at", *body.DeletedAt, goa.FormatDateTime))
	}
	return
}

// ValidateBaseNavStateResponseBody runs the validations defined on
// BaseNavStateResponseBody
func ValidateBaseNavStateResponseBody(body *BaseNavStateResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.BaseID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("baseId", "body"))
	}
	if body.Index == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("index", "body"))
	}
	if body.ID != nil {
		if utf8.RuneCountInString(*body.ID) < 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.id", *body.ID, utf8.RuneCountInString(*body.ID), 24, true))
		}
	}
	if body.ID != nil {
		if utf8.RuneCountInString(*body.ID) > 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.id", *body.ID, utf8.RuneCountInString(*body.ID), 24, false))
		}
	}
	if body.BaseID != nil {
		if utf8.RuneCountInString(*body.BaseID) < 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.baseId", *body.BaseID, utf8.RuneCountInString(*body.BaseID), 24, true))
		}
	}
	if body.BaseID != nil {
		if utf8.RuneCountInString(*body.BaseID) > 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.baseId", *body.BaseID, utf8.RuneCountInString(*body.BaseID), 24, false))
		}
	}
	if body.Index != nil {
		if *body.Index < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.index", *body.Index, 1, true))
		}
	}
	if body.DeletedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.deleted_at", *body.DeletedAt, goa.FormatDateTime))
	}
	return
}
