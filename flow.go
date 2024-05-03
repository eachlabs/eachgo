package each

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/eachlabs/eachgo/types"
)

// GetFlows returns a list of flows.
func (e *Client) GetFlows(ctx context.Context) ([]types.Flow, error) {
	req, err := e.newRequest(ctx, BackendServiceFlow, http.MethodGet, "/", nil)
	if err != nil {
		return nil, err
	}

	var resp types.GetFlowsResponse
	if err := e.doRequest(ctx, req, &resp); err != nil {
		return nil, err
	}

	if resp.Status != "success" {
		return nil, errors.New(resp.Message)
	}

	return resp.Flows, nil
}

// GetFlow returns a flow by ID.
func (e *Client) GetFlow(ctx context.Context, id string) (*types.Flow, error) {
	req, err := e.newRequest(ctx, BackendServiceFlow, http.MethodGet, fmt.Sprintf("/%s", id), nil)
	if err != nil {
		return nil, err
	}

	var resp types.GetFlowResponse
	if err := e.doRequest(ctx, req, &resp); err != nil {
		return nil, err
	}

	return &resp.Flow, nil
}

// TriggerFlow triggers a flow by ID.
func (e *Client) TriggerFlow(ctx context.Context, id string, input map[string]interface{}) (string, error) {

	ib, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	req, err := e.newRequest(ctx, BackendServiceFlow, http.MethodPost, fmt.Sprintf("/%s/trigger", id), io.Reader(bytes.NewReader(ib)))
	if err != nil {
		return "", err
	}

	var resp types.TriggerFlowResponse
	if err := e.doRequest(ctx, req, &resp); err != nil {
		return "", err
	}

	return resp.TriggerID, nil
}

// GetExecution returns all executions by Flow ID.
func (e *Client) GetExecutions(ctx context.Context, flowId string) ([]types.Execution, error) {
	req, err := e.newRequest(ctx, BackendServiceFlow, http.MethodGet, fmt.Sprintf("/%s/executions", flowId), nil)
	if err != nil {
		return nil, err
	}

	var resp types.GetExecutionsResp
	if err := e.doRequest(ctx, req, &resp); err != nil {
		return nil, err
	}

	if resp.Status != "success" {
		return nil, fmt.Errorf("failed to get executions: %s", resp.Message)
	}

	return resp.Executions, nil
}

// GetExecution returns an execution by ID.
func (e *Client) GetExecution(ctx context.Context, flowId, executionId string) (*types.Execution, error) {
	req, err := e.newRequest(ctx, BackendServiceFlow, http.MethodGet, fmt.Sprintf("/%s/executions/%s", flowId, executionId), nil)
	if err != nil {
		return nil, err
	}

	var execution types.Execution
	if err := e.doRequest(ctx, req, &execution); err != nil {
		return nil, err
	}

	return &execution, nil
}
