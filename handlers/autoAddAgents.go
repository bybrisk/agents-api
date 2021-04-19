
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/agents-api/data"
)

// swagger:route POST /agents/create/auto agents autoCreateAgents
// Automatically add required number of agents (Delivery partner, contracter, etc) to your business account.
//
// responses:
//	200: autoAgentsPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Agents) AutoAddNewAgent (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> agents-api Module")
	agents := &data.AutoNewAgentsRequest{}

	err:=agents.FromJSONToAutoNewAgentsRequest(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = agents.ValidateAutoNewAgentsRequest()
	if err!=nil {
		p.l.Println("Validation error in POST request -> agents-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//create agent
	agentWithID := data.AutoAddAgents(agents)

	//writing to the io.Writer
	err = agentWithID.AutoAgentsPostSuccessToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}