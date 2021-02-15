package example1_test

import (
	"context"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/examples/example1"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"
)

func TestUpdateexample1(t *testing.T) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		t.Fatal(err)
	}

	database := client.Database("tpm_morphia")
	example1Collection := database.Collection("example1s")

	if err := removeAll(t, err, example1Collection, ctx); err != nil {
		t.Fatal(err)
	}

	if err := insertAFirstRecord(t, ctx, example1Collection, "John", "Smith"); err != nil {
		t.Fatal(err)
	}

	if err := insertASecondRecord(t, ctx, example1Collection, "Colin", "Smith"); err != nil {
		t.Fatal(err)
	}

	if err := insertAThirdRecord(t, ctx, example1Collection, "Susan", "Pitt"); err != nil {
		t.Fatal(err)
	}

	f := example1.Filter{}
	_ = f.Or().AndFirstNameEqTo("Susan")
	opts := options.Update().SetUpsert(true)

	a := example1.Address{City: "b.city", Strt: "b.street"}

	upd := example1.UpdateDocument{}
	upd.Set().SetAddress(a).SetShipAddress(a)

	brel := example1.BusinessRel{
		PublisherId: "Einaudi",
		Contracts: map[string]example1.Contract{
			"C2": {ContractId: "C2", ContractDescr: "le nuvole"}},
	}
	upd.Set().SetBusinessRelsS("Einaudi", brel)

	brel2 := example1.BusinessRel{
		PublisherId: "Longanesi",
		Contracts: map[string]example1.Contract{
			"C3": {ContractId: "C3", ContractDescr: "le nuvole"}},
	}
	upd.Set().SetBusinessRelsS("Longanesi", brel2)

	if ures, err := example1Collection.UpdateOne(ctx, f.Build(), upd.Build(), opts); err != nil {
		t.Fatal(err)
	} else {
		if ures.ModifiedCount != 1 {
			t.Fatal(ures)
		}

		t.Log(ures)
	}

	upd2 := example1.UpdateDocument{}
	ctrt := example1.Contract{ContractId: "C4", ContractDescr: "il mare"}

	upd2.Set().SetBusinessRelsSContractsT("Longanesi", "C4", ctrt)

	if ures, err := example1Collection.UpdateOne(ctx, f.Build(), upd2.Build(), opts); err != nil {
		t.Fatal(err)
	} else {
		if ures.ModifiedCount != 1 {
			t.Fatal(ures)
		}

		t.Log(ures)
	}

}
