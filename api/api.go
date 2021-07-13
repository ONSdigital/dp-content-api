package api

import (
	"context"
	"encoding/json"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

//API provides a struct to wrap the api around
type API struct {
	Router       *mux.Router
	contentStore ContentStore
}

//Setup function sets up the api and returns an api
func Setup(ctx context.Context, r *mux.Router, contentStore ContentStore) *API {
	api := &API{
		Router:       r,
		contentStore: contentStore,
	}

	//r.HandleFunc("/collections", api.AddCollectionHandler).Methods(http.MethodPost)
	return api
}

// WriteJSONBody marshals the provided interface into json, and writes it to the response body.
func WriteJSONBody(ctx context.Context, v interface{}, w http.ResponseWriter, data log.Data) error {

	body, err := json.Marshal(v)
	if err != nil {
		handleError(ctx, err, w, data)
		return err
	}

	if _, err := w.Write(body); err != nil {
		return err
	}

	return nil
}

// NewID returns a new UUID
var NewID = func() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}
