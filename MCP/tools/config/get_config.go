package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/modelcontextprotocol-inspector-api/mcp-server/config"
	"github.com/modelcontextprotocol-inspector-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_configHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		url := fmt.Sprintf("%s/config", cfg.BaseURL)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["origin"]; ok {
			req.Header.Set("origin", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-mcp-proxy-auth"]; ok {
			req.Header.Set("x-mcp-proxy-auth", fmt.Sprintf("%v", val))
		}

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
		var result map[string]interface{}
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

func CreateGet_configTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_config",
		mcp.WithDescription(""),
		mcp.WithString("origin", mcp.Description("")),
		mcp.WithString("x-mcp-proxy-auth", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_configHandler(cfg),
	}
}
