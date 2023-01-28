package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hecomp/catchall/internal/models"
	"github.com/hecomp/catchall/internal/utils"
	"github.com/hecomp/catchall/pkg/services"
	"log"
	"net/http"
)

// CatchAll is a http.Handler
type CatchAll struct {
	l       *log.Logger
	service services.CatchAllService
}

var (
	ErrDomainNameInvalid  = errors.New("domain Name %s is invalid")
	ErrPutDeliveredEvent  = errors.New("error put delivered event")
	ErrDomainNameNotFound = errors.New("domain Name not found")

	MakeDeliveredEventSuccess = fmt.Sprintf("delivered event completed!")
	MakeBoncedEventSuccess    = fmt.Sprintf("delivered event completed!")
	MakeGetDomainNameSuccess  = fmt.Sprintf("delivered event completed!")
)

// NewCatchAllHandler creates a CatchAll handler with the given logger
func NewCatchAllHandler(logger *log.Logger, service services.CatchAllService) *CatchAll {
	return &CatchAll{l: logger, service: service}
}

func (c *CatchAll) DeliveredEventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	domainName := mux.Vars(r)["domain_name"]
	isValid, domainN := utils.ValidateDomainName(domainName)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := models.CatchAllResponse{
			Err: fmt.Sprintf("%s %s", domainName, ErrDomainNameInvalid.Error()),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := c.service.PutDeliveredEvent(domainN); err != nil {
		log.Fatalln(ErrPutDeliveredEvent)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&models.CatchAllResponse{
			Err: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.CatchAllResponse{
		Message: MakeDeliveredEventSuccess,
	})
}

func (c *CatchAll) EventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	domainName := mux.Vars(r)["domain_name"]
	isValid, domainN := utils.ValidateDomainName(domainName)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.CatchAllResponse{
			Err: fmt.Sprintf("%s %s", domainName, ErrDomainNameInvalid.Error()),
		})
		return
	}
	if err := c.service.PutBouncedEvent(domainN); err != nil {
		log.Fatalln(ErrPutDeliveredEvent)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&models.CatchAllResponse{
			Err: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.CatchAllResponse{
		Message: MakeBoncedEventSuccess,
	})
}

func (c *CatchAll) GetDomainsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	domainName := mux.Vars(r)["domain_name"]
	isValid, domainN := utils.ValidateDomainName(domainName)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.CatchAllResponse{
			Err: fmt.Sprintf("%s %s", domainName, ErrDomainNameInvalid.Error()),
		})
		return
	}
	resp, err := c.service.GetDomains(domainN)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&models.CatchAllResponse{
			Err: ErrDomainNameNotFound.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&models.CatchAllResponse{
		Message: MakeGetDomainNameSuccess,
		Data:    resp,
	})
}

func (c *CatchAll) HealthHandler(w http.ResponseWriter, r *http.Request) {
	c.l.Println("Checking application health")
	response := map[string]string{
		"status": "UP",
	}
	json.NewEncoder(w).Encode(response)
}
