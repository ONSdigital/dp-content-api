package api

import (
	"context"
	"github.com/ONSdigital/dp-content-api/models"
	dprequest "github.com/ONSdigital/dp-net/request"
	"time"
)

//go:generate moq -out mock/contentstore.go -pkg mock . ContentStore

// ContentStore defines the required methods from the data store of content
type ContentStore interface {
	UpsertContent(ctx context.Context, content *models.Content) error
	GetInProgressContentByURL(ctx context.Context, url string) (*models.Content, error)
	GetCollectionContentByURL(ctx context.Context, url, collectionID string) (*models.Content, error)
	PatchContent(ctx context.Context, url, collectionID string, patches []dprequest.Patch) error
	GetPublishedContentByURL(ctx context.Context, url string) (*models.Content, error)
	GetNextPublishDate(ctx context.Context, url string) (*time.Time, error)
}
