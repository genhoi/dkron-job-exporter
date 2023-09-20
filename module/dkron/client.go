package dkron

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	api    string
	client *http.Client
}

func CreateClient(api string) *Client {
	return &Client{
		api: api,
		client: &http.Client{
			Timeout: time.Duration(3) * time.Second,
		},
	}
}

func (c *Client) GetJobs() ([]Job, error) {
	u, err := url.Parse(c.api)
	if err != nil {
		return nil, err
	}
	u.Path = "/v1/jobs"

	q := u.Query()
	q.Add("_start", "0")
	q.Add("_end", "1000")
	q.Add("_sort", "id")
	q.Add("_order", "ASC")

	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	jobs := make([]Job, 0)
	err = json.NewDecoder(resp.Body).Decode(&jobs)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (c *Client) ListExecutionsByJob(name string) ([]Execution, error) {
	u, err := url.Parse(c.api)
	if err != nil {
		return nil, err
	}

	u.Path = fmt.Sprintf("/v1/jobs/%s/executions", name)

	q := u.Query()
	q.Add("output_size_limit", "0")

	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	executions := make([]Execution, 0)
	err = json.NewDecoder(resp.Body).Decode(&executions)
	if err != nil {
		return nil, err
	}

	return executions, nil
}
