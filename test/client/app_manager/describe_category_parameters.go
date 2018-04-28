// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeCategoryParams creates a new DescribeCategoryParams object
// with the default values initialized.
func NewDescribeCategoryParams() *DescribeCategoryParams {
	var ()
	return &DescribeCategoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeCategoryParamsWithTimeout creates a new DescribeCategoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeCategoryParamsWithTimeout(timeout time.Duration) *DescribeCategoryParams {
	var ()
	return &DescribeCategoryParams{

		timeout: timeout,
	}
}

// NewDescribeCategoryParamsWithContext creates a new DescribeCategoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeCategoryParamsWithContext(ctx context.Context) *DescribeCategoryParams {
	var ()
	return &DescribeCategoryParams{

		Context: ctx,
	}
}

// NewDescribeCategoryParamsWithHTTPClient creates a new DescribeCategoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeCategoryParamsWithHTTPClient(client *http.Client) *DescribeCategoryParams {
	var ()
	return &DescribeCategoryParams{
		HTTPClient: client,
	}
}

/*DescribeCategoryParams contains all the parameters to send to the API endpoint
for the describe category operation typically these are written to a http.Request
*/
type DescribeCategoryParams struct {

	/*CategoryID*/
	CategoryID []string
	/*Limit*/
	Limit *int64
	/*Name*/
	Name []string
	/*Offset*/
	Offset *int64
	/*Owner*/
	Owner []string
	/*SearchWord*/
	SearchWord *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe category params
func (o *DescribeCategoryParams) WithTimeout(timeout time.Duration) *DescribeCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe category params
func (o *DescribeCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe category params
func (o *DescribeCategoryParams) WithContext(ctx context.Context) *DescribeCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe category params
func (o *DescribeCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe category params
func (o *DescribeCategoryParams) WithHTTPClient(client *http.Client) *DescribeCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe category params
func (o *DescribeCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryID adds the categoryID to the describe category params
func (o *DescribeCategoryParams) WithCategoryID(categoryID []string) *DescribeCategoryParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the describe category params
func (o *DescribeCategoryParams) SetCategoryID(categoryID []string) {
	o.CategoryID = categoryID
}

// WithLimit adds the limit to the describe category params
func (o *DescribeCategoryParams) WithLimit(limit *int64) *DescribeCategoryParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe category params
func (o *DescribeCategoryParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the describe category params
func (o *DescribeCategoryParams) WithName(name []string) *DescribeCategoryParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the describe category params
func (o *DescribeCategoryParams) SetName(name []string) {
	o.Name = name
}

// WithOffset adds the offset to the describe category params
func (o *DescribeCategoryParams) WithOffset(offset *int64) *DescribeCategoryParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe category params
func (o *DescribeCategoryParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the describe category params
func (o *DescribeCategoryParams) WithOwner(owner []string) *DescribeCategoryParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the describe category params
func (o *DescribeCategoryParams) SetOwner(owner []string) {
	o.Owner = owner
}

// WithSearchWord adds the searchWord to the describe category params
func (o *DescribeCategoryParams) WithSearchWord(searchWord *string) *DescribeCategoryParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe category params
func (o *DescribeCategoryParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesCategoryID := o.CategoryID

	joinedCategoryID := swag.JoinByFormat(valuesCategoryID, "multi")
	// query array param category_id
	if err := r.SetQueryParam("category_id", joinedCategoryID...); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	valuesName := o.Name

	joinedName := swag.JoinByFormat(valuesName, "multi")
	// query array param name
	if err := r.SetQueryParam("name", joinedName...); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOwner := o.Owner

	joinedOwner := swag.JoinByFormat(valuesOwner, "multi")
	// query array param owner
	if err := r.SetQueryParam("owner", joinedOwner...); err != nil {
		return err
	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}