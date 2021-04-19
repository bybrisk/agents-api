package data

import (
	"fmt"
	//"github.com/bybrisk/structs"
	"strconv"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	log "github.com/sirupsen/logrus"
)

var resultID string 

//Database Funcs
func CreateAgents (agents *NewAgentsRequest) string {
	collectionName := shashankMongo.DatabaseName.Collection("agents")
	result, insertErr := collectionName.InsertOne(shashankMongo.CtxForDB, agents)
	if insertErr != nil {
		log.Error("Create Agents ERROR:")
		log.Error(insertErr)
	} else {
		fmt.Println("createAgents() API result:", result)

		newID := result.InsertedID
		fmt.Println("createAgents() newID:", newID)
		resultID = newID.(primitive.ObjectID).Hex()
	}
	return resultID
}

func UpdateAgentsByID(agents *UpdateAgentsRequest) int64 {
	collectionName := shashankMongo.DatabaseName.Collection("agents")
	id, _ := primitive.ObjectIDFromHex(agents.BybID)
	filter := bson.M{"_id": id}
	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.D{{Key: "$set", Value: agents}})
	if err != nil {
		log.Error("updateAgentsByID ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}

func AddBybIDToAgents(agents *NewAgentsRequest) int64 {
	collectionName := shashankMongo.DatabaseName.Collection("agents")
	id, _ := primitive.ObjectIDFromHex(agents.BybID)
	filter := bson.M{"_id": id}
	updateResult, err := collectionName.UpdateOne(shashankMongo.CtxForDB, filter, bson.D{{Key: "$set", Value: agents}})
	if err != nil {
		log.Error("AddBybIDToAgents ERROR:")
		log.Error(err)
	}
	return updateResult.ModifiedCount
}

func AllAgentsByBusinessID (docID string) []AgentSummaryDetail {
	var agents []AgentSummaryDetail
	collectionName := shashankMongo.DatabaseName.Collection("agents")

	cursor, err := collectionName.Find(shashankMongo.CtxForDB, bson.M{"businessid":docID})
	if err != nil {
		log.Error("AllAgentsByBusinessID Read ERROR : ")
		log.Error(err)
	}
	if err = cursor.All(shashankMongo.CtxForDB, &agents); err != nil {
		log.Error("AllAgentsByBusinessID Write ERROR : ")
		log.Error(err)
	}

	return agents
}

func DeleteAgentFromDB (docID string) int64{
	collectionName := shashankMongo.DatabaseName.Collection("agents")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}

	deleteResult, err := collectionName.DeleteOne(shashankMongo.CtxForDB, filter)
	if err != nil {
		log.Error("DeleteAgentFromDB ERROR:")
		log.Error(err)
	}
	return deleteResult.DeletedCount
	
}

func FetchAgentFromDB (docID string) *SingleAgentResponse {
	var agent *SingleAgentResponse
	collectionName := shashankMongo.DatabaseName.Collection("agents")
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
	
	err:= collectionName.FindOne(shashankMongo.CtxForDB, filter).Decode(&agent)
	if err != nil {
		log.Error("FetchAgentFromDB ERROR:")
		log.Error(err)
	}
	return agent
}

func AddRequiredAgents(d *AutoNewAgentsRequest) []string {
	collectionName := shashankMongo.DatabaseName.Collection("agents")

	//initial_interface := []interface{}{}
	meta_interface:= []interface{}{}
	
	for i := 0; i<int(d.AgentNum); i++ {
		var newAgent AutoNewAgentsRequest
		newAgent.AgentName = "Agent" + strconv.Itoa(i+1)
		newAgent.BusinessID = d.BusinessID
		newAgent.PhoneNumber = "9122XXXXXX"
		newAgent.MaxWeightCapacity = d.MaxWeightCapacity
		newAgent.MaxHourCapacity = 8
		newAgent.AgentID = "Agent00" + strconv.Itoa(i+1)
		newAgent.TypeOfVehicle = "TWO WHEELER"

		meta_interface = append([]interface{}{newAgent}, meta_interface...)
	}

	//fmt.Println(meta_interface)
	

	insertManyResult, err := collectionName.InsertMany(shashankMongo.CtxForDB, meta_interface)
	if err != nil {
		log.Error("AddRequiredAgents ERROR:")
		log.Error(err)
	}

	var resultIDArr []string

	//fmt.Println(insertManyResult.InsertedIDs)
	for _,val:=range insertManyResult.InsertedIDs{
		resultID = val.(primitive.ObjectID).Hex()
		resultIDArr = append(resultIDArr,resultID)
	}
	
	return resultIDArr
}