
package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bybrisk/agents-api/data"
)

// swagger:route GET /agents/delete/{id} agents deleteAgents
// Delete an existing agent registered to a business account.
//
// responses:
//	200: agentsPostResponse
//  422: errorValidation
//  501: errorResponse

func (p *Agents) DeleteAgents(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> agents-api Module")
	
	vars := mux.Vars(r)
	id := vars["id"]

	lp := data.DeleteAgentsByID(id)

	err := lp.ResultToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}