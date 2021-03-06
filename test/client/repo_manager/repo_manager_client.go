// Code generated by go-swagger; DO NOT EDIT.

package repo_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new repo manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for repo manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateRepo creates repo
*/
func (a *Client) CreateRepo(params *CreateRepoParams) (*CreateRepoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateRepo",
		Method:             "POST",
		PathPattern:        "/v1/repos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRepoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRepoOK), nil

}

/*
DeleteRepo deletes repo
*/
func (a *Client) DeleteRepo(params *DeleteRepoParams) (*DeleteRepoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRepo",
		Method:             "DELETE",
		PathPattern:        "/v1/repos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRepoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRepoOK), nil

}

/*
DescribeRepos describes repos with filter
*/
func (a *Client) DescribeRepos(params *DescribeReposParams) (*DescribeReposOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeReposParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeRepos",
		Method:             "GET",
		PathPattern:        "/v1/repos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeReposReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeReposOK), nil

}

/*
ModifyRepo modifies repo
*/
func (a *Client) ModifyRepo(params *ModifyRepoParams) (*ModifyRepoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyRepo",
		Method:             "PATCH",
		PathPattern:        "/v1/repos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyRepoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyRepoOK), nil

}

/*
ValidateRepo validates repo
*/
func (a *Client) ValidateRepo(params *ValidateRepoParams) (*ValidateRepoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ValidateRepo",
		Method:             "GET",
		PathPattern:        "/v1/repos/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValidateRepoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ValidateRepoOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
