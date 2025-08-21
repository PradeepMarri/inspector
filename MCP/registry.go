package main

import (
	"github.com/modelcontextprotocol-inspector-api/mcp-server/config"
	"github.com/modelcontextprotocol-inspector-api/mcp-server/models"
	tools_health "github.com/modelcontextprotocol-inspector-api/mcp-server/tools/health"
	tools_mcp "github.com/modelcontextprotocol-inspector-api/mcp-server/tools/mcp"
	tools_message "github.com/modelcontextprotocol-inspector-api/mcp-server/tools/message"
	tools_sse "github.com/modelcontextprotocol-inspector-api/mcp-server/tools/sse"
	tools_stdio "github.com/modelcontextprotocol-inspector-api/mcp-server/tools/stdio"
	tools_config "github.com/modelcontextprotocol-inspector-api/mcp-server/tools/config"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_health.CreateGet_healthTool(cfg),
		tools_mcp.CreateGet_mcpTool(cfg),
		tools_mcp.CreatePost_mcpTool(cfg),
		tools_mcp.CreateDelete_mcpTool(cfg),
		tools_message.CreatePost_messageTool(cfg),
		tools_sse.CreateGet_sseTool(cfg),
		tools_stdio.CreateGet_stdioTool(cfg),
		tools_config.CreateGet_configTool(cfg),
	}
}
