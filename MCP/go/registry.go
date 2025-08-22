package main

import (
	"github.com/shotstack/mcp-server/config"
	"github.com/shotstack/mcp-server/models"
	tools_serve "github.com/shotstack/mcp-server/tools/serve"
	tools_edit "github.com/shotstack/mcp-server/tools/edit"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_serve.CreateDeleteassetTool(cfg),
		tools_serve.CreateGetassetTool(cfg),
		tools_edit.CreatePostrenderTool(cfg),
		tools_edit.CreateGetrenderTool(cfg),
		tools_serve.CreateGetassetbyrenderidTool(cfg),
	}
}
