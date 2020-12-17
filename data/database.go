package data

import (
	"fmt"
	//"github.com/bybrisk/structs"
	//"go.mongodb.org/mongo-driver/bson"
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
