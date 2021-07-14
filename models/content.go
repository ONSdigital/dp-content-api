package models

import (
	"time"
)

// Content represents information related to a single piece of content
type Content struct {
	ID           string     `bson:"_id,omitempty"           json:"id,omitempty"`
	URL          string     `bson:"url,omitempty"           json:"url,omitempty"`
	CollectionID string     `bson:"collection_id,omitempty" json:"collection_id,omitempty"`
	ContentType  string     `bson:"content_type,omitempty"  json:"content_type,omitempty"`
	Content      string     `bson:"content,omitempty"       json:"content,omitempty"`
	PublishDate  *time.Time `bson:"publish_date,omitempty"  json:"publish_date,omitempty"`
	LastUpdated  time.Time  `bson:"last_updated,omitempty"  json:"-"`
}
