
package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bybrisk/agents-api/data"
)

// swagger:route GET /agents/all/{id} agents getAllAgents
// Get summerized details of all the agents registed on a business account.
//
// responses:
//	200: getAllAgentsDetails
//  422: errorValidation
//  501: errorResponse

func (p *Agents) GetAllAgents(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> agents-api Module")
	
	vars := mux.Vars(r)
	id := vars["id"]

	lp := data.GetAllAgents(id)

	err := lp.GetAllAgentsResultToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}