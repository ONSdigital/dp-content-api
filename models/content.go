package models

import (
	"time"
)

// Content represents information related to a single piece of content
type Content struct {
	ID          string     `bson:"_id,omitempty"          json:"id,omitempty"`
	Name        string     `bson:"name,omitempty"         json:"name,omitempty"`
	PublishDate *time.Time `bson:"publish_date,omitempty" json:"publish_date,omitempty"`
	LastUpdated time.Time  `bson:"last_updated,omitempty" json:"-"`
}
