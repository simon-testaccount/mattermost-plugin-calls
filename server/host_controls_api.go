package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (p *Plugin) handleMakeHost(w http.ResponseWriter, r *http.Request) {
	var res httpResponse
	defer p.httpAudit("handleMakeHost", &res, w, r)

	userID := r.Header.Get("Mattermost-User-Id")
	callID := mux.Vars(r)["call_id"]

	var params struct {
		NewHostID string `json:"new_host_id"`
	}
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, requestBodyMaxSizeBytes)).Decode(&params); err != nil {
		res.Err = err.Error()
		res.Code = http.StatusBadRequest
		return
	}

	if err := p.changeHost(userID, callID, params.NewHostID); err != nil {
		p.LogError("handleMakeHost: failed to changeHost", "err", err.Error())
		res.Code = http.StatusInternalServerError
		res.Err = err.Error()
		return
	}

	res.Code = http.StatusOK
	res.Msg = "success"
}