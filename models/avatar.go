package models

import "github.com/go-openapi/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Avatar avatar

swagger:model Avatar
*/
type Avatar struct {

	/* color
	 */
	Color string `json:"color,omitempty"`

	/* hash
	 */
	Hash string `json:"hash,omitempty"`

	/* kind
	 */
	Kind string `json:"kind,omitempty"`

	/* name
	 */
	Name string `json:"name,omitempty"`
}

// Validate validates this avatar
func (m *Avatar) Validate(formats strfmt.Registry) error {
	return nil
}
