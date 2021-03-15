package example0_test

import (
	"context"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

func TestFind(t *testing.T) {

	logger := system.GetLogger()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, collection, err := Setup(logger, ctx)

	if client != nil {
		defer client.Disconnect(ctx)
	}

	if err != nil {
		_ = level.Error(logger).Log("msg", err.Error())
		return
	}

	if err := find(logger, ctx, collection, "John", "Ward", "Naples"); err != nil {
		_ = level.Error(logger).Log("msg", err.Error())
		return
	}

	if err := find(logger, ctx, collection, "John", "Ward", "Atlanta"); err != nil {
		_ = level.Error(logger).Log("msg", err.Error())
		return
	}
}

func find(logger log.Logger, ctx context.Context, aCollection *mongo.Collection, fn string, ln string, cy string) error {

	filter := bson.D{
		{"$or", bson.A{bson.D{{"fn", fn}}, bson.D{{"ln", ln}, {"addr.city", cy}}}},
	}

	_ = level.Info(logger).Log("using_filter: ", fmt.Sprintf("%v", filter))

	cur, err := aCollection.Find(ctx, filter)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	numRecords := consume(logger, ctx, cur)
	_ = level.Info(logger).Log("number_of_recs_found", numRecords)
	return nil
}
