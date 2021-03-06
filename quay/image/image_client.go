package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new image API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for image API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetImage Get the information available for the specified image.
*/
func (a *Client) GetImage(params *GetImageParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getImage",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/image/{image_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetImageReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetImageOK), nil
}

/*
ListRepositoryImages List the images for the specified repository.
*/
func (a *Client) ListRepositoryImages(params *ListRepositoryImagesParams, authInfo runtime.ClientAuthInfoWriter) (*ListRepositoryImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRepositoryImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRepositoryImages",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/image/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRepositoryImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRepositoryImagesOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
