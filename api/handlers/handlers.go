package handlers

import (
	"encoding/json"
	"github.com/esvas/FinalProject/internal/storages"
	"net/http"
)

func ConnectionHandler(w http.ResponseWriter, _ *http.Request) {
	res := storages.GetResultData()
	resBytes, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resBytes)
}