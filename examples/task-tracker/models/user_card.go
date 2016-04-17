package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/swag"
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*UserCard A minimal representation of a user.

This representation of a user is mainly meant for inclusion in other models, or for list views.


swagger:model UserCard
*/
type UserCard struct {

	/* When true this user is an admin.

	Only employees of the owning company can be admins.
	Admins are like project owners but have access to all the projects in the application.
	There aren't many admins, and it's only used for extremly critical issues with the application.


	Read Only: true
	*/
	Admin *bool `json:"admin,omitempty"`

	/* The amount of karma this user has available.

	In this application users get a cerain amount of karma alotted.
	This karma can be donated to other users to show appreciation, or it can be used
	by a user to vote on issues.
	Once an issue is closed or rejected, the user gets his karma back.


	Read Only: true
	Maximum: < 1000
	*/
	AvailableKarma float64 `json:"availableKarma,omitempty"`

	/* A unique identifier for a user.

	This id is automatically generated on the server when a user is created.


	Required: true
	Read Only: true
	*/
	ID int64 `json:"id"`

	/* The screen name for the user.

	This is used for vanity type urls as well as login credentials.


	Required: true
	Max Length: 255
	Min Length: 3
	Pattern: \w[\w_-]+
	*/
	ScreenName *string `json:"screenName"`
}

// Validate validates this user card
func (m *UserCard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableKarma(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScreenName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserCard) validateAvailableKarma(formats strfmt.Registry) error {

	if swag.IsZero(m.AvailableKarma) { // not required
		return nil
	}

	if err := validate.Maximum("availableKarma", "body", float64(m.AvailableKarma), 1000, true); err != nil {
		return err
	}

	return nil
}

func (m *UserCard) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *UserCard) validateScreenName(formats strfmt.Registry) error {

	if err := validate.Required("screenName", "body", m.ScreenName); err != nil {
		return err
	}

	if err := validate.MinLength("screenName", "body", string(*m.ScreenName), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("screenName", "body", string(*m.ScreenName), 255); err != nil {
		return err
	}

	if err := validate.Pattern("screenName", "body", string(*m.ScreenName), `\w[\w_-]+`); err != nil {
		return err
	}

	return nil
}
