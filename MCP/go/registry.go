package main

import (
	"github.com/access-context-manager-api/mcp-server/config"
	"github.com/access-context-manager-api/mcp-server/models"
	tools_accesspolicies "github.com/access-context-manager-api/mcp-server/tools/accesspolicies"
	tools_services "github.com/access-context-manager-api/mcp-server/tools/services"
	tools_organizations "github.com/access-context-manager-api/mcp-server/tools/organizations"
	tools_operations "github.com/access-context-manager-api/mcp-server/tools/operations"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_setiampolicyTool(cfg),
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_accesslevels_listTool(cfg),
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_accesslevels_createTool(cfg),
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_serviceperimeters_replaceallTool(cfg),
		tools_services.CreateAccesscontextmanager_services_getTool(cfg),
		tools_organizations.CreateAccesscontextmanager_organizations_gcpuseraccessbindings_listTool(cfg),
		tools_organizations.CreateAccesscontextmanager_organizations_gcpuseraccessbindings_createTool(cfg),
		tools_organizations.CreateAccesscontextmanager_organizations_gcpuseraccessbindings_getTool(cfg),
		tools_organizations.CreateAccesscontextmanager_organizations_gcpuseraccessbindings_patchTool(cfg),
		tools_organizations.CreateAccesscontextmanager_organizations_gcpuseraccessbindings_deleteTool(cfg),
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_getiampolicyTool(cfg),
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_serviceperimeters_listTool(cfg),
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_serviceperimeters_createTool(cfg),
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_serviceperimeters_testiampermissionsTool(cfg),
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_authorizedorgsdescs_listTool(cfg),
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_authorizedorgsdescs_createTool(cfg),
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_listTool(cfg),
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_createTool(cfg),
		tools_services.CreateAccesscontextmanager_services_listTool(cfg),
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_accesslevels_replaceallTool(cfg),
		tools_operations.CreateAccesscontextmanager_operations_cancelTool(cfg),
		tools_accesspolicies.CreateAccesscontextmanager_accesspolicies_serviceperimeters_commitTool(cfg),
	}
}
