package prover

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/OffchainLabs/prysm/v7/monitoring/tracing/trace"
	"github.com/OffchainLabs/prysm/v7/network/httputil"
)

// SubmitExecutionProof handles POST requests to /eth/v1/prover/execution_proofs.
// It receives execution proofs from provers and logs them.
func (s *Server) SubmitExecutionProof(w http.ResponseWriter, r *http.Request) {
	_, span := trace.StartSpan(r.Context(), "prover.SubmitExecutionProof")
	defer span.End()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		httputil.HandleError(w, "Could not read request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(body) == 0 {
		httputil.HandleError(w, "No data submitted", http.StatusBadRequest)
		return
	}

	// Parse the JSON to extract fields for logging
	var proof map[string]any
	if err := json.Unmarshal(body, &proof); err != nil {
		httputil.HandleError(w, "Could not decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	log.Info("Received execution proof")

	w.WriteHeader(http.StatusOK)
}
