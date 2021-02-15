package example1_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/examples/example1"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"testing"
	"time"
)

func TestFilterexample1(t *testing.T) {

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

	if err := queryUsingFilter(t, err, example1Collection, ctx); err != nil {
		t.Fatal(err)
	}
}

func queryUsingFilter(t *testing.T, err error, example1Collection *mongo.Collection, ctx context.Context) error {

	t.Log("queryUsingFilter")

	/*
	 *
	 */
	var f example1.Filter
	t.Log("######## fn = 'John' or fn = 'Colin'")
	f = example1.Filter{}

	f.Or().AndFirstNameEqTo("John")
	f.Or().AndFirstNameEqTo("Colin")

	filterD := f.Build()
	t.Log("resulting filter: ", filterD)
	cur, err := example1Collection.Find(ctx, filterD)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if nr := consumeResultset(t, ctx, cur); nr != 2 {
		t.Error("# records found: ", nr, " expected 2")
	}

	/*
	 *
	 */
	t.Log("######## fn = 'John' and ln = 'Smith'")
	f = example1.Filter{}
	f.Or().AndFirstNameEqTo("John").AndLastNameEqTo("Smith")

	filterD = f.Build()
	t.Log("resulting filter: ", filterD)
	cur, err = example1Collection.Find(ctx, filterD)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if nr := consumeResultset(t, ctx, cur); nr != 1 {
		t.Error("# records found: ", nr, " expected 1")
	}

	/*
	 *
	 */
	t.Log("######## ln = 'Smith' and age gt 29")
	f = example1.Filter{}
	f.Or().AndLastNameEqTo("Smith").AndAgeGt(29)

	filterD = f.Build()
	t.Log("resulting filter: ", filterD)
	cur, err = example1Collection.Find(ctx, filterD)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if nr := consumeResultset(t, ctx, cur); nr != 2 {
		t.Error("# records found: ", nr, " expected 2")
	}

	/*
	 *
	 */
	t.Log("######## (fn = 'John' or fn = 'Colin') and ln = 'Smith'")
	t.Log("that is: (fn = 'John'  and ln = 'Smith') or (fn = 'Colin'  and ln = 'Smith')")

	f = example1.Filter{}
	f.Or().AndFirstNameEqTo("John").AndLastNameEqTo("Smith")
	f.Or().AndFirstNameEqTo("Colin").AndLastNameEqTo("Smith")

	filterD = f.Build()
	t.Log("resulting filter: ", filterD)
	cur, err = example1Collection.Find(ctx, filterD)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if nr := consumeResultset(t, ctx, cur); nr != 2 {
		t.Error("# records found: ", nr, " expected 2")
	}

	/*
	 *
	 */
	t.Log("######## (addr.city = 'Rome' and addr.strt = 'St.Peter Square' and ln = 'Smith' and fn = 'Colin') or (fn = 'John' and books.title = 'My Best Seller')")

	f = example1.Filter{}
	ca := f.Or().AndLastNameEqTo("Smith").AndFirstNameEqTo("Colin")
	ca.AndAddressCityEqTo("Rome").AndAddressStrtEqTo("St.Peter Square")

	ca = f.Or().AndFirstNameEqTo("John")
	ca.AndBooksTitleEqTo("My Best Seller")
	// ca.books().titleEqTo("My Best Seller", 0)

	filterD = f.Build()
	t.Log("resulting filter: ", filterD)
	cur, err = example1Collection.Find(ctx, filterD)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if nr := consumeResultset(t, ctx, cur); nr != 2 {
		t.Error("# records found: ", nr, " expected 2")
	}

	/*
	 *
	 */
	t.Log("######## (fn in [ 'Colin', 'John'])")

	f = example1.Filter{}
	f.Or().AndFirstNameIn([]string{"John", "Colin"})

	filterD = f.Build()
	t.Log("resulting filter: ", filterD)
	cur, err = example1Collection.Find(ctx, filterD)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if nr := consumeResultset(t, ctx, cur); nr != 2 {
		t.Error("# records found: ", nr, " expected 2")
	}
	return nil
}

func removeAll(t *testing.T, err error, example1Collection *mongo.Collection, ctx context.Context) error {
	deleteResult, err := example1Collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		return err
	}
	t.Log(deleteResult)

	return nil
}

func consumeResultset(t *testing.T, ctx context.Context, cur *mongo.Cursor) int {

	numberOfRecords := 0
	for cur.Next(ctx) {
		numberOfRecords++
		var document bson.D
		err := cur.Decode(&document)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(document)
	}

	if err := cur.Err(); err != nil {
		t.Fatal(err)
	}

	return numberOfRecords
}

func insertAFirstRecord(t *testing.T, ctx context.Context, example1Collection *mongo.Collection, fn string, ln string) error {
	a := example1.Author{
		FirstName: fn,
		LastName:  ln,
		Age:       30,
		Address:   example1.Address{City: "Rome", Strt: "St.Peter Square"},
		Books:     []example1.Book{{Title: "My Best Seller", Isbn: "001-12345678", CoAuthors: []string{"co-example1-1", "co-example1-2"}}},
		BusinessRels: map[string]example1.BusinessRel{
			"MC_GRAW_HILL": {
				PublisherId: "MC_GRAW_HILL", PublisherName: "Mc Graw Hill",
				Contracts: map[string]example1.Contract{
					"C1": {ContractId: "C1", ContractDescr: "nuovo libro Harry Potter"}},
			}},
	}

	t.Log("Inserting a record: ", a)

	r, err := example1Collection.InsertOne(ctx, a)
	if err != nil {
		return err
	} else {
		a.OId = r.InsertedID.(primitive.ObjectID)
		if b, err := json.Marshal(a); err != nil {
			return err
		} else {
			fmt.Println(string(b))
		}
	}

	return nil
}

func insertASecondRecord(t *testing.T, ctx context.Context, example1Collection *mongo.Collection, fn string, ln string) error {
	a := example1.Author{
		FirstName: fn,
		LastName:  ln,
		Age:       40,
		Address:   example1.Address{City: "Rome", Strt: "St.Peter Square"},
		Books:     []example1.Book{{Title: "The life Journey"}},
	}

	t.Log("Inserting a record: ", a)

	_, err := example1Collection.InsertOne(ctx, a)
	if err != nil {
		return err
	}

	return nil
}

func insertAThirdRecord(t *testing.T, ctx context.Context, example1Collection *mongo.Collection, fn string, ln string) error {
	a := example1.Author{
		FirstName: fn,
		LastName:  ln,
		Age:       40,
		Address:   example1.Address{City: "Rome", Strt: "St.Peter Square"},
		Books:     []example1.Book{{Title: "The life Journey"}},
	}

	t.Log("Inserting a record: ", a)

	_, err := example1Collection.InsertOne(ctx, a)
	if err != nil {
		return err
	}

	return nil
}
