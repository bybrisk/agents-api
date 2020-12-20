package data_test

import (
	"testing"
	//"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
	"github.com/bybrisk/agents-api/data"
)

func TestValidateAddAgentRequest(t *testing.T) {
	agentRequest := data.NewAgentsRequest{
		PicURL : "img/pic.jpg",
		AgentName : "Shashank Prakash",
		AgentType : "DELIVERY",
		BusinessID : "ivuivrevrhhr",
		PhoneNumber : "9340212623",
		Address : "A.G Colony, Chetna Samiti",
		MaxWeightCapacity : 30,
		MaxHourCapacity : 8,	
		AgentID : "39hf4r",
	}
	if agentRequest.Validate()!= nil {
		t.Fail()
	}

}

