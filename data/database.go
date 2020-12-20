package data

import (
	"fmt"
	//"github.com/bybrisk/structs"
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
