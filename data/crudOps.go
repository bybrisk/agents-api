package data

func AddData (d *NewAgentsRequest) *AgentsPostSuccess{
	//save data to database and return ID
	id := CreateAgents(d)
	d.BybID = id
	res:= AddBybIDToAgents(d)
	//sending response
	var response AgentsPostSuccess
	if res == 1 {
		response = AgentsPostSuccess{
			BybID: id,
			Message: "Agent added successfully with agentID",
		}
	} else {
		response = AgentsPostSuccess{
			BybID: id,
			Message: "Error adding agents!",
		}
	}

	return &response
}

func UpdateAgents(d *UpdateAgentsRequest) *AgentsPostSuccess {

	res := UpdateAgentsByID(d)

	var response AgentsPostSuccess
	//sending response
	if res == 1 {
		response = AgentsPostSuccess{
			BybID: d.BybID,
			Message: "Update Done Successfully",
		}
	} else {
		response = AgentsPostSuccess{
			BybID: d.BybID,
			Message: "Error updating agents!",
		}
	}

	return &response
}

func GetAllAgents (docID string) *AgentResponseBulk {
	var response AgentResponseBulk
	agentsArray := AllAgentsByBusinessID(docID)

	response = AgentResponseBulk{
		Result: agentsArray,
		BusinessID: docID,
	}
	
	return &response
}

func DeleteAgentsByID (docID string) *AgentsPostSuccess{

	var response AgentsPostSuccess
	res := DeleteAgentFromDB(docID)

	if res == 1 {
		response = AgentsPostSuccess{
			BybID: docID,
			Message: "Agent deleted successfully",
		}
	} else {
		response = AgentsPostSuccess{
			BybID: docID,
			Message: "Error deleting agents!",
		}
	}

	return &response
}

func GetSingleAgent (docID string) *SingleAgentResponse {
	
	res := FetchAgentFromDB(docID)
	return res
}