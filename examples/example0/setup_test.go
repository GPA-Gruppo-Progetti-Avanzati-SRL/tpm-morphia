package example0_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/examples/example0"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MongoUrl        = "mongodb://localhost:27017"
	MongoDb         = "tpm_morphia"
	MongoCollection = "examples"
)

func Setup(logger log.Logger, ctx context.Context) (*mongo.Client, *mongo.Collection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoUrl))
	if err != nil {
		return nil, nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return client, nil, err
	}

	database := client.Database(MongoDb)
	collection := database.Collection(MongoCollection)

	if err := removeAll(logger, collection, ctx); err != nil {
		return client, collection, err
	}

	if err := insertARecord(logger, ctx, collection, "John", "Smith", "Atlanta", "Marietta St."); err != nil {
		return client, collection, err
	}

	if err := insertARecord(logger, ctx, collection, "Colin", "Ward", "Naples", "5th Ave"); err != nil {
		return client, collection, err
	}

	if err := insertARecord(logger, ctx, collection, "Susan", "Red", "", ""); err != nil {
		return client, collection, err
	}

	return client, collection, nil
}

func insertARecord(logger log.Logger, ctx context.Context, aCollection *mongo.Collection, fn string, ln string, city string, strt string) error {
	a := example0.Author{
		FirstName: fn,
		LastName:  ln,
		Age:       30,
		Address:   example0.Address{City: city, Street: strt},
	}

	_ = level.Info(logger).Log("inserting_record", fmt.Sprintf("%v", a))

	r, err := aCollection.InsertOne(ctx, a)
	if err != nil {
		return err
	} else {
		a.OId = r.InsertedID.(primitive.ObjectID)
		if b, err := json.Marshal(a); err != nil {
			return err
		} else {
			_ = level.Info(logger).Log("document", string(b))
		}
	}

	return nil
}

/*
 * Boilerplate code to clear the collection
 */
func removeAll(logger log.Logger, example1Collection *mongo.Collection, ctx context.Context) error {
	deleteResult, err := example1Collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		return err
	}
	_ = level.Info(logger).Log("number_of_documents_deleted", deleteResult.DeletedCount)

	return nil
}

/*
 * Boilerplate code to read the resultset and return the number of records.
 */
func consume(logger log.Logger, ctx context.Context, cur *mongo.Cursor) int {

	numberOfRecords := 0
	for cur.Next(ctx) {
		numberOfRecords++
		var document bson.D
		err := cur.Decode(&document)
		if err != nil {
			_ = level.Error(logger).Log("err", err.Error())
			panic(err.Error())
		}
		_ = level.Info(logger).Log("found_document", fmt.Sprintf("%v", document))
	}

	if err := cur.Err(); err != nil {
		_ = level.Error(logger).Log("err", err.Error())
		panic(err.Error())
	}

	return numberOfRecords
}
