package example1_test

import (
	"context"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/examples/example1"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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

	f := example1.Filter{}
	f.Or().AndFirstNameEqTo(fn).AndLastNameEqTo(ln)

	filterDocument := f.Build()
	log.Info().Msg("resulting_filter: " + fmt.Sprintf("%v", filterDocument))

	updateDoc := example1.UpdateDocument{}
	updateDoc.SetAddressCity(cy)

	if ur, err := aCollection.UpdateOne(ctx, filterDocument, updateDoc.Build(), opts); err != nil {
		return err
	} else {
		log.Info().Msgf("update result - upsertedCound: %d, modifiedCount: %d", ur.UpsertedCount, ur.ModifiedCount)
	}

	return nil
}
