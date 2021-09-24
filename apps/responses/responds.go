package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// respondwithJSON write json response format
func RespondwithJSON(w http.ResponseWriter, code int, response []byte) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func RespondWithError(w http.ResponseWriter, res *ErrResponse) {
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	RespondwithJSON(w, res.HTTPStatusCode, b)
}
