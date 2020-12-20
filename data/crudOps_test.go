package data_test

import (
	"testing"
	//"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
	"github.com/bybrisk/agents-api/data"
)

func TestAddData(t *testing.T) {
	agentRequest := &data.NewAgentsRequest{
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
	res:= data.AddData(agentRequest) 
	if res.Message!="Agent added successfully with agentID"{
		t.Fail()
	}

}

func TestUpdateAgents(t *testing.T) {
	updateAgentRequest := &data.UpdateAgentsRequest{
		BybID : "5fddc00a0f3d08cd4146ef7f",
		PicURL : "newimg/pic.jpg",
		AgentName : "BestEver Shashank Prakash",
		AgentType : "DELIVERY",
		PhoneNumber : "8340212623",
		Address : "New A.G Colony, Chetna Samiti",
		MaxWeightCapacity : 30,
		MaxHourCapacity : 6,	
		AgentID : "NEW-39hf4r",
	}
	res:= data.UpdateAgents(updateAgentRequest) 
	if res.Message!="Update Done Successfully"{
		t.Fail()
	}

}