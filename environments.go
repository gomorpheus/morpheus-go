// Morpheus API types and Client methods for Environments
package morpheus

import (
	"fmt"
)

// globals
var (
	EnvironmentsPath = "/api/environments"
)

// Environment structures for use in request and response payloads

type Environment struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Visibility  string `json:"visibility"`
}

type ListEnvironmentsResult struct {
	Environments *[]Environment `json:"environments"`
	Meta         *MetaResult    `json:"meta"`
}

type GetEnvironmentResult struct {
	Environment *Environment `json:"environment"`
}

type CreateEnvironmentResult struct {
	Success     bool              `json:"success"`
	Message     string            `json:"msg"`
	Errors      map[string]string `json:"errors"`
	Environment *Environment      `json:"environment"`
}

type UpdateEnvironmentResult struct {
	CreateEnvironmentResult
}

type DeleteEnvironmentResult struct {
	DeleteResult
}

// Client request methods

func (client *Client) ListEnvironments(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        EnvironmentsPath,
		QueryParams: req.QueryParams,
		Result:      &ListEnvironmentsResult{},
	})
}

func (client *Client) GetEnvironment(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", EnvironmentsPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetEnvironmentResult{},
	})
}

func (client *Client) CreateEnvironment(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        EnvironmentsPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateEnvironmentResult{},
	})
}

func (client *Client) UpdateEnvironment(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", EnvironmentsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateEnvironmentResult{},
	})
}

func (client *Client) DeleteEnvironment(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", EnvironmentsPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteEnvironmentResult{},
	})
}

// helper functions
func (client *Client) FindEnvironmentByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListEnvironments(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListEnvironmentsResult)
	environmentsCount := len(*listResult.Environments)
	if environmentsCount != 1 {
		return resp, fmt.Errorf("found %d Environments for %v", environmentsCount, name)
	}
	firstRecord := (*listResult.Environments)[0]
	environmentID := firstRecord.ID
	return client.GetEnvironment(environmentID, &Request{})
}
