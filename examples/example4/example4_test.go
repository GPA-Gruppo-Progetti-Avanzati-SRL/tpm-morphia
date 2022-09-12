package example4_test

import (
	"context"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/examples/example4"
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
	MongoCollection = "examples"
)

var collection *mongo.Collection

func TestMain(m *testing.M) {

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoUrl))
	requireNoErr(err)

	ctx := context.Background()
	err = client.Connect(ctx)
	requireNoErr(err)

	defer func() {
		err := client.Disconnect(ctx)
		requireNoErr(err)
	}()

	err = client.Ping(ctx, readpref.Primary())
	requireNoErr(err)

	database := client.Database(MongoDb)
	colls, err := database.ListCollectionNames(ctx, bson.D{})
	requireNoErr(err)

	found := false
	for _, n := range colls {
		if n == MongoCollection {
			found = true
			break
		}
	}

	if !found {
		err = database.CreateCollection(ctx, MongoCollection)
		requireNoErr(err)
	}

	collection = database.Collection(MongoCollection)
	removeAll(collection, ctx)

	exitVal := m.Run()

	if !found {
		log.Info().Msg("dropping collection")
		err = collection.Drop(ctx)
		if err != nil {
			log.Error().Err(err).Msg("dropping collection")
		}
	}
	os.Exit(exitVal)
}

func requireNoErr(err error) {
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func removeAll(example1Collection *mongo.Collection, ctx context.Context) {
	deleteResult, err := example1Collection.DeleteMany(ctx, bson.D{})
	requireNoErr(err)
	log.Info().Int64("delete-count", deleteResult.DeletedCount).Msg("remove all")
}

func TestExample4(t *testing.T) {
	t.Log("example4 test case")

	t.Log("insert record")
	err := insert(collection, &example4Data)
	require.NoError(t, err)

	t.Log("add new address to record")
	err = updateWithNewAddress(collection)
	require.NoError(t, err)

	t.Log("modify address to record")
	err = updateExistingAddress(collection)
	require.NoError(t, err)

	t.Log("remove address from record")
	err = deleteExistingAddress(collection)
	require.NoError(t, err)
}

var example4Data = example4.Cliente{
	Ndg:           "123456789",
	CodiceFiscale: "MPRMLS62S21G337J",
	PartitaIVA:    "0123456789",
	Natura:        "SNC",
	Stato:         "A",
	Indirizzi: map[string]example4.Indirizzo{
		"NDR": example4.Indirizzo{
			Indirizzo: "Via del pero",
			Cap:       "01010",
			Localita:  "Paperopoli",
			Provincia: "PP",
			Nazione:   "WD",
		},
	},
	Legati: []example4.Legame{
		{
			Ndg:           "56789",
			Cognome:       "Paperino",
			Nome:          "Paolino",
			CodiceFiscale: "PPRPLN62S21G337J",
			PartitaIVA:    "987654",
			Natura:        "SAS",
		},
	},
	Leganti: []example4.Legame{
		{
			Ndg:           "56789",
			Cognome:       "Paperino",
			Nome:          "Paolino",
			CodiceFiscale: "PPRPLN62S21G337J",
			PartitaIVA:    "987654",
			Natura:        "SAS",
		},
	},
}

func insert(c *mongo.Collection, dto *example4.Cliente) error {
	r, err := c.InsertOne(context.Background(), dto)
	if err != nil {
		return err
	}
	log.Info().Interface("obj-id", r.InsertedID).Msg("insert one")
	return nil
}

func updateWithNewAddress(c *mongo.Collection) error {
	newCli := example4.Cliente{}
	ud := example4.GetUpdateDocument(&newCli)
	ud.SetIndirizziS("ND1", example4.Indirizzo{
		Indirizzo: "Via del pioppo",
		Cap:       "01011",
		Localita:  "Topolinia",
		Provincia: "TP",
		Nazione:   "WD",
	})

	filter := example4.Filter{}
	filter.Or().AndNdgEqTo("123456789")

	upsert := true
	ur, err := c.UpdateOne(context.Background(), filter.Build(), ud.Build(), &options.UpdateOptions{Upsert: &upsert})
	if err != nil {
		return err
	}

	if ur.ModifiedCount != 1 {
		return fmt.Errorf("number of modified documents != 1 (%d)", ur.ModifiedCount)
	}

	log.Info().Interface("update result", ur).Msg("add new address")
	return nil
}

func updateExistingAddress(c *mongo.Collection) error {
	newCli := example4.Cliente{}
	ud := example4.GetUpdateDocument(&newCli)
	ud.SetIndirizziS("ND1", example4.Indirizzo{
		Indirizzo: "Via del pioppo",
		Cap:       "88888",
		Localita:  "Topolinia",
		Provincia: "TP",
		Nazione:   "WD",
	})

	filter := example4.Filter{}
	filter.Or().AndNdgEqTo("123456789")

	upsert := true
	ur, err := c.UpdateOne(context.Background(), filter.Build(), ud.Build(), &options.UpdateOptions{Upsert: &upsert})
	if err != nil {
		return err
	}

	if ur.ModifiedCount != 1 {
		return fmt.Errorf("number of modified documents != 1 (%d)", ur.ModifiedCount)
	}

	log.Info().Interface("update result", ur).Msg("update existing address")
	return nil
}

func deleteExistingAddress(c *mongo.Collection) error {
	newCli := example4.Cliente{}
	ud := example4.GetUpdateDocument(&newCli)
	ud.UnsetIndirizziS("ND1")

	filter := example4.Filter{}
	filter.Or().AndNdgEqTo("123456789")

	upsert := false
	ur, err := c.UpdateOne(context.Background(), filter.Build(), ud.Build(), &options.UpdateOptions{Upsert: &upsert})
	if err != nil {
		return err
	}

	if ur.ModifiedCount != 1 {
		return fmt.Errorf("number of modified documents != 1 (%d)", ur.ModifiedCount)
	}

	log.Info().Interface("update result", ur).Msg("delete existing address")
	return nil
}
