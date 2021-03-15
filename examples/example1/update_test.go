package example1_test

import (
	"context"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/examples/example1"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {

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

	if err := update(logger, ctx, collection, "Susan", "Red", "Atlanta"); err != nil {
		_ = level.Info(logger).Log("err", err.Error())
	}

}

func update(logger log.Logger, ctx context.Context, aCollection *mongo.Collection, fn string, ln string, cy string) error {

	opts := options.Update().SetUpsert(true)

	f := example1.Filter{}
	f.Or().AndFirstNameEqTo(fn).AndLastNameEqTo(ln)

	filterDocument := f.Build()
	_ = level.Info(logger).Log("resulting_filter: ", fmt.Sprintf("%v", filterDocument))

	updateDoc := example1.UpdateDocument{}
	updateDoc.SetAddressCity(cy)

	if ur, err := aCollection.UpdateOne(ctx, filterDocument, updateDoc.Build(), opts); err != nil {
		return err
	} else {
		_ = level.Info(logger).Log("msg", "update result", "upsertedCound", ur.UpsertedCount, "modifiedCount", ur.ModifiedCount)
	}

	return nil
}
