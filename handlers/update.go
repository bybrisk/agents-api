
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/agents-api/data"
)

// swagger:route POST /agents/update agents updateAgents
// Modify an existing agent using BybID of this agent.
//
// responses:
//	200: agentsPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Agents) UpdateAgent (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> agents-api Module")
	NewAgents := &data.UpdateAgentsRequest{}

	err:=NewAgents.FromJSONUpdateRequest(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data 
	err = NewAgents.ValidateUpdateRequest()
	if err!=nil {
		p.l.Println("Validation error in POST request -> agents-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//update agents
	response := data.UpdateAgents(NewAgents)

	//writing to the io.Writer
	err = response.ResultToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}