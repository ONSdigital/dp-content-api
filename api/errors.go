package api

import (
	"context"
	"errors"
	"github.com/ONSdigital/dp-content-api/models"
	"github.com/ONSdigital/log.go/v2/log"
	"net/http"
)

var (
	// errors that should return a 400 status
	badRequest = map[error]bool{
		ErrUnableToParseJSON: true,
	}

	ErrUnableToParseJSON = errors.New("failed to parse json body")
)

func handleError(ctx context.Context, err error, w http.ResponseWriter, logData log.Data) {
	var status int
	switch {

	case badRequest[err]:
		status = http.StatusBadRequest
	default:
		status = http.StatusInternalServerError
	}

	if logData == nil {
		logData = log.Data{}
	}

	response := models.ErrorsResponse{
		Errors: []models.ErrorResponse{
			{
				Message: err.Error(),
			},
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	WriteJSONBody(ctx, response, w, logData)
	log.Error(ctx, "request unsuccessful", err, logData)
}
