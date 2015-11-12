package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
BuildTriggerActivateRequest build trigger activate request

swagger:model BuildTriggerActivateRequest
*/
type BuildTriggerActivateRequest struct {

	/* Arbitrary json.

	Required: true
	*/
	Config map[string]interface{} `json:"config"`

	/* The name of the robot that will be used to pull images.
	 */
	PullRobot string `json:"pull_robot,omitempty"`
}

// Validate validates this build trigger activate request
func (m *BuildTriggerActivateRequest) Validate(formats strfmt.Registry) error {

	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}

func (m *BuildTriggerActivateRequest) validateConfig(formats strfmt.Registry) error {

	return nil
}