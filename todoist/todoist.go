package todoist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	todoistAPIBaseURL  = "https://api.todoist.com/rest/v2"
	todositAPITasksURL = todoistAPIBaseURL + "/tasks"
)

// Task represents a task in Todoist.
// Plase refer to the Todoist API documentation for more information:
// https://developer.todoist.com/rest/v2/#tasks
type Task struct {
	CreatorID    string    `json:"creator_id,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	AssigneeID   string    `json:"assignee_id,omitempty"`
	AssignerID   string    `json:"assigner_id,omitempty"`
	CommentCount int       `json:"comment_count,omitempty"`
	IsCompleted  bool      `json:"is_completed,omitempty"`
	Content      string    `json:"content"`
	Description  string    `json:"description,omitempty"`
	Due          Due       `json:"due,omitempty"`
	Duration     *Duration `json:"duration,omitempty"`
	ID           string    `json:"id,omitempty"`
	Labels       []string  `json:"labels,omitempty"`
	Order        int       `json:"order,omitempty"`
	Priority     int       `json:"priority,omitempty"`
	ProjectID    string    `json:"project_id,omitempty"`
	SectionID    string    `json:"section_id,omitempty"`
	ParentID     string    `json:"parent_id,omitempty"`
	URL          string    `json:"url,omitempty"`
}

// Due represents the due date object of a task.
// Plase refer to the Todoist API documentation for more information:
// https://developer.todoist.com/rest/v2/#tasks
type Due struct {
	Date        string `json:"date"`
	IsRecurring bool   `json:"is_recurring"`
	Datetime    string `json:"datetime"`
	String      string `json:"string"`
	Timezone    string `json:"timezone"`
}

// Duration represents the duration object of a task.
// Plase refer to the Todoist API documentation for more information:
// https://developer.todoist.com/rest/v2/#tasks
type Duration struct {
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}

// Client provides a client to the Todoist API.
// Currently only supports listing and creating tasks.
type Client struct {
	token      string
	httpClient http.Client
}

// NewClient returns a new todoist client.
// Requires a Todoist API token.
func NewClient(token string, opts ...ClientOption) *Client {
	httpClient := &http.Client{}
	client := &Client{
		token:      token,
		httpClient: *httpClient,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// ClientOption represents functional options to configure the client.
type ClientOption func(*Client)

// WithHTTPClient configures the client to use the given http.Client.
func WithHTTPTransport(transport http.RoundTripper) func(*Client) {
	return func(c *Client) {
		c.httpClient.Transport = transport
	}
}

// ListTasks returns a list of tasks from the Todoist service.
func (c *Client) ListTasks() ([]Task, error) {
	resp, err := c.makeRequest(http.MethodGet, todositAPITasksURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error making get task request: %w", err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error making get task request: %s", err)
	}

	var tasks []Task
	json.Unmarshal(body, &tasks)

	return tasks, nil
}

// CreateTask creates a task in the Todoist service.
func (c *Client) CreateTask(task Task) error {
	jsonValue, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("error marshalling task: %w", err)
	}

	resp, err := c.makeRequest(http.MethodPost, todositAPITasksURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		return fmt.Errorf("error making create task request: %w", err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error making create task request: %s", body)
	}

	return nil
}

func (c *Client) makeRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, todositAPITasksURL, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	return resp, nil
}
