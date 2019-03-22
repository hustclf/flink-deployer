package flink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
)

type terminateRequest struct {
	Mode string `json:"mode"`
}

// Cancel terminates a running job specified by job ID
func (c FlinkRestClient) Terminate(jobID string, mode string) error {
	terminateRequest := terminateRequest{
		Mode: mode,
	}

	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(terminateRequest)

	req, err := retryablehttp.NewRequest("PATCH", c.constructURL(fmt.Sprintf("jobs/%v?mode=%v", jobID, mode)), nil)

	if err != nil {
		return err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 202 {
		return fmt.Errorf("Unexpected response status %v", res.StatusCode)
	}

	return nil
}
