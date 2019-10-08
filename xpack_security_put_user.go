// Copyright 2012-2019 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/olivere/elastic/v7/uritemplates"
)

// XPackSecurityPutUserService retrieves a user by its name.
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.0/security-api-put-user.html.
type XPackSecurityPutUserService struct {
	client   *Client
	pretty   bool
	username string
	body     interface{}
}

// NewXPackSecurityPutUserService creates a new XPackSecurityPutUserService.
func NewXPackSecurityPutUserService(client *Client) *XPackSecurityPutUserService {
	return &XPackSecurityPutUserService{
		client: client,
	}
}

// Name is the name of the user to create.
func (s *XPackSecurityPutUserService) Name(username string) *XPackSecurityPutUserService {
	s.username = username
	return s
}

// Pretty indicates that the JSON response be indented and human readable.
func (s *XPackSecurityPutUserService) Pretty(pretty bool) *XPackSecurityPutUserService {
	s.pretty = pretty
	return s
}

// Body specifies the user. Use a string or a type that will get serialized as JSON.
func (s *XPackSecurityPutUserService) Body(body interface{}) *XPackSecurityPutUserService {
	s.body = body
	return s
}

// buildURL builds the URL for the operation.
func (s *XPackSecurityPutUserService) buildURL() (string, url.Values, error) {
	// Build URL
	path, err := uritemplates.Expand("/_security/user/{username}", map[string]string{
		"username": s.username,
	})
	if err != nil {
		return "", url.Values{}, err
	}

	// Add query string parameters
	params := url.Values{}
	if s.pretty {
		params.Set("pretty", "true")
	}
	return path, params, nil
}

// Validate checks if the operation is valid.
func (s *XPackSecurityPutUserService) Validate() error {
	var invalid []string
	if s.username == "" {
		invalid = append(invalid, "Name")
	}
	if s.body == nil {
		invalid = append(invalid, "Body")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Do executes the operation.
func (s *XPackSecurityPutUserService) Do(ctx context.Context) (*XPackSecurityPutUserResponse, error) {
	// Check pre-conditions
	if err := s.Validate(); err != nil {
		return nil, err
	}

	// Get URL for request
	path, params, err := s.buildURL()
	if err != nil {
		return nil, err
	}

	// Get HTTP response
	res, err := s.client.PerformRequest(ctx, PerformRequestOptions{
		Method: "PUT",
		Path:   path,
		Params: params,
		Body:   s.body,
	})
	if err != nil {
		return nil, err
	}

	// Return operation response
	ret := new(XPackSecurityPutUserResponse)
	if err := json.Unmarshal(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// XPackSecurityPutUserResponse is the response of XPackSecurityPutUserService.Do.
type XPackSecurityPutUserResponse struct {
	User XPackSecurityPutUser
}

// XPackSecurityPutUser is the response containing the creation information
type XPackSecurityPutUser struct {
	Created bool `json:"created"`
}
