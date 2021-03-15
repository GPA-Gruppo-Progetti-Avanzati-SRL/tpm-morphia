package example0_test

import (
	"context"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {

	logger := system.GetLogger()

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	client, collection, err := Setup(logger, ctx)

	if client != nil {
		defer client.Disconnect(ctx)
	}

	if err != nil {
		_ = level.Info(logger).Log("err", err.Error())
	}

	if err := update(logger, ctx, collection, "Susan", "Red", "Atlanta"); err != nil {
		_ = level.Info(logger).Log("err", err.Error())
	}
}

func update(logger log.Logger, ctx context.Context, aCollection *mongo.Collection, fn string, ln string, cy string) error {

	opts := options.Update().SetUpsert(true)

	filter := bson.D{
		{"$and", bson.A{bson.D{{"fn", fn}}, bson.D{{"ln", ln}}}},
	}

	updateDoc := bson.D{{"$set", bson.D{{"addr.city", cy}}}}
	if ur, err := aCollection.UpdateOne(ctx, filter, updateDoc, opts); err != nil {
		return err
	} else {
		_ = level.Info(logger).Log("msg", "update result", "upsertedCound", ur.UpsertedCount, "modifiedCount", ur.ModifiedCount)
	}

	return nil
}
