package mongo

import (
	"context"
	"errors"
	"github.com/ONSdigital/dp-content-api/models"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	dpMongoHealth "github.com/ONSdigital/dp-mongodb/v2/health"
	dpMongoDriver "github.com/ONSdigital/dp-mongodb/v2/mongodb"
	dprequest "github.com/ONSdigital/dp-net/request"
	"github.com/ONSdigital/log.go/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	connectTimeoutInSeconds = 5
	queryTimeoutInSeconds   = 15
)

// Mongo represents a simplistic MongoDB configuration.
type Mongo struct {
	healthClient      *dpMongoHealth.CheckMongoClient
	Database          string
	ContentCollection string
	Connection        *dpMongoDriver.MongoConnection
	Username          string
	Password          string
	URI               string
	IsSSL             bool
}

func (m *Mongo) getConnectionConfig() *dpMongoDriver.MongoConnectionConfig {
	return &dpMongoDriver.MongoConnectionConfig{
		ConnectTimeoutInSeconds:       connectTimeoutInSeconds,
		QueryTimeoutInSeconds:         queryTimeoutInSeconds,
		Username:                      m.Username,
		Password:                      m.Password,
		ClusterEndpoint:               m.URI,
		Database:                      m.Database,
		Collection:                    m.ContentCollection,
		IsSSL:                         m.IsSSL,
		IsWriteConcernMajorityEnabled: false,
		IsStrongReadConcernEnabled:    false,
	}
}

// Init creates a new mongoConnection with a strong consistency and a write mode of "majority".
func (m *Mongo) Init() error {
	if m.Connection != nil {
		return errors.New("datastore connection already exists")
	}

	mongoConnection, err := dpMongoDriver.Open(m.getConnectionConfig())
	if err != nil {
		return err
	}
	m.Connection = mongoConnection
	databaseCollectionBuilder := make(map[dpMongoHealth.Database][]dpMongoHealth.Collection)
	databaseCollectionBuilder[(dpMongoHealth.Database)(m.Database)] = []dpMongoHealth.Collection{(dpMongoHealth.Collection)(m.ContentCollection)}

	// Create client and health client from session AND collections
	client := dpMongoHealth.NewClientWithCollections(mongoConnection, databaseCollectionBuilder)

	m.healthClient = &dpMongoHealth.CheckMongoClient{
		Client:      *client,
		Healthcheck: client.Healthcheck,
	}

	return nil
}

// Close closes the mongo session and returns any error
func (m *Mongo) Close(ctx context.Context) error {
	if m.Connection == nil {
		return errors.New("cannot close a empty connection")
	}
	return m.Connection.Close(ctx)
}

// Checker is called by the health check library to check the health state of this mongoDB instance
func (m *Mongo) Checker(ctx context.Context, state *healthcheck.CheckState) error {
	return m.healthClient.Checker(ctx, state)
}

func (m *Mongo) UpsertContent(ctx context.Context, content *models.Content) error {

	update := bson.M{
		"$set": content,
		"$setOnInsert": bson.M{
			"last_updated": time.Now(),
		},
	}

	_, err := m.Connection.C(m.ContentCollection).UpsertId(ctx, content.ID, update)

	return err
}

func (m *Mongo) PatchContent(ctx context.Context, url string, patches []dprequest.Patch) error {

	query := bson.D{{"url", url}}

	// create update query from updatedFilter and newly generated eTag
	updates := bson.M{}

	// iterate patches and add updates
	for _, patch := range patches {
		switch patch.Path {
		case "publish_date":
			updates["publish_date"] = patch.Value
		case "approved":
			updates["approved"] = patch.Value
		case "content":
			updates["content"] = patch.Value
		}
	}

	update, err := dpMongoDriver.WithUpdates(bson.M{
		"$set": updates,
	})
	if err != nil {
		return err
	}

	result, err := m.Connection.C(m.ContentCollection).Update(ctx, query, update)

	log.Info(ctx, "patched content", log.Data{"url": url, "result": result})

	return err
}

// GetInProgressContentByURL retrieves content for the given URL that is not yet approved
func (m *Mongo) GetInProgressContentByURL(ctx context.Context, url string) (*models.Content, error) {

	// todo: add status to query, so only in progress content is retrieved

	query := bson.D{{"url", url}}
	result := &models.Content{}

	err := m.Connection.
		C(m.ContentCollection).
		FindOne(ctx, query, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
