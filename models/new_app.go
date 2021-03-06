package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*NewApp Description of a new organization application.

swagger:model NewApp
*/
type NewApp struct {

	/* The URI for the application's homepage
	 */
	ApplicationURI string `json:"application_uri,omitempty"`

	/* The e-mail address of the avatar to use for the application
	 */
	AvatarEmail string `json:"avatar_email,omitempty"`

	/* The human-readable description for the application
	 */
	Description string `json:"description,omitempty"`

	/* The name of the application

	Required: true
	*/
	Name *string `json:"name"`

	/* The URI for the application's OAuth redirect
	 */
	RedirectURI string `json:"redirect_uri,omitempty"`
}

// Validate validates this new app
func (m *NewApp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewApp) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
