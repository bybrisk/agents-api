package data

import (
	"encoding/json"
	"io"
)	

func (d *NewAgentsRequest) FromJSON (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}


func (d *AgentsPostSuccess) ResultToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *AgentResponseBulk) GetAllAgentsResultToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *UpdateAgentsRequest) FromJSONUpdateRequest (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}