// Code generated by goa v3.16.2, DO NOT EDIT.
//
// Base HTTP client CLI support package
//
// Command:
// $ goa gen mai.today/api/design

package client

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
	base "mai.today/api/gen/base"
)

// BuildReceiveCreateBasePayload builds the payload for the Base
// receiveCreateBase endpoint from CLI flags.
func BuildReceiveCreateBasePayload(baseReceiveCreateBaseChannel string, baseReceiveCreateBaseJWT string) (*base.ReceiveCreateBasePayload, error) {
	var err error
	var channel string
	{
		channel = baseReceiveCreateBaseChannel
		if utf8.RuneCountInString(channel) < 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("channel", channel, utf8.RuneCountInString(channel), 24, true))
		}
		if utf8.RuneCountInString(channel) > 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("channel", channel, utf8.RuneCountInString(channel), 24, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if baseReceiveCreateBaseJWT != "" {
			jwt = &baseReceiveCreateBaseJWT
		}
	}
	v := &base.ReceiveCreateBasePayload{}
	v.Channel = channel
	v.JWT = jwt

	return v, nil
}

// BuildCreateBasePayload builds the payload for the Base CreateBase endpoint
// from CLI flags.
func BuildCreateBasePayload(baseCreateBaseBody string, baseCreateBaseJWT string) (*base.CreateBasePayload, error) {
	var err error
	var body CreateBaseRequestBody
	{
		err = json.Unmarshal([]byte(baseCreateBaseBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"color\": \"#FF5733\",\n      \"index\": 0,\n      \"logo\": \"http://example.com/logo.png\",\n      \"name\": \"基地 A\"\n   }'")
		}
		if utf8.RuneCountInString(body.Color) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.color", body.Color, utf8.RuneCountInString(body.Color), 1, true))
		}
		if utf8.RuneCountInString(body.Color) > 16 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.color", body.Color, utf8.RuneCountInString(body.Color), 16, false))
		}
		if body.Index < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.index", body.Index, 1, true))
		}
		if utf8.RuneCountInString(body.Logo) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.logo", body.Logo, utf8.RuneCountInString(body.Logo), 1, true))
		}
		if utf8.RuneCountInString(body.Logo) > 1024 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.logo", body.Logo, utf8.RuneCountInString(body.Logo), 1024, false))
		}
		if utf8.RuneCountInString(body.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 1, true))
		}
		if utf8.RuneCountInString(body.Name) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 128, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if baseCreateBaseJWT != "" {
			jwt = &baseCreateBaseJWT
		}
	}
	v := &base.CreateBasePayload{
		Color: body.Color,
		Index: body.Index,
		Logo:  body.Logo,
		Name:  body.Name,
	}
	{
		var zero int
		if v.Index == zero {
			v.Index = 1
		}
	}
	v.JWT = jwt

	return v, nil
}

// BuildReceiveDeleteBasePayload builds the payload for the Base
// receiveDeleteBase endpoint from CLI flags.
func BuildReceiveDeleteBasePayload(baseReceiveDeleteBaseChannel string, baseReceiveDeleteBaseJWT string) (*base.ReceiveDeleteBasePayload, error) {
	var err error
	var channel string
	{
		channel = baseReceiveDeleteBaseChannel
		if utf8.RuneCountInString(channel) < 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("channel", channel, utf8.RuneCountInString(channel), 24, true))
		}
		if utf8.RuneCountInString(channel) > 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("channel", channel, utf8.RuneCountInString(channel), 24, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if baseReceiveDeleteBaseJWT != "" {
			jwt = &baseReceiveDeleteBaseJWT
		}
	}
	v := &base.ReceiveDeleteBasePayload{}
	v.Channel = channel
	v.JWT = jwt

	return v, nil
}

// BuildDeleteBasePayload builds the payload for the Base DeleteBase endpoint
// from CLI flags.
func BuildDeleteBasePayload(baseDeleteBaseID string, baseDeleteBaseJWT string) (*base.DeleteBasePayload, error) {
	var err error
	var id string
	{
		id = baseDeleteBaseID
		if utf8.RuneCountInString(id) < 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("id", id, utf8.RuneCountInString(id), 24, true))
		}
		if utf8.RuneCountInString(id) > 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("id", id, utf8.RuneCountInString(id), 24, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if baseDeleteBaseJWT != "" {
			jwt = &baseDeleteBaseJWT
		}
	}
	v := &base.DeleteBasePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildReceiveUpdateBaseInfoPayload builds the payload for the Base
// receiveUpdateBaseInfo endpoint from CLI flags.
func BuildReceiveUpdateBaseInfoPayload(baseReceiveUpdateBaseInfoChannel string, baseReceiveUpdateBaseInfoJWT string) (*base.ReceiveUpdateBaseInfoPayload, error) {
	var err error
	var channel string
	{
		channel = baseReceiveUpdateBaseInfoChannel
		if utf8.RuneCountInString(channel) < 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("channel", channel, utf8.RuneCountInString(channel), 24, true))
		}
		if utf8.RuneCountInString(channel) > 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("channel", channel, utf8.RuneCountInString(channel), 24, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if baseReceiveUpdateBaseInfoJWT != "" {
			jwt = &baseReceiveUpdateBaseInfoJWT
		}
	}
	v := &base.ReceiveUpdateBaseInfoPayload{}
	v.Channel = channel
	v.JWT = jwt

	return v, nil
}

// BuildUpdateBaseInfoPayload builds the payload for the Base UpdateBaseInfo
// endpoint from CLI flags.
func BuildUpdateBaseInfoPayload(baseUpdateBaseInfoBody string, baseUpdateBaseInfoID string, baseUpdateBaseInfoJWT string) (*base.UpdateBaseInfoPayload, error) {
	var err error
	var body UpdateBaseInfoRequestBody
	{
		err = json.Unmarshal([]byte(baseUpdateBaseInfoBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"color\": \"#FF5733\",\n      \"logo\": \"http://example.com/logo.png\",\n      \"name\": \"基地 A\"\n   }'")
		}
		if utf8.RuneCountInString(body.Color) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.color", body.Color, utf8.RuneCountInString(body.Color), 1, true))
		}
		if utf8.RuneCountInString(body.Color) > 16 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.color", body.Color, utf8.RuneCountInString(body.Color), 16, false))
		}
		if utf8.RuneCountInString(body.Logo) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.logo", body.Logo, utf8.RuneCountInString(body.Logo), 1, true))
		}
		if utf8.RuneCountInString(body.Logo) > 1024 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.logo", body.Logo, utf8.RuneCountInString(body.Logo), 1024, false))
		}
		if utf8.RuneCountInString(body.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 1, true))
		}
		if utf8.RuneCountInString(body.Name) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 128, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = baseUpdateBaseInfoID
		if utf8.RuneCountInString(id) < 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("id", id, utf8.RuneCountInString(id), 24, true))
		}
		if utf8.RuneCountInString(id) > 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("id", id, utf8.RuneCountInString(id), 24, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if baseUpdateBaseInfoJWT != "" {
			jwt = &baseUpdateBaseInfoJWT
		}
	}
	v := &base.UpdateBaseInfoPayload{
		Color: body.Color,
		Logo:  body.Logo,
		Name:  body.Name,
	}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildReceiveReorderBaseNavStatesPayload builds the payload for the Base
// receiveReorderBaseNavStates endpoint from CLI flags.
func BuildReceiveReorderBaseNavStatesPayload(baseReceiveReorderBaseNavStatesChannel string, baseReceiveReorderBaseNavStatesJWT string) (*base.ReceiveReorderBaseNavStatesPayload, error) {
	var err error
	var channel string
	{
		channel = baseReceiveReorderBaseNavStatesChannel
		if utf8.RuneCountInString(channel) < 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("channel", channel, utf8.RuneCountInString(channel), 24, true))
		}
		if utf8.RuneCountInString(channel) > 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("channel", channel, utf8.RuneCountInString(channel), 24, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if baseReceiveReorderBaseNavStatesJWT != "" {
			jwt = &baseReceiveReorderBaseNavStatesJWT
		}
	}
	v := &base.ReceiveReorderBaseNavStatesPayload{}
	v.Channel = channel
	v.JWT = jwt

	return v, nil
}

// BuildReorderBaseNavStatesPayload builds the payload for the Base
// ReorderBaseNavStates endpoint from CLI flags.
func BuildReorderBaseNavStatesPayload(baseReorderBaseNavStatesBody string, baseReorderBaseNavStatesJWT string) (*base.ReorderBaseNavStatesPayload, error) {
	var err error
	var body ReorderBaseNavStatesRequestBody
	{
		err = json.Unmarshal([]byte(baseReorderBaseNavStatesBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"669071b8a650a662b82285ca\",\n      \"newIndex\": 0\n   }'")
		}
		if utf8.RuneCountInString(body.ID) < 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.id", body.ID, utf8.RuneCountInString(body.ID), 24, true))
		}
		if utf8.RuneCountInString(body.ID) > 24 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.id", body.ID, utf8.RuneCountInString(body.ID), 24, false))
		}
		if body.NewIndex < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.newIndex", body.NewIndex, 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if baseReorderBaseNavStatesJWT != "" {
			jwt = &baseReorderBaseNavStatesJWT
		}
	}
	v := &base.ReorderBaseNavStatesPayload{
		ID:       body.ID,
		NewIndex: body.NewIndex,
	}
	v.JWT = jwt

	return v, nil
}