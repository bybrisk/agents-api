package data

import (
	"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
)

//post request to add agents
type NewAgentsRequest struct{

	BybID string `json:"-"`
	// The url of the profile pic for this agent
	//
	// required: false
	// max length: 1000
	PicURL string `json: "picurl"`

	// The Name of the agent
	//
	// required: true
	// max length: 1000
	AgentName string `json: "agentName" validate:"required"`

	// Type of Agent
	//
	// required: true
	// example: DELIVERY | CONTRACT
	AgentType string `json:"agentType" validate:"required"`

	// BybID of the business the agent belongs to
	//
	// required: true
	// max length: 1000
	BusinessID string `json: "businessID" validate:"required"`

	// 10 Digit Phone Number of the agent
	//
	// required: true
	// max length: 1000
	PhoneNumber string `json: "phoneNumber" validate:"required"`

	// Permanent Address of the Agent
	//
	// required: true
	// max length: 10000
	Address string `json: "address" validate:"required"`

	// Maximum weight (in Kg) the agent can carry in a single trip
	//
	// required: true
	// max length: 1000
	MaxWeightCapacity int64 `json: "maxWeightCapacity" validate:"required"`

	// Maximum duration (in hr) the agent can work on a single trip
	//
	// required: true
	// max length: 1000
	MaxHourCapacity int64 `json: "maxHourCapacity" validate:"required"`

	// Agent Username/ UserID given by Business (unique to that business)
	//
	// required: true
	// max length: 1000	
	AgentID string `json: "agentID" validate:"required"`
}

//update agents request
type UpdateAgentsRequest struct {

	// BybID of the agent from database
	//
	// required: true
	// max length: 1000
	BybID string `json:"bybID" validate:"required"`

	// The url of the profile pic for this agent
	//
	// required: false
	// max length: 1000
	PicURL string `json:"picurl" validate:"required"`

	// Modified Name of the agent
	//
	// required: true
	// max length: 1000
	AgentName string `json: "agentName" validate:"required"`

	// Type of Agent
	//
	// required: true
	// example: DELIVERY | CONTRACT
	AgentType string `json:"agentType" validate:"required"`

	// 10 Digit Phone Number of the agent
	//
	// required: true
	// max length: 1000
	PhoneNumber string `json: "phoneNumber" validate:"required"`

	// New Permanent Address of the Agent
	//
	// required: true
	// max length: 10000
	Address string `json: "address" validate:"required"`

	// New Maximum weight (in Kg) the agent can carry in a single trip
	//
	// required: true
	// max length: 1000
	MaxWeightCapacity int64 `json: "maxWeightCapacity" validate:"required"`

	// New Maximum duration (in hr) the agent can work on a single trip
	//
	// required: true
	// max length: 1000
	MaxHourCapacity int64 `json: "maxHourCapacity" validate:"required"`

	// Modified Agent Username/ UserID given by Business (unique to that business)
	//
	// required: true
	// max length: 1000
	AgentID string `json: "agentID" validate:"required"`
}

//get all Agents Response struct
type AgentResponseBulk struct {
	// Array of agent details
	//
	Result []AgentSummaryDetail `json:"result"`

	// BybID of business
	//
	BusinessID string `json:"businessid"`

}

type AgentSummaryDetail struct {
	// BybID of each agent
	//
	BybID string `json:"bybid"`

	// Agent Username/ UserID given by Business (unique to that business)
	//
	AgentID string `json: "agentID" validate:"required"`

	// The url of the profile pic for this agent
	//
	PicURL string `json: "picurl"`
	// Name of the agent
	//
	AgentName string `json: "agentName" validate:"required"`

	// Type of Agent
	//
	AgentType string `json:"agentType" validate:"required"`

	// 10 Digit Phone Number of the agent
	//
	PhoneNumber string `json: "phoneNumber" validate:"required"`
}

//Single Agent Get Response
type SingleAgentResponse struct {
	// The url of the profile pic for this agent
	//
	PicURL string `json: "picurl"`

	// The Name of the agent
	//
	AgentName string `json: "agentName"`

	// Type of Agent
	//
	AgentType string `json:"agentType"`

	// BybID of the business the agent belongs to
	//
	BusinessID string `json: "businessID"`

	// 10 Digit Phone Number of the agent
	//
	PhoneNumber string `json: "phoneNumber"`

	// Permanent Address of the Agent
	//
	Address string `json: "address"`

	// Maximum weight (in Kg) the agent can carry in a single trip
	//
	MaxWeightCapacity int64 `json: "maxWeightCapacity"`

	// Maximum duration (in hr) the agent can work on a single trip
	//
	MaxHourCapacity int64 `json: "maxHourCapacity"`

	// Agent Username/ UserID given by Business (unique to that business)
	//	
	AgentID string `json: "agentID"`
}

//post response
type AgentsPostSuccess struct {
	// uuid of the object (agent, business, etc.)
	//
	BybID string `json:"bybID"`
	// Response message of success or failure
	//
	Message string `json:"message"`
}

func (d *NewAgentsRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *UpdateAgentsRequest) ValidateUpdateRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}