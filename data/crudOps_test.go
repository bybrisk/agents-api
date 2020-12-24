package data_test

import (
	"testing"
	//"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
	"github.com/bybrisk/agents-api/data"
)

func TestGetSingleAgents(t *testing.T) {
	id:="5fddc00a0f3d08cd4146ef7f"
	res:= data.GetSingleAgent(id) 
	if res==nil{
		t.Fail()
	}
}