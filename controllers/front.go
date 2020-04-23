package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	tc := newWorldTimeController()
	http.Handle("/time", *tc)
	http.Handle("/time/", *tc)
	http.Handle("/times", *tc)
	http.Handle("/times/", *tc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
