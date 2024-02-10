package example5_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/examples/example5"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"testing"
)

const (
	MongoUrl        = "mongodb://localhost:27017"
	MongoDb         = "tpm_morphia"
	MongoCollection = "example5"
)

var collection *mongo.Collection

func TestMain(m *testing.M) {

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoUrl))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	database := client.Database(MongoDb)
	collection = database.Collection(MongoCollection)

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestExample5(t *testing.T) {
	ctx := context.Background()
	err := removeAll(collection, ctx)
	require.NoError(t, err)

	err = insertARecord(ctx, collection, "123-stella")
	require.NoError(t, err)

	err = findRecord(ctx, collection, "123-stella")
	require.NoError(t, err)

}

func findRecord(ctx context.Context, aCollection *mongo.Collection, sid string) error {
	f := example5.Filter{}
	f.Or().AndSidEqTo(sid)

	filterDocument := f.Build()
	log.Info().Msgf("resulting_filter: %s", fmt.Sprintf("%v", filterDocument))

	cur, err := aCollection.Find(ctx, filterDocument)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	numRecords := consume(ctx, cur)
	log.Info().Msgf("number_of_recs_found %d", numRecords)

	return nil
}

func insertARecord(ctx context.Context, aCollection *mongo.Collection, sid string) error {
	a := example5.Session{
		Sid:        sid,
		Nickname:   "paolino-paperino",
		Remoteaddr: "192.168.1.1",
		Flags:      "none",
	}

	log.Info().Msgf("inserting_record %s", fmt.Sprintf("%v", a))

	r, err := aCollection.InsertOne(ctx, a)

	if err != nil {
		return err
	} else {
		log.Info().Interface("insert-result", r).Send()
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
		log.Info().Msgf("found_document %s", fmt.Sprintf("%v", document))
	}

	if err := cur.Err(); err != nil {
		log.Error().Err(err).Send()
		panic(err.Error())
	}

	return numberOfRecords
}
