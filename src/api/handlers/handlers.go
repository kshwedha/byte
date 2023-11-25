package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "/")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("components/templates", "home.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LookUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Look up handler...")
}

func GetModulesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Getting all modules...")
}

func GetModuleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	moduleID := vars["module"]
	fmt.Fprintf(w, "Getting module with ID: %s", moduleID)
}

func CreateModuleHandler(w http.ResponseWriter, r *http.Request) {
	// verify user privilege
	fmt.Fprintln(w, "Creating a new module...")
}

func UpdateModuleHandler(w http.ResponseWriter, r *http.Request) {
	// verify user privilege
	vars := mux.Vars(r)
	moduleID := vars["module"]
	fmt.Fprintf(w, "Updating user with ID: %s", moduleID)
}

func DeleteModuleHandler(w http.ResponseWriter, r *http.Request) {
	// verify user privilege
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Fprintf(w, "Deleting user with ID: %s", userID)
}

func DeleteModulesHandler(w http.ResponseWriter, r *http.Request) {
	// verify user privilege
	fmt.Fprintln(w, "Deleting all modules...")
}
