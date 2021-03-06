// Code generated by go-swagger; DO NOT EDIT.

package repo_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewDeleteRepoParams creates a new DeleteRepoParams object
// with the default values initialized.
func NewDeleteRepoParams() *DeleteRepoParams {
	var ()
	return &DeleteRepoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRepoParamsWithTimeout creates a new DeleteRepoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRepoParamsWithTimeout(timeout time.Duration) *DeleteRepoParams {
	var ()
	return &DeleteRepoParams{

		timeout: timeout,
	}
}

// NewDeleteRepoParamsWithContext creates a new DeleteRepoParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRepoParamsWithContext(ctx context.Context) *DeleteRepoParams {
	var ()
	return &DeleteRepoParams{

		Context: ctx,
	}
}

// NewDeleteRepoParamsWithHTTPClient creates a new DeleteRepoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRepoParamsWithHTTPClient(client *http.Client) *DeleteRepoParams {
	var ()
	return &DeleteRepoParams{
		HTTPClient: client,
	}
}

/*DeleteRepoParams contains all the parameters to send to the API endpoint
for the delete repo operation typically these are written to a http.Request
*/
type DeleteRepoParams struct {

	/*Body*/
	Body *models.OpenpitrixDeleteRepoRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete repo params
func (o *DeleteRepoParams) WithTimeout(timeout time.Duration) *DeleteRepoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete repo params
func (o *DeleteRepoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete repo params
func (o *DeleteRepoParams) WithContext(ctx context.Context) *DeleteRepoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete repo params
func (o *DeleteRepoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete repo params
func (o *DeleteRepoParams) WithHTTPClient(client *http.Client) *DeleteRepoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete repo params
func (o *DeleteRepoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete repo params
func (o *DeleteRepoParams) WithBody(body *models.OpenpitrixDeleteRepoRequest) *DeleteRepoParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete repo params
func (o *DeleteRepoParams) SetBody(body *models.OpenpitrixDeleteRepoRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRepoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
