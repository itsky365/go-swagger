package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/swag"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewAddCommentToTaskParams creates a new AddCommentToTaskParams object
// with the default values initialized.
func NewAddCommentToTaskParams() AddCommentToTaskParams {
	var ()
	return AddCommentToTaskParams{}
}

// AddCommentToTaskParams contains all the bound params for the add comment to task operation
// typically these are obtained from a http.Request
//
// swagger:parameters addCommentToTask
type AddCommentToTaskParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The comment to add
	  In: body
	*/
	Body AddCommentToTaskBody
	/*The id of the item
	  Required: true
	  In: path
	*/
	ID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddCommentToTaskParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if httpkit.HasBody(r) {
		defer r.Body.Close()
		var body AddCommentToTaskBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {

			if len(res) == 0 {
				o.Body = body
			}
		}

	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddCommentToTaskParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	return nil
}
