package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aelmel/ports-infra/client_api/internal/client/port"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NewHTTPHandler(client port.Client) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/port/{key}", PortHandler(client))
	return r
}

func PortHandler(client port.Client) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		portKey := vars["key"]
		log.Info(fmt.Sprintf("Get details for key %s", portKey))

		portDetails, err := client.GetPort(context.Background(), portKey)
		if err != nil {
			log.Warn("error received ", err)
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		port, err := json.Marshal(portDetails)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Write(port)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
	}

}
