package example0_test

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

func TestFind(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, collection, err := Setup(ctx)

	if client != nil {
		defer client.Disconnect(ctx)
	}
	require.NoError(t, err)

	err = find(ctx, collection, "John", "Ward", "Naples")
	require.NoError(t, err)

	err = find(ctx, collection, "John", "Ward", "Atlanta")
	require.NoError(t, err)
}

func find(ctx context.Context, aCollection *mongo.Collection, fn string, ln string, cy string) error {

	filter := bson.D{
		{"$or", bson.A{bson.D{{"fn", fn}}, bson.D{{"ln", ln}, {"addr.city", cy}}}},
	}

	log.Info().Str("filter", fmt.Sprintf("%v", filter)).Msg("using filter")

	cur, err := aCollection.Find(ctx, filter)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	numRecords := consume(ctx, cur)
	log.Info().Int("number_of_recs_found", numRecords).Send()
	return nil
}
