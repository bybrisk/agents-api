package data

func AddData (d *NewAgentsRequest) *AgentsPostSuccess{
	//save data to database and return ID
	id := CreateAgents(d)

	//sending response
	var response = AgentsPostSuccess{
		BybID: id,
		Message: "Agent added successfully with agentID",
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