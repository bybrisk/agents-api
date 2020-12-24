
package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bybrisk/agents-api/data"
)

// swagger:route GET /agents/one/{id} agents getOneAgents
// Get full details of a single agent by it's BybID.
//
// responses:
//	200: getSingleAgentsDetails
//  422: errorValidation
//  501: errorResponse

func (p *Agents) GetOneAgent(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> agents-api Module")
	
	vars := mux.Vars(r)
	id := vars["id"]

	lp := data.GetSingleAgent(id)

	err := lp.GetOneAgentResultToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}