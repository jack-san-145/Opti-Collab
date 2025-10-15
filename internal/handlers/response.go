package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, r *http.Request, data any) {
	w.Header().Set("Content-Type", "application/json")
	enc_json := json.NewEncoder(w)
	enc_json.SetEscapeHTML(false) // Don’t change &, <, or > to /u0026,/u003c,/u003e ,just keep them as they are in my JSON
	err := enc_json.Encode(data)
	if err != nil {
		fmt.Println("error while encoding data - ", err)
		return
	}
}
