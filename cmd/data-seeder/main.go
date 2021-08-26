package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ONSdigital/dp-content-api/api"
	"github.com/ONSdigital/dp-content-api/config"
	"github.com/ONSdigital/dp-content-api/models"
	"github.com/ONSdigital/dp-content-api/mongo"
	"github.com/ONSdigital/log.go/v2/log"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	err, mongodb := initMongo(ctx)
	if err != nil {
		log.Error(ctx, "error initialising mongo", err)
		os.Exit(1)
	}

	processLines := func(s []string) {
		addContentToMongoDB(s, ctx, mongodb)
	}

	processFileLines(processLines)
}

func addContentToMongoDB(lines []string, ctx context.Context, mongodb *mongo.Mongo) {

	now := time.Now()

	var updates []interface{}

	for _, line := range lines {
		content := &models.Content{
			URL:          line,
			CollectionID: "coll-123",
			ContentType:  "test",
			Content:      "123",
			Approved:     true,
			PublishDate:  &now,
		}
		id, err := api.NewID()
		if err != nil {
			log.Error(ctx, "error initialising mongo", err)
			os.Exit(1)
		}

		content.ID = id
		updates = append(updates, content)
	}

	_, err := mongodb.Connection.C(mongodb.ContentCollection).Insert(ctx, updates)
	if err != nil {
		log.Error(ctx, "error inserting content into mongo", err)
		os.Exit(1)
	}
}

func processFileLines(processFunc func([]string)) {

	f, err := os.Open("cmd/data-seeder/onsurls.txt")
	if err != nil {
		fmt.Println("error opening file ", err)
		os.Exit(1)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	var lines []string
	i := 0
	batchSize := 100

	for s.Scan() {
		lines = append(lines, s.Text())
		i++
		if i == batchSize {
			processFunc(lines)
			lines = []string{}
		}
	}
	processFunc(lines)
}

func initMongo(ctx context.Context) (error, *mongo.Mongo) {
	cfg, err := config.Get()
	if err != nil {
		return err, nil
	}
	log.Event(ctx, "loaded config", log.INFO, log.Data{"config": cfg})

	mongodb := &mongo.Mongo{
		ContentCollection: cfg.MongoConfig.ContentCollection,
		Database:          cfg.MongoConfig.ContentDatabase,
		Username:          cfg.MongoConfig.Username,
		Password:          cfg.MongoConfig.Password,
		IsSSL:             cfg.MongoConfig.IsSSL,
		URI:               cfg.MongoConfig.BindAddr,
	}
	err = mongodb.Init()
	if err != nil {
		return err, nil
	}
	return err, mongodb
}
