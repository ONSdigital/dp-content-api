package api

import (
	"context"
	"encoding/json"
	"github.com/ONSdigital/dp-content-api/models"
	dphttp "github.com/ONSdigital/dp-net/http"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func (api *API) AddContentHandler(w http.ResponseWriter, r *http.Request) {

	defer dphttp.DrainBody(r)
	ctx := r.Context()
	logData := log.Data{}

	content, err := ParseContent(ctx, r)
	if err != nil {
		handleError(ctx, err, w, logData)
		return
	}

	// fail if the URL with an unapproved status exists - only approved should exist?

	if err = api.contentStore.UpsertContent(ctx, content); err != nil {
		handleError(ctx, err, w, logData)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = WriteJSONBody(ctx, content, w, logData)
	if err != nil {
		handleError(ctx, err, w, logData)
		return
	}

	log.Event(ctx, "add content request completed successfully", log.INFO, logData)
}

func ParseContent(ctx context.Context, r *http.Request) (*models.Content, error) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var content models.Content
	err = json.Unmarshal(b, &content)
	if err != nil {
		log.Error(ctx, "failed to parse content json body", err)
		return nil, ErrUnableToParseJSON
	}

	content.ID, err = NewID()
	if err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	content.CollectionID = vars["collection_id"]
	content.URL = "/" + vars["url"]

	return &content, nil
}
