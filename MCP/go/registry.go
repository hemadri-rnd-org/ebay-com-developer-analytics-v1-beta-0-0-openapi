package main

import (
	"github.com/analytics-api/mcp-server/config"
	"github.com/analytics-api/mcp-server/models"
	tools_user_rate_limit "github.com/analytics-api/mcp-server/tools/user_rate_limit"
	tools_rate_limit "github.com/analytics-api/mcp-server/tools/rate_limit"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_user_rate_limit.CreateGetuserratelimitsTool(cfg),
		tools_rate_limit.CreateGetratelimitsTool(cfg),
	}
}
