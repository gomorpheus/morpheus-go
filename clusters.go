// Morpheus API types and Client methods for Clusters
package morpheus

import (
	"fmt"
)

// globals

var (
	ClustersPath = "/api/clusters"
)

// types

type Cluster struct {
	ID          int64                  `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Type        string                 `json:"type"`
	Layout      string                 `json:"layout"`
	Group       map[string]interface{} `json:"group"`
	Cloud       map[string]interface{} `json:"cloud"`
	Server      Server                 `json:"server"`
	Status      string                 `json:"status"`
}

type Server struct {
	Config            map[string]interface{}    `json:"config"`
	Name              string                    `json:"name"`
	HostName          string                    `json:"hostname"`
	Plan              map[string]interface{}    `json:"plan"`
	Volumes           *[]map[string]interface{} `json:"volumes"`
	NetworkInterfaces *[]map[string]interface{} `json:"networkInterfaces"`
	Visibility        string                    `json:"visibility"`
	NodeCount         int64                     `json:"nodeCount"`
}

type ListClustersResult struct {
	Clusters *[]Cluster  `json:"clusteres"`
	Meta     *MetaResult `json:"meta"`
}

type GetClusterResult struct {
	Cluster *Cluster `json:"clustere"`
}

type CreateClusterResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	Cluster *Cluster          `json:"clustere"`
}

type UpdateClusterResult struct {
	CreateClusterResult
}

type DeleteClusterResult struct {
	DeleteResult
}

// API endpoints

func (client *Client) ListClusters(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        ClustersPath,
		QueryParams: req.QueryParams,
		Result:      &ListClustersResult{},
	})
}

func (client *Client) GetCluster(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%d", ClustersPath, id),
		QueryParams: req.QueryParams,
		Result:      &GetClusterResult{},
	})
}

func (client *Client) CreateCluster(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "POST",
		Path:        ClustersPath,
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &CreateClusterResult{},
	})
}

func (client *Client) UpdateCluster(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "PUT",
		Path:        fmt.Sprintf("%s/%d", ClustersPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &UpdateClusterResult{},
	})
}

func (client *Client) DeleteCluster(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%d", ClustersPath, id),
		QueryParams: req.QueryParams,
		Body:        req.Body,
		Result:      &DeleteClusterResult{},
	})
}

// helper functions

func (client *Client) FindClusterByName(name string) (*Response, error) {
	// Find by name, then get by ID
	resp, err := client.ListClusters(&Request{
		QueryParams: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return resp, err
	}
	listResult := resp.Result.(*ListClustersResult)
	clustereCount := len(*listResult.Clusters)
	if clustereCount != 1 {
		return resp, fmt.Errorf("Found %d Clusters for %v", clustereCount, name)
	}
	firstRecord := (*listResult.Clusters)[0]
	clustereId := firstRecord.ID
	return client.GetCluster(clustereId, &Request{})
}
