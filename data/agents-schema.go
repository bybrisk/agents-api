package data

import (
	"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
)

//post request to add agents
type NewAgentsRequest struct{
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

//post response
type AgentsPostSuccess struct {
	BybID string `json:"bybID"`
	Message string `json:"message"`
}

func (d *NewAgentsRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}