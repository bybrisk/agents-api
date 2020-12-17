
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/agents-api/data"
)

// swagger:route POST /agents/create agents createAgents
// Add an agent (Delivery partner, contracter, etc) to your business account.
//
// responses:
//	200: agentsPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Agents) AddNewAgent (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> agents-api Module")
	agents := &data.NewAgentsRequest{}

	err:=agents.FromJSON(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = agents.Validate()
	if err!=nil {
		p.l.Println("Validation error in POST request -> agents-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//create agent
	agentWithID := data.AddData(agents)

	//writing to the io.Writer
	err = agentWithID.ResultToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}