// Package classification of Agents API
//
// Documentation for Agents API
//
//	Schemes: https
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package handlers
import "github.com/bybrisk/agents-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// Success message on adding a single Agent
// swagger:response agentsPostResponse
type accountPostResponseWrapper struct {
	// Success message on newly created Agents
	// in: body
	Body data.AgentsPostSuccess
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters createAgents
type createAgentsParamsWrapper struct {
	// Agents data structure to Create.
	// Note: the id field is ignored by create operations
	// in: body
	// required: true
	Body data.NewAgentsRequest
}

// swagger:parameters updateAgents
type updateAgentsParamsWrapper struct {
	// Data structure to update an existing agent.
	// Note: You need to send all the details required or they will be overwritten
	// in: body
	// required: true
	Body data.UpdateAgentsRequest
}

// Summary of all the agents of a business account
// swagger:response getAllAgentsDetails
type agentsGetAllResponseWrapper struct {
	// Summarised details of all the agents of a business account
	// in: body
	Body data.AgentResponseBulk
}

// Details of a single agents of a particular business account
// swagger:response getSingleAgentsDetails
type agentsGetSingleResponseWrapper struct {
	// Complete details of a single agents of a business account
	// in: body
	Body data.SingleAgentResponse
}