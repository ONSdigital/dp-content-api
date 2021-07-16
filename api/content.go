package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/ONSdigital/dp-content-api/models"
	"github.com/ONSdigital/dp-mongodb/v2/mongodb"
	dphttp "github.com/ONSdigital/dp-net/http"
	dprequest "github.com/ONSdigital/dp-net/request"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

func (api *API) AddCollectionContentHandler(w http.ResponseWriter, r *http.Request) {

	defer dphttp.DrainBody(r)
	ctx := r.Context()
	logData := log.Data{}

	content, err := ParseContent(ctx, r)
	if err != nil {
		handleError(ctx, err, w, logData)
		return
	}

	logData["collection_id"] = content.CollectionID
	logData["url"] = content.URL

	_, err = api.contentStore.GetInProgressContentByURL(ctx, content.URL)
	if err != nil && !mongodb.IsErrNoDocumentFound(err) {
		handleError(ctx, err, w, logData)
		return
	}
	if err == nil {
		handleError(ctx, ErrContentAlreadyInProgress, w, logData)
		return
	}

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

func (api *API) GetCollectionContentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	collectionID := vars["collection_id"]
	url := "/" + vars["url"]

	logData := log.Data{
		"collection_id": collectionID,
		"url":           url,
	}

	content, err := api.contentStore.GetInProgressContentByURL(ctx, url)
	if err != nil {
		// todo: do not check for mongo specific errors here - create a 'domain' error type to return from mongo package
		if mongodb.IsErrNoDocumentFound(err) {
			handleError(ctx, err, w, logData)
			return
		}

		handleError(ctx, err, w, logData)
		return
	}

	if collectionID != content.CollectionID {
		// content is being edited in another collection
		// return not found? or something more specific to suggest it's in another collection?
		handleError(ctx, ErrInProgressContentNotFound, w, logData)
		return
	}

	rawContent, err := base64.StdEncoding.DecodeString(content.Content)
	if err != nil {
		handleError(ctx, err, w, logData)
		return
	}

	// todo: handle other content types - currently assumes JSON, but should set other content type headers as required
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(rawContent); err != nil {
		handleError(ctx, err, w, logData)
		return
	}
}

func (api *API) PatchCollectionContentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collectionID := vars["collection_id"]
	url := "/" + vars["url"]

	logData := log.Data{
		"collection_id": collectionID,
		"url":           url,
	}
	ctx := r.Context()

	// unmarshal and validate the patch array
	patches, err := CreatePatches(r.Body)
	if err != nil {
		handleError(ctx, err, w, logData)
		return
	}
	logData["patch_list"] = patches

	err = api.contentStore.PatchContent(ctx, url, patches)
	if err != nil {
		handleError(ctx, err, w, logData)
		return
	}

	// set content type, marshal and write response
	setJSONPatchContentType(w)
	if err := WriteJSONBody(ctx, patches, w, logData); err != nil {
		handleError(ctx, err, w, logData)
		return
	}
}

func setJSONPatchContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json-patch+json")
}

func CreatePatches(reader io.Reader) ([]dprequest.Patch, error) {
	patches := []dprequest.Patch{}

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return []dprequest.Patch{}, ErrReadingBody
	}

	err = json.Unmarshal(bytes, &patches)
	if err != nil {
		return []dprequest.Patch{}, ErrParsingBody
	}

	for _, patch := range patches {
		if err := patch.Validate(dprequest.OpAdd, dprequest.OpRemove, dprequest.OpReplace); err != nil {
			return []dprequest.Patch{}, err
		}
	}
	return patches, nil
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
