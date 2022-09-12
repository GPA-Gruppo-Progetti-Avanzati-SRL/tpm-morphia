package example0_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/examples/example0"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MongoUrl        = "mongodb://localhost:27017"
	MongoDb         = "tpm_morphia"
	MongoCollection = "examples"
)

func Setup(ctx context.Context) (*mongo.Client, *mongo.Collection, error) {
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

	if err := removeAll(collection, ctx); err != nil {
		return client, collection, err
	}

	if err := insertARecord(ctx, collection, "John", "Smith", "Atlanta", "Marietta St."); err != nil {
		return client, collection, err
	}

	if err := insertARecord(ctx, collection, "Colin", "Ward", "Naples", "5th Ave"); err != nil {
		return client, collection, err
	}

	if err := insertARecord(ctx, collection, "Susan", "Red", "", ""); err != nil {
		return client, collection, err
	}

	return client, collection, nil
}

func insertARecord(ctx context.Context, aCollection *mongo.Collection, fn string, ln string, city string, strt string) error {
	a := example0.Author{
		FirstName: fn,
		LastName:  ln,
		Age:       30,
		Address:   example0.Address{City: city, Street: strt},
		Document:  bson.M{"f1": fn, "f2": ln},
	}

	log.Info().Msgf("inserting_record %s", fmt.Sprintf("%v", a))

	r, err := aCollection.InsertOne(ctx, a)
	if err != nil {
		return err
	} else {
		a.OId = r.InsertedID.(primitive.ObjectID)
		if b, err := json.Marshal(a); err != nil {
			return err
		} else {
			log.Info().Msgf("document %s", string(b))
		}
	}

	return nil
}

/*
 * Boilerplate code to clear the collection
 */
func removeAll(example1Collection *mongo.Collection, ctx context.Context) error {
	deleteResult, err := example1Collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		return err
	}
	log.Info().Msgf("number_of_documents_deleted %d", deleteResult.DeletedCount)

	return nil
}

/*
 * Boilerplate code to read the resultset and return the number of records.
 */
func consume(ctx context.Context, cur *mongo.Cursor) int {

	numberOfRecords := 0
	for cur.Next(ctx) {
		numberOfRecords++
		var document bson.D
		err := cur.Decode(&document)
		if err != nil {
			log.Error().Err(err).Send()
			panic(err.Error())
		}
		log.Info().Str("docs", fmt.Sprintf("%v", document)).Msg("found_document")
	}

	if err := cur.Err(); err != nil {
		log.Error().Err(err).Send()
		panic(err.Error())
	}

	return numberOfRecords
}
