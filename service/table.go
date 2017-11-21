package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func tableHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id := req.FormValue("id")
		name := req.FormValue("name")
		major := req.FormValue("major")
		formatter.HTML(w, http.StatusOK, "table_in", struct {
			ID    string
			NAME  string
			MAJOR string
		}{ID: id, NAME: name, MAJOR: major})
	}
}
