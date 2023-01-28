package app

import (
	"github.com/gorilla/mux"
)

// Routes returns a router with the routes defined
func Routes(catchAll *CatchAll) *mux.Router {
	// create a new serve mux and register the handlers
	r := mux.NewRouter()
	r.HandleFunc("/events/{domain_name}/delivered", catchAll.DeliveredEventsHandler).Methods("PUT")
	r.HandleFunc("/events/{domain_name}/bounced", catchAll.EventsHandler).Methods("PUT")
	r.HandleFunc("/domains/{domain_name}", catchAll.GetDomainsHandler).Methods("GET")
	r.HandleFunc("/health", catchAll.HealthHandler).Methods("GET")
	return r
}
