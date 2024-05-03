package types

import "time"

type Input struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	DefaultValue string `json:"default_value"`
}

type Output struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Params struct {
	APIKey       string `json:"api_key"`
	Model        string `json:"model"`
	SystemPrompt string `json:"system_prompt"`
	UserPrompt   string `json:"user_prompt"`
}

type Step struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	ModelSlug    string `json:"model_slug"`
	ModelVersion string `json:"model_version"`
	Params       Params `json:"params"`
}

type Flow struct {
	WorkspaceID   string    `json:"workspace_id"`
	UserID        string    `json:"user_id"`
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	ThumbnailURL  string    `json:"thumbnail_url"`
	Status        string    `json:"status"`
	Inputs        []Input   `json:"inputs"`
	Outputs       []Output  `json:"outputs"`
	Steps         []Step    `json:"steps"`
	Verified      bool      `json:"verified"`
	TriggerCount  int       `json:"trigger_count"`
	Popularity    int       `json:"popularity"`
	LastUpdatedBy string    `json:"last_updated_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type GetFlowResponse struct {
	Flow    Flow   `json:"workflow"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type TriggerFlowResponse struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	TriggerID string `json:"trigger_id"`
}

type GetFlowsResponse struct {
	Flows   []Flow `json:"flows"`
	Message string `json:"message"`
	Status  string `json:"status"`
	Total   int    `json:"total"`
}

type StepResult struct {
	StepID string `json:"step_id"`
	Status string `json:"status"`
	Output string `json:"output"`
}

type Execution struct {
	FlowID          string       `json:"flow_id"`
	OrganizationID  string       `json:"organization_id"`
	APIKey          string       `json:"api_key"`
	ExecutionID     string       `json:"execution_id"`
	SourceIPAddress string       `json:"source_ip_address"`
	Parameters      any          `json:"parameters,omitempty"`
	StepResults     []StepResult `json:"step_results"`
	Status          string       `json:"status"`
	Output          string       `json:"output"`
	CreatedAt       time.Time    `json:"created_at"`
	StartedAt       time.Time    `json:"started_at"`
	EndedAt         time.Time    `json:"ended_at"`
	UpdatedAt       string       `json:"updated_at"`
	DeletedAt       string       `json:"deleted_at"`
}

type GetExecutionsResp struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Total      int         `json:"total"`
	Executions []Execution `json:"executions"`
}
