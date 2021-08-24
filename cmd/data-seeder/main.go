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

	processLine := func(s string) {
		addContentToMongoDB(s, ctx, mongodb)
	}

	processFileLines(processLine)
}

func addContentToMongoDB(s string, ctx context.Context, mongodb *mongo.Mongo) {
	content := &models.Content{
		URL:          s,
		CollectionID: "coll-123",
		ContentType:  "test",
		Content:      "123",
		Approved:     true,
		PublishDate:  &time.Time{},
	}
	id, err := api.NewID()
	if err != nil {
		log.Error(ctx, "error initialising mongo", err)
		os.Exit(1)
	}

	content.ID = id
	mongodb.UpsertContent(ctx, content)
}

func processFileLines(processFunc func(string)) {

	f, err := os.Open("cmd/data-seeder/onsurls.txt")
	if err != nil {
		fmt.Println("error opening file ", err)
		os.Exit(1)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	for s.Scan() {
		processFunc(s.Text())
		//fmt.Println(s.Text())
	}
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
