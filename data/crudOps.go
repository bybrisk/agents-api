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