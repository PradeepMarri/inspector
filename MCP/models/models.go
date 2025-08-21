package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Message string `json:"message,omitempty"`
	Code int `json:"code,omitempty"`
}
