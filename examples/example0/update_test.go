package example0_test

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	client, collection, err := Setup(ctx)

	if client != nil {
		defer client.Disconnect(ctx)
	}

	require.NoError(t, err)

	err = update(ctx, collection, "Susan", "Red", "Atlanta")
	require.NoError(t, err)

}

func update(ctx context.Context, aCollection *mongo.Collection, fn string, ln string, cy string) error {

	opts := options.Update().SetUpsert(true)

	filter := bson.D{
		{"$and", bson.A{bson.D{{"fn", fn}}, bson.D{{"ln", ln}}}},
	}

	updateDoc := bson.D{{"$set", bson.D{{"addr.city", cy}}}}
	if ur, err := aCollection.UpdateOne(ctx, filter, updateDoc, opts); err != nil {
		return err
	} else {
		log.Info().Msgf("update result - upsertedCound %d, modifiedCount %d", ur.UpsertedCount, ur.ModifiedCount)
	}

	return nil
}
