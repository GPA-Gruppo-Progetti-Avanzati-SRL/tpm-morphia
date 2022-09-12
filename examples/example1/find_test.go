package example1_test

import (
	"context"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/examples/example1"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

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

	f := example1.Filter{}
	f.Or().AndFirstNameEqTo(fn)
	f.Or().AndLastNameEqTo(ln).AndAddressCityEqTo(cy)

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
