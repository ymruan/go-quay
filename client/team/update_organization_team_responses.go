package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type UpdateOrganizationTeamReader struct {
	formats strfmt.Registry
}

func (o *UpdateOrganizationTeamReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result UpdateOrganizationTeamOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result UpdateOrganizationTeamBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("updateOrganizationTeamBadRequest", &result, response.Code())

	case 401:
		var result UpdateOrganizationTeamUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("updateOrganizationTeamUnauthorized", &result, response.Code())

	case 403:
		var result UpdateOrganizationTeamForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("updateOrganizationTeamForbidden", &result, response.Code())

	case 404:
		var result UpdateOrganizationTeamNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("updateOrganizationTeamNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Successful invocation
*/
type UpdateOrganizationTeamOK struct {
}

func (o *UpdateOrganizationTeamOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type UpdateOrganizationTeamBadRequest struct {
}

func (o *UpdateOrganizationTeamBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Session required
*/
type UpdateOrganizationTeamUnauthorized struct {
}

func (o *UpdateOrganizationTeamUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type UpdateOrganizationTeamForbidden struct {
}

func (o *UpdateOrganizationTeamForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type UpdateOrganizationTeamNotFound struct {
}

func (o *UpdateOrganizationTeamNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}