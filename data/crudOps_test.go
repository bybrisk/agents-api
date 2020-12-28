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
		AgentName : "Sanjay sinha",
		AgentType : "INSTANT",
		BusinessID : "abcd",
		PhoneNumber : "9234512623",
		Address : "A.G Colony, R Block",
		MaxWeightCapacity : 30,
		MaxHourCapacity : 6,	
		AgentID : "rty1",
	}
	res:= data.AddData(agentRequest) 
	if res.Message!="Agent added successfully with agentID"{
		t.Fail()
	}

}