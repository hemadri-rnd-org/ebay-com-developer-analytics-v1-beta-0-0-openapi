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
	Category string `json:"category,omitempty"` // Identifies the type of erro.
	Inputrefids []string `json:"inputRefIds,omitempty"` // An array of request elements most closely associated to the error.
	Longmessage string `json:"longMessage,omitempty"` // A more detailed explanation of the error.
	Outputrefids []string `json:"outputRefIds,omitempty"` // An array of request elements most closely associated to the error.
	Subdomain string `json:"subdomain,omitempty"` // Further helps indicate which subsystem the error is coming from. System subcategories include: Initialization, Serialization, Security, Monitoring, Rate Limiting, etc.
	Domain string `json:"domain,omitempty"` // Name for the primary system where the error occurred. This is relevant for application errors.
	Errorid int `json:"errorId,omitempty"` // A unique number to identify the error.
	Message string `json:"message,omitempty"` // Information on how to correct the problem, in the end user's terms and language where applicable.
	Parameters []ErrorParameter `json:"parameters,omitempty"` // An array of name/value pairs that describe details the error condition. These are useful when multiple errors are returned.
}

// ErrorParameter represents the ErrorParameter schema from the OpenAPI specification
type ErrorParameter struct {
	Name string `json:"name,omitempty"` // The object of the error.
	Value string `json:"value,omitempty"` // The value of the object.
}

// Rate represents the Rate schema from the OpenAPI specification
type Rate struct {
	Reset string `json:"reset,omitempty"` // The data and time the time window and accumulated calls for this resource reset. <br><br>When the <b>reset</b> time is reached, the <b>remaining</b> value is reset to the value of <b>limit</b>, and this <b>reset</b> value is reset to the current time plus the number of seconds defined by the <b>timeWindow</b> value. <br><br>The time stamp is formatted as an <a href="http://www.iso.org/iso/home/standards/iso8601.htm " target="_blank">ISO 8601</a> string, which is based on the 24-hour Universal Coordinated Time (UTC) clock. <br><br><b>Format:</b> <code>[YYYY]-[MM]-[DD]T[hh]:[mm]:[ss].[sss]Z</code> <br><b>Example:</b> <code>2018-08-04T07:09:00.000Z</code>
	Timewindow int `json:"timeWindow,omitempty"` // A period of time, expressed in seconds. The call quota for a resource is applied to the period of time defined by the value of this field.
	Limit int `json:"limit,omitempty"` // The maximum number of requests that can be made to this resource during a set time period. The length of time to which the limit is applied is defined by the associated <b>timeWindow</b> value. <br><br>This value is often referred to as the "call quota" for the resource.
	Remaining int `json:"remaining,omitempty"` // The remaining number of requests that can be made to this resource before the associated time window resets.
}

// RateLimit represents the RateLimit schema from the OpenAPI specification
type RateLimit struct {
	Apicontext string `json:"apiContext,omitempty"` // The context of the API for which rate-limit data is returned. For example <code>buy</code>, <code>sell</code>, <code>commerce</code>, <code>developer</code> or <code>tradingapi</code>.
	Apiname string `json:"apiName,omitempty"` // The name of the API for which rate-limit data is returned. For example <code>browse</code> for the Buy API, <code>inventory</code> for the Sell API, <code>taxonomy</code> for the Commerce API, or <code>tradingapi</code> for Trading API.
	Apiversion string `json:"apiVersion,omitempty"` // The version of the API for which rate-limit data is returned. For example <code>v1</code> or <code>v2</code>.
	Resources []Resource `json:"resources,omitempty"` // A list of the methods for which rate-limit data is returned. For example <code>item</code> for the Feed API, <code>getOrder</code> for the Fulfillment API, <code>getProduct</code> for the Catalog API, <code>AddItems</code> for the Trading API.
}

// RateLimitsResponse represents the RateLimitsResponse schema from the OpenAPI specification
type RateLimitsResponse struct {
	Ratelimits []RateLimit `json:"rateLimits,omitempty"` // The rate-limit data for the specified APIs. The rate-limit data is returned for all the methods in the specified APIs and data pertains to the current time window.
}

// Resource represents the Resource schema from the OpenAPI specification
type Resource struct {
	Name string `json:"name,omitempty"` // The name of the resource (an API or an API method) to which the rate-limit data applies.
	Rates []Rate `json:"rates,omitempty"` // A list of rate-limit data, where each list element represents the rate-limit data for a specific resource.
}
