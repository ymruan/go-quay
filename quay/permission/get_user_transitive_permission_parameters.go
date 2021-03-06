package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUserTransitivePermissionParams creates a new GetUserTransitivePermissionParams object
// with the default values initialized.
func NewGetUserTransitivePermissionParams() *GetUserTransitivePermissionParams {
	var ()
	return &GetUserTransitivePermissionParams{}
}

/*GetUserTransitivePermissionParams contains all the parameters to send to the API endpoint
for the get user transitive permission operation typically these are written to a http.Request
*/
type GetUserTransitivePermissionParams struct {

	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Username
	  The username of the user to which the permissions apply

	*/
	Username string
}

// WithRepository adds the repository to the get user transitive permission params
func (o *GetUserTransitivePermissionParams) WithRepository(Repository string) *GetUserTransitivePermissionParams {
	o.Repository = Repository
	return o
}

// WithUsername adds the username to the get user transitive permission params
func (o *GetUserTransitivePermissionParams) WithUsername(Username string) *GetUserTransitivePermissionParams {
	o.Username = Username
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserTransitivePermissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
