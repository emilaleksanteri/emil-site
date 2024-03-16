package web

import (
	"emil/cmd/web/components"
	"net/http"
)

func HelloWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	name := r.FormValue("name")
	component := HelloPost(name)
	component.Render(r.Context(), w)
}

func CommandsHandler(w http.ResponseWriter, r *http.Request) {
	command := r.URL.Query().Get("command")
	switch command {
	case "./intro":
		component := Intro()
		component.Render(r.Context(), w)
	case "./projects":
		component := Projects()
		component.Render(r.Context(), w)
	case "./contact":
		component := Contact()
		component.Render(r.Context(), w)
	case "./help":
		component := Help()
		component.Render(r.Context(), w)
	case "ls":
		component := Ls()
		component.Render(r.Context(), w)
	case "":
		component := components.TextBlock("commands")
		component.Render(r.Context(), w)
	default:
		component := NoCommand(command)
		component.Render(r.Context(), w)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	component := HomePage()
	component.Render(r.Context(), w)
}
