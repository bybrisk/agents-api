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

	// 12 Digit Aadhar Card number of the Agent
	//
	// required: true
	// max length: 1000	
	AadharNumber string `json: "aadharNumber" validate:"required"`

	// Driving licence number of the agent
	//
	// required: true
	// max length: 1000	
	DrivingLicenceNumber string `json: "drivingLicenceNumber" validate:"required"`

	// Type of Vehicle Agent is using
	//
	// required: true
	// max length: 1000	
	// example: TWO WHEELER | VAN | MINI TRUCK | TRUCK
	TypeOfVehicle string `json: "typeOfVehicle" validate:"required"`

	// PAN Card number of the Agent 
	//
	// required: false
	// max length: 1000	
	PANCardNumber string `json: "panCardNumber"`
}

type AutoNewAgentsRequest struct{

	BybID string `json:"-"`

	// The Name of the agent
	//
	// required: false
	// max length: 1000
	AgentName string `json: "agentName"`

	// BybID of the business these agent belongs to
	//
	// required: true
	// max length: 1000
	BusinessID string `json: "businessID" validate:"required"`

	// Number of agents required
	//
	// required: true
	// max length: 1000
	AgentNum int64 `json: "agentNum" validate:"required"`

	// 10 Digit Phone Number of the agent
	//
	// required: false
	// max length: 1000
	PhoneNumber string `json: "phoneNumber"`

	// Maximum weight (in Kg) the agent can carry in a single trip
	//
	// required: true
	// max length: 1000
	MaxWeightCapacity int64 `json: "maxWeightCapacity" validate:"required"`

	// Maximum duration (in hr) the agent can work on a single trip
	//
	// required: false
	// max length: 1000
	MaxHourCapacity int64 `json: "maxHourCapacity"`

	// Agent Username/ UserID given by Business (unique to that business)
	//
	// required: false
	// max length: 1000	
	AgentID string `json: "agentID"`

	// Type of Vehicle Agent is using
	//
	// required: false
	// max length: 1000	
	// example: TWO WHEELER | VAN | MINI TRUCK | TRUCK
	TypeOfVehicle string `json: "typeOfVehicle"`
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
	PicURL string `json: "picurl"`

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

	// 12 Digit Aadhar Card number of the Agent
	//
	// required: true
	// max length: 1000	
	AadharNumber string `json: "aadharNumber" validate:"required"`

	// Driving licence number of the agent
	//
	// required: true
	// max length: 1000	
	DrivingLicenceNumber string `json: "drivingLicenceNumber" validate:"required"`

	// Type of Vehicle Agent is using
	//
	// required: true
	// max length: 1000	
	// example: TWO WHEELER | VAN | MINI TRUCK | TRUCK
	TypeOfVehicle string `json: "typeOfVehicle" validate:"required"`

	// PAN Card number of the Agent 
	//
	// required: false
	// max length: 1000	
	PANCardNumber string `json: "panCardNumber"`
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

	// 12 Digit Aadhar Card number of the Agent
	//
	AadharNumber string `json: "aadharNumber" validate:"required"`

	// Driving licence number of the agent
	//
	DrivingLicenceNumber string `json: "drivingLicenceNumber" validate:"required"`

	// Type of Vehicle Agent is using
	//
	TypeOfVehicle string `json: "typeOfVehicle" validate:"required"`

	// PAN Card number of the Agent 
	//	
	PANCardNumber string `json: "panCardNumber"`
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

type AutoAgentsPostSuccess struct {
	// BybID of the businessID
	//
	BybID string `json:"bybID"`

	// uuid (bybID) of all the created agents
	//
	AgentIDs []string `json:"agentIDs"`

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

func (d *AutoNewAgentsRequest) ValidateAutoNewAgentsRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}