package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/access-context-manager-api/mcp-server/config"
	"github.com/access-context-manager-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Accesscontextmanager_accesspolicies_serviceperimeters_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		parentVal, ok := args["parent"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: parent"), nil
		}
		parent, ok := parentVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: parent"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["pageSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageSize=%v", val))
		}
		if val, ok := args["pageToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageToken=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/%s/servicePerimeters%s", cfg.BaseURL, parent, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.ListServicePerimetersResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateAccesscontextmanager_accesspolicies_serviceperimeters_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_parent_servicePerimeters",
		mcp.WithDescription("Lists all service perimeters for an access policy."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. Resource name for the access policy to list Service Perimeters from. Format: `accessPolicies/{policy_id}`")),
		mcp.WithNumber("pageSize", mcp.Description("Number of Service Perimeters to include in the list. Default 100.")),
		mcp.WithString("pageToken", mcp.Description("Next page token for the next batch of Service Perimeter instances. Defaults to the first page of results.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Accesscontextmanager_accesspolicies_serviceperimeters_listHandler(cfg),
	}
}
