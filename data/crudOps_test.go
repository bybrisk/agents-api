package data_test

import (
	"testing"
	"fmt"
	//"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
	"github.com/bybrisk/agents-api/data"
)

/*func TestAddData(t *testing.T) {
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

}*/

/*func TestGetAgents(t *testing.T) {
	res:=data.GetSingleAgent("600e9c8ae3ffe1ac1f5f3a32")
	fmt.Println(res)
}*/

func TestAutoAddAgents(t *testing.T) {
	autoAdd := &data.AutoNewAgentsRequest{
		BusinessID:"6079d92f44956a457f5f275d",
		AgentNum:10,
		MaxWeightCapacity:30,
	}
	res:= data.AutoAddAgents(autoAdd)
	fmt.Println(res)
}