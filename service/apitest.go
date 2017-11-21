package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ID    string `json:"id"`
			NAME  string `json:"name"`
			MAJOR string `json:"major"`
		}{ID: "15331157", NAME: "Caroline", MAJOR: "Software Engineering"})
	}
}
