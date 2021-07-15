package api

import (
	"context"
	"github.com/ONSdigital/dp-content-api/models"
)

//go:generate moq -out mock/contentstore.go -pkg mock . ContentStore

// ContentStore defines the required methods from the data store of content
type ContentStore interface {
	UpsertContent(ctx context.Context, content *models.Content) error
	GetInProgressContentByURL(ctx context.Context, url string) (*models.Content, error)
}
