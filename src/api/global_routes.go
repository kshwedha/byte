package api

import (
	"github.com/gorilla/mux"
	"github.com/kshwedha/core/src/api/handlers"
)

// api design
// Global
// {domain}/{endpoint}
// {domain}/{version}/{endpoint}
// regular
// {domain}/{catalog}/{version}
func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.RootHandler).Methods("GET")
	router.HandleFunc("/home", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/lookup", handlers.LookUpHandler).Methods("GET")
	router.HandleFunc("/modules", handlers.GetModulesHandler).Methods("GET")
	router.HandleFunc("/module/{module}", handlers.GetModuleHandler).Methods("GET")
	router.HandleFunc("/modules", handlers.CreateModuleHandler).Methods("POST")
	router.HandleFunc("/module/{module}", handlers.UpdateModuleHandler).Methods("PUT")
	router.HandleFunc("/module/{module}", handlers.DeleteModuleHandler).Methods("DELETE")
	return router
}
